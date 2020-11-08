package util

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestUUID(t *testing.T) {
	uuid := GenUUID()
	fmt.Println("uuid => ", uuid)
}

func TestTimeStamp(t *testing.T) {
	// timeNow := time.Now().Add(-15 * time.Minute)
	// timeNow := time.Now().AddDate(0, 0, -2).Add(9 * time.Hour)
	timeNow := time.Now().AddDate(0, 0, -6).Add(4 * time.Hour).Add(-37 * time.Minute)
	// timeNow, _ := time.Parse("2006-01-02 15:04:05", "2018-09-12 18:05:23") // wrong
	fmt.Println("time_now => ", timeNow)
	timestamp := timeNow.Unix()
	// timestamp := 1536748480
	fmt.Println("timestamp => ", timestamp)
}

func TestSignature(t *testing.T) {
	// origin := strings.Join([]string{"5255421148254165520", "4743935625164463028", "JZ9S4ZyNRwS57dX95uNPnBMLHxXzgJ8a", "4743935625164463028", "1531299087", "466893"}, "")
	// signature := GenMD5HashCode(origin)
	// fmt.Println("signature => ", signature)

	origin := strings.Join([]string{"5031701270659586583", "5503296779358656963", "Sw3GVXjva2RtDgHcxnMb9LCq6EJTANkm", "5503296779358656963", "1548669832", "536048"}, "")
	signature := GenMD5HashCode(origin)
	fmt.Println("signature => ", signature)
}

func TestTimeStampToTime(t *testing.T) {
	i, err := strconv.ParseInt("1531299087", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0).Format("2006-01-02 15:04:05")
	fmt.Println("time parsed from timestamp => ", tm)
}

func TestNanoSeconds(t *testing.T) {
	nanos := time.Now().UnixNano()
	// 1533111814897704794
	fmt.Println("nanos => ", nanos)
}

func TestFormat2String(t *testing.T) {
	numbers := []float64{123, 43543, 546.56, 2365879787, 4, 6, 7444357, 4832593876584, 213423, 654.435, 3465546}
	for _, v := range numbers {
		numStr := FormatNumber2String(int64(v))
		fmt.Println("numStr => ", numStr)
	}
}

func TestFileSize(t *testing.T) {
	size := (2 * 1 << 30) / (1024 * 1024 * 1024)
	fmt.Println("size => ", size)
}

func TestArigth(t *testing.T) {
	v := 5 / 4
	fmt.Println(v)
}

func TestDuration(t *testing.T) {
	dur := time.Duration(2 * time.Second)
	fmt.Println("dur => ", dur)
}

func TestTimeSub(t *testing.T) {
	createTime, _ := time.Parse("2006-01-02 15:04:05", "2020-06-16 15:47:11")
	now, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	minutes := now.Sub(createTime).Minutes()
	fmt.Println("\ncreateTime => ", createTime)
	fmt.Println("now => ", time.Now())
	fmt.Printf("minutes => %f\n", minutes)
}
