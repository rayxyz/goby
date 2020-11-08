package util

import (
	"errors"
	"time"
)

// 日期、时间处理相关

// FirstDayOfISOWeek : 某年某周的第一天
func FirstDayOfISOWeek(year int, week int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.Now().Local().Location())
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	return date
}

// LastDayOfISOweek : 某年某周的最后一天
func LastDayOfISOweek(year int, week int) time.Time {
	dateTime := FirstDayOfISOWeek(year, week)
	dateTime = time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day()+6, 23, 59, 59, 0, time.Now().Local().Location())
	return dateTime
}

// FirstDayOfMonth : 某年某月的第一天
func FirstDayOfMonth(year int, month int) time.Time {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Local().Location())
	return firstDay
}

// LastDayOfMonth : 某年某月的最后一天
func LastDayOfMonth(year int, month int) time.Time {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Local().Location())
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDay
}

// FirstDayOfLastMonth : 上个月最后一天
func FirstDayOfLastMonth(year int, month int) time.Time {
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Local().Location())
	if month == 1 {
		date = date.AddDate(-1, 11, 0)
	} else {
		date = date.AddDate(0, -1, 0)
	}
	return date
}

// GetFirstAndLastSecondOfDate : 某天第一秒和最后一秒
func GetFirstAndLastSecondOfDate(dateTime time.Time) (time.Time, time.Time) {
	firstSecond := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, time.Now().Local().Location())
	lastSecond := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 23, 59, 59, 0, time.Now().Local().Location())
	return firstSecond, lastSecond
}

// GetFirstAndLastSecondOfDateString : 某天第一秒和最后一秒（字符串）
func GetFirstAndLastSecondOfDateString(dateTime time.Time) (string, string) {
	firstSecond, lastSecond := GetFirstAndLastSecondOfDate(dateTime)
	return firstSecond.Format("2006-01-02 15:04:05"), lastSecond.Format("2006-01-02 15:04:05")
}

// ParseRFC3339ToMDHM : time.RFC3339 to `yyyy-mm-dd hh:mm:ss`
func ParseRFC3339ToMDHM(timeStr string) string {
	timeparsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return timeparsed.Format("01-02 15:04")
}

// ParseRFC3339ToYMDHMS : time.RFC3339 to `yyyy-mm-dd hh:mm:ss`
func ParseRFC3339ToYMDHMS(timeStr string) string {
	timeparsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return timeparsed.Format("2006-01-02 15:04:05")
}

// ParseRFC3339ToYMDHM : time.RFC3339 to `yyyy-mm-dd hh:mm`
func ParseRFC3339ToYMDHM(timeStr string) string {
	timeparsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return timeparsed.Format("2006-01-02 15:04")
}

// ParseRFC3339ToYMD : time.RFC3339 to `yyyy-mm-dd hh:mm:ss`
func ParseRFC3339ToYMD(timeStr string) string {
	timeparsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return timeparsed.Format("2006-01-02")
}

// ParseRFC3339ToHMS : time.RFC3339 to `hh:mm:ss`
func ParseRFC3339ToHMS(timeStr string) string {
	timeparsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return timeparsed.Format("15:04:05")
}

// ParseRFC3339ToHM : time.RFC3339 to `hh:mm`
func ParseRFC3339ToHM(timeStr string) string {
	timeparsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	time.Now().Month()
	return timeparsed.Format("15:04")
}

// CheckHMTimeStrValid : Checking time string "hh:mm" valid
func CheckHMTimeStrValid(hmstr string) bool {
	if _, err := time.Parse("2006-01-02 15:04", time.Now().Format("2006-01-02")+" "+hmstr); err != nil {
		return false
	}
	return true
}

// CheckYMDHMSTimeStrValid : checking time string "yyyy-mm-dd hh:mm:ss" valid
func CheckYMDHMSTimeStrValid(ymdhms string) bool {
	if _, err := time.Parse("2006-01-02 15:04:05", ymdhms); err != nil {
		return false
	}
	return true
}

// GetDateListByRange : 根据日期范围返回日期列表
func GetDateListByRange(startDate, endDate time.Time) ([]time.Time, error) {
	if startDate.IsZero() || endDate.IsZero() || startDate.After(endDate) {
		return nil, errors.New("error of generating date list")
	}
	var dateList []time.Time
	for d := startDate; d.Before(endDate) || d.Equal(endDate); d = d.AddDate(0, 0, 1) {
		// fmt.Println("now_date => ", d.Format("2006-01-02"))
		dateList = append(dateList, d)
	}
	return dateList, nil
}

// TimeStamp2YMDHMS : convert unix timestamp to ymdhms format datetime string
func TimeStamp2YMDHMS(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// Week :
type Week struct {
	StartDateTime time.Time
	EndDateTime   time.Time
}

// GenWeeks : 生成周（不包含星期一和星期日）
// mondayDateTime => 星期一的日期
// weekCount => 要生成的周数
func GenWeeks(mondayDateTime time.Time, weekCount int) ([]*Week, error) {
	var weekList []*Week
	if mondayDateTime.Weekday() != time.Monday {
		return nil, errors.New("the start date is not Monday")
	}
	if weekCount <= 0 {
		return nil, errors.New("invalid week count value")
	}
	startDateTime := mondayDateTime
	endDateTime := mondayDateTime
	for i := 0; i < weekCount; i++ {
		if endDateTime.Equal(startDateTime) {
			endDateTime = endDateTime.AddDate(0, 0, 4)
			weekList = append(weekList, &Week{
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			})
			endDateTime = endDateTime.AddDate(0, 0, 2)
			continue
		}
		startDateTime = endDateTime.AddDate(0, 0, 1)
		endDateTime = startDateTime.AddDate(0, 0, 4)
		weekList = append(weekList, &Week{
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		})
		endDateTime = endDateTime.AddDate(0, 0, 2)
	}

	return weekList, nil
}

// Time2UTCTime : convert system time to UTC time
func Time2UTCTime(goTime time.Time) time.Time {
	timeNowParsed, _ := time.Parse("2006-01-02 15:04:05", goTime.Format("2006-01-02 15:04:05"))
	return timeNowParsed
}

// Weekday2CN :
func Weekday2CN(weekday int) string {
	weekdayCN := ""

	switch weekday {
	case 0:
		weekdayCN = "星期日"
	case 1:
		weekdayCN = "星期一"
	case 2:
		weekdayCN = "星期二"
	case 3:
		weekdayCN = "星期三"
	case 4:
		weekdayCN = "星期四"
	case 5:
		weekdayCN = "星期五"
	case 6:
		weekdayCN = "星期六"
	}

	return weekdayCN
}
