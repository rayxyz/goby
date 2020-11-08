package ratelimit

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"goby/pkg/dict"

	"golang.org/x/time/rate"
)

const (
	defaultTTL              = 60 * 10 * time.Second // default time to live
	defaultLockTimeDuration = 60 * time.Second      // lock 1 min by default
)

// RateLimiter :
type RateLimiter struct {
	sync.RWMutex
	limiters        map[string]*Limiter
	cleanupInterval time.Duration
}

// Limiter :
type Limiter struct {
	Limiter                *rate.Limiter `json:"-"`
	LastVisitTime          time.Time     `json:"last_visit_time"`
	LastVisitTimeBeautiful string        `json:"last_visit_time_beautiful"`
	Key                    string        `json:"key"`
	SourceName             string        `json:"source_name"` // ex: IP, User ID, etc.
	Source                 string        `json:"source"`      // The actual IP, User ID, etc.
	Match                  string        `json:"match"`
	Matched                string        `json:"matched"` // The actual matched string by match pattern, ex: an URI, etc.
	Rate                   rate.Limit    `json:"rate"`
	Bursts                 int           `json:"bursts"`
	TTL                    time.Duration `json:"ttl"`             // Time to live
	Watch                  bool          `json:"watch"`           // If watched, the watch time duration and max_count settings enabled.
	WatchTime              int32         `json:"watch_time"`      // A number works with WatchTimeUnit to parse as a period of time by seconds.
	WatchTimeUnit          int32         `json:"watch_time_unit"` // Unit: 1 => minute, 2 => hour,3 =>  day, 4 => week, 5 => month.
	MaxCount               int64         `json:"max_count"`       // Number of max operations in defined duration.
	CurrentCount           int64         `json:"current_count"`   // Number of operations executed by far.
	NeedToLock             bool          `json:"need_to_lock"`    // Need to lock calling?
	LockTime               int32         `json:"lock_time"`       // How long to lock up the current actor.
	LockTimeUnit           int32         `json:"lock_time_unit"`
	LockTimeDuration       time.Duration `json:"lock_time_duration"`
	Locked                 bool          `json:"locked"` // Lock up the current actor by logic.
	Cached                 bool          `json:"-"`
}

func (l *Limiter) parseWatchTimeDuration() (time.Duration, error) {
	if l.WatchTime <= 0 {
		return 0, errors.New("invalid watch time")
	}

	return doParseDuration(l.WatchTime, l.WatchTimeUnit)
}

func (l *Limiter) parseLockTimeDuration() (time.Duration, error) {
	if l.LockTime <= 0 {
		return 0, errors.New("invalid lock time")
	}

	return doParseDuration(l.LockTime, l.LockTimeUnit)
}

func doParseDuration(timeDigit, timeUnit int32) (time.Duration, error) {
	var duration time.Duration

	if _, ok := dict.RateLimitRuleUnitMap[timeUnit]; !ok {
		return 0, errors.New("invalid time unit")
	}

	switch timeUnit {
	case dict.RateLimitRuleUnitMinute:
		duration = time.Minute * time.Duration(timeDigit)
	case dict.RateLimitRuleUnitHour:
		duration = time.Hour * time.Duration(timeDigit)
	case dict.RateLimitRuleUnitDay:
		duration = time.Hour * time.Duration(timeDigit)
	case dict.RateLimitRuleUnitWeek:
		duration = 7 * 24 * time.Hour * time.Duration(timeDigit)
	case dict.RateLimitRuleUnitMonth:
		duration = 30 * 24 * time.Hour * time.Duration(timeDigit)
	default:
		return 0, errors.New("no time unit matched => " + string(timeUnit))
	}

	return duration, nil
}

func (l *Limiter) parseRaw() error {
	l.Key = strings.Trim(l.Key, " ")
	if strings.EqualFold(l.Key, "") {
		return errors.New("empty key")
	}
	if l.Rate <= 0 {
		return errors.New("invalid rate")
	}
	if l.Bursts <= 0 {
		return errors.New("invalid bursts")
	}
	if l.TTL <= 0 {
		log.Println("no TTL(Time to Live) set or invalid, watch => ", l.Watch)
		if l.Watch {
			d, err := l.parseWatchTimeDuration()
			if err == nil {
				l.TTL = d
			} else {
				l.TTL = defaultTTL
			}
		} else {
			l.TTL = defaultTTL
		}
	}
	if l.NeedToLock {
		if l.LockTime > 0 {
			d, err := l.parseLockTimeDuration()
			if err == nil {
				l.LockTimeDuration = d
			} else {
				l.LockTimeDuration = defaultLockTimeDuration
			}
		} else {
			l.LockTimeDuration = defaultLockTimeDuration
		}
		if l.LockTimeDuration > l.TTL {
			l.TTL = l.LockTimeDuration
		}
	}

	l.Limiter = rate.NewLimiter(l.Rate, l.Bursts)
	l.LastVisitTime = time.Now()

	return nil
}

// Lockup :
func (l *Limiter) Lockup() {
	log.Printf("Lock up limiter of key %s\n", l.Key)
	l.Locked = true
	time.Sleep(l.LockTimeDuration)
	log.Printf("Unlock limiter of key %s\n", l.Key)
	l.Locked = false
	if l.Watch {
		l.CurrentCount = 0 // reset the count
	}
}

// NewRateLimiter : init rate limiters
func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		limiters:        make(map[string]*Limiter),
		cleanupInterval: 20 * time.Second,
	}

	go rl.Cleanup()

	return rl
}

// Add : add limiter to limiter map
func (rl *RateLimiter) Add(l *Limiter) (*Limiter, error) {
	rl.Lock()
	defer rl.Unlock()

	limiter, ok := rl.limiters[l.Key]
	if !ok {
		if err := l.parseRaw(); err != nil {
			return nil, err
		}
		rl.limiters[l.Key] = l
	}
	limiter, _ = rl.limiters[l.Key]

	limiter.LastVisitTime = time.Now()

	return limiter, nil
}

// Get : get the keyed rate limiter
func (rl *RateLimiter) Get(key string) *Limiter {
	rl.Lock()
	defer rl.Unlock()

	limiter, ok := rl.limiters[key]
	if !ok {
		return nil
	}

	return limiter
}

// GetAll :
func (rl *RateLimiter) GetAll() []*Limiter {
	var list []*Limiter
	for _, v := range rl.limiters {
		list = append(list, v)
	}

	return list
}

// Count :
func (rl *RateLimiter) Count() int {
	return len(rl.limiters)
}

// Delete limiters
func (rl *RateLimiter) Delete(keysList []string) {
	fmt.Println("in delete rate limiters...")
	rl.Lock()
	defer rl.Unlock()

	for _, v := range keysList {
		delete(rl.limiters, v)
	}
}

// Cleanup :
func (rl *RateLimiter) Cleanup() {
	for {
		time.Sleep(rl.cleanupInterval)

		go func() {
			rl.Lock()
			defer rl.Unlock()

			for k, v := range rl.limiters {
				if time.Now().Sub(v.LastVisitTime) > v.TTL {
					log.Println("Delete limiter of key => ", k)
					if !v.Locked {
						delete(rl.limiters, k)
					}
				}
			}
		}()
	}
}
