package utils

import (
	"fmt"
	"strconv"
	"time"
)

var (
	Time = &defaultTime{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"20060102",
		"15:04:05",
	}
)

type defaultTime struct {
	dateFmt     string
	dateTimeFmt string
	notSign     string
	timeFmt     string
}

func (d *defaultTime) OnlyTime() string {
	return time.Now().Format(d.timeFmt)
}

func (d *defaultTime) DateTime() string {
	return time.Unix(time.Now().Unix(), 0).Format(d.dateFmt)
}

func (d *defaultTime) TimestampString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func (d *defaultTime) Timestamp() int {
	return int(time.Now().Unix())
}

func (d *defaultTime) TomorrowTime() time.Time {
	ts := time.Now().Format(d.dateFmt)
	t2, _ := time.ParseInLocation(d.dateFmt, ts, time.Local)
	return t2.AddDate(0, 0, 1)
}

func (d *defaultTime) TomorrowSubNow() time.Duration {
	now := time.Now()
	return d.TomorrowTime().Sub(now)
}

func (d *defaultTime) DateAtInt() int {
	at := d.DateAt()
	i, _ := strconv.Atoi(at)
	return i
}

// DateAt 获取无符号时间
func (d *defaultTime) DateAt() string {
	return time.Now().Format(d.notSign)
}

// ParseNotSignDate 解析无符号时间
func (d *defaultTime) ParseNotSignDate(date string) (string, error) {
	parsed, err := time.ParseInLocation(d.notSign, date, time.Local)
	if err != nil {
		return "", err
	}

	return parsed.Format(d.dateFmt), nil
}

func (d *defaultTime) ParseWithTime(t string) (time.Time, error) {
	ts := fmt.Sprintf("%s %s", time.Now().Format(d.dateFmt), t)
	return time.ParseInLocation(d.dateTimeFmt, ts, time.Local)
}

func (d *defaultTime) IsAfterTimes(first string, second string) bool {
	fst, _ := d.ParseWithTime(first)
	sec, _ := d.ParseWithTime(second)
	return fst.After(sec)
}

func (d *defaultTime) after(current time.Time, target string) bool {
	if target == "" {
		return true
	}

	tgt, _ := d.ParseWithTime(target)
	return current.After(tgt)
}

// BetweenWithTimes 当前时间是否在某个时间段内 例如 10:00:00 ~ 23:59:59
func (d *defaultTime) BetweenWithTimes(first, second string) bool {
	current := time.Now()
	a, b := 0, 0
	if d.after(current, first) {
		a = 1
	}

	if d.after(current, second) {
		b = 1
	}

	return a&b == 1
}
