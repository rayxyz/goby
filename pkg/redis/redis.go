package redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// ConnectConf :
type ConnectConf struct {
	Addr     string
	Password string
	DB       int
}

// NewClient : New a redis client
func NewClient(conf *ConnectConf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	return client
}
