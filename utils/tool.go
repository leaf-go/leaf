package utils

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SortItem struct {
	Position int //下标位置
	Value    int //需要排序的值
}

func TimeYM() string {
	return time.Now().Format("200601")
}

func TimeYMD() string {
	return time.Now().Format("20060102")
}

func TimeYMDH(tt ...time.Time) string {
	var t time.Time

	if len(tt) > 0 {
		t = tt[0]
	} else {
		t = time.Now()
	}

	return t.Format("2006010215")
}

func TimeYMDHIS(tt ...time.Time) string {
	var t time.Time

	if len(tt) > 0 {
		t = tt[0]
	} else {
		t = time.Now()
	}

	return t.Format("2006-01-02 15:04:05")
}

// @des 时间转换函数
//@param atime string 要转换的时间戳（秒）
func Time2Str(atime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = tempStr + unit[i]
		}
		break
	}
	return res
}

// @des 时间转换函数
//@param atime string 要转换的时间戳（秒）
func TimeInDetail(atime int64) string {
	todayStr := TimeYMD()
	at := time.Unix(atime, 0)
	adayStr := at.Format("20060102")

	if todayStr == adayStr {
		return at.Format("15:04")
	}

	if time.Now().Year() > at.Year() {
		return at.Format("2006/01/02 15:04")
	}

	today, _ := strconv.Atoi(todayStr)
	aday, _ := strconv.Atoi(adayStr)

	if today-aday == 1 {
		return "昨天" + at.Format("15:04")
	} else {
		return at.Format("01/02 15:04")
	}
}

// @des 时间转换函数
//@param atime string 要转换的时间戳（秒）
func TimeInList(atime int64) string {
	if atime == 0 || (time.Now().Unix()-atime) < 60 {
		return "刚刚"
	}

	at := time.Unix(atime, 0)
	today := TimeYMDH()
	aday := TimeYMDH(at)

	if today == aday {
		return fmt.Sprintf("%d分钟前", time.Now().Minute()-at.Minute())
	}

	todayStr := TimeYMD()
	adayStr := at.Format("20060102")

	if todayStr == adayStr {
		return fmt.Sprintf("%d小时前", time.Now().Hour()-at.Hour())
	}

	todays, _ := strconv.Atoi(todayStr)
	adays, _ := strconv.Atoi(adayStr)

	if (todays - adays) == 1 {
		return "昨天"
	}

	if time.Now().Year() == at.Year() {
		return at.Format("01/02")
	} else {
		return at.Format("2006/01/02")
	}
}

func NumericToString(v interface{}) string {
	switch v.(type) {
	case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		return fmt.Sprintf("%d", v)
	case float64, float32:
		return fmt.Sprintf("%f", v)
	default:
		if val, ok := v.(string); ok {
			return val
		}
		return ""
	}
}

func FormatNumericFloat64(v float64, decimal int) (data float64) {
	d := float64(1)

	if decimal > 0 {
		d = math.Pow10(decimal)
	}

	vv := strconv.FormatFloat(math.Trunc(v*d)/d, 'f', -1, 64)
	data, _ = strconv.ParseFloat(vv, 64)

	return
}

func FormatNumericString(val string, decimal int) (data float64) {
	v, _ := strconv.ParseFloat(val, 64)
	return FormatNumericFloat64(v, decimal)
}

func SortMapStringInterface(data map[string]interface{}, key, val string) []map[string]interface{} {
	keys := []string{}

	for k, _ := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	res := []map[string]interface{}{}

	for _, k := range keys {
		res = append(res, map[string]interface{}{
			key: k,
			val: data[k],
		})
	}

	return res
}

func SortMapIntInterface(data map[int]interface{}, key, val string) []map[string]interface{} {
	keys := []int{}

	for k, _ := range data {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	res := []map[string]interface{}{}

	for _, k := range keys {
		res = append(res, map[string]interface{}{
			key: k,
			val: data[k],
		})
	}

	return res
}

func InArrayInt64(v int64, arr []int64) bool {
	if len(arr) == 0 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func InArrayString(v string, arr []string) bool {
	if len(arr) == 0 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}

	return false
}

func QuickSort(data []SortItem) {
	if len(data) <= 1 {
		return
	}
	mid := data[0].Value
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i].Value > mid {
			//交换值和交换值对应的位置
			data[i].Value, data[tail].Value = data[tail].Value, data[i].Value
			data[i].Position, data[tail].Position = data[tail].Position, data[i].Position
			tail--
		} else {
			data[i].Value, data[head].Value = data[head].Value, data[i].Value
			data[i].Position, data[head].Position = data[head].Position, data[i].Position
			head++
			i++
		}
	}

	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

func ParseVersion(v string) (version int, err error) {
	val := strings.Split(v, ".")
	count := len(val)

	if count == 0 || count > 3 {
		err = errors.New("version parse error")
		return
	}

	var data []int

	for _, v := range val {
		i, _ := strconv.Atoi(v)
		data = append(data, i)
	}

	if count == 3 {
		major, minor, patch := data[0], data[1], data[2]
		version = major*1000000 + minor*1000 + patch
	} else if count == 2 {
		major, minor := data[0], data[1]
		version = major*1000000 + minor*1000
	} else {
		major := data[0]
		version = major * 1000000
	}

	return
}

func IF(is bool, t interface{}, f interface{}) interface{} {
	if is {
		return t
	} else {
		return f
	}
}
