package utils

import (
	"reflect"
)

var (
	Slice = newSlice()
)

type defaultSlice struct {
}

func newSlice() *defaultSlice {
	return &defaultSlice{}
}

func (s *defaultSlice) KindIn(needle reflect.Kind , arrays []reflect.Kind) (num int , exist bool)  {
	for k, v := range arrays {
		if v == needle {
			return k, true
		}
	}

	return -1, false
}


func (d defaultSlice) InSliceInt(needle int, arrays []int) (index int, exist bool) {
	for k, v := range arrays {
		if v == needle {
			return k, true
		}
	}

	return -1, false
}

func (d defaultSlice) Value(slice interface{}, index int, def interface{}) interface{} {
	reflectSlice := reflect.TypeOf(slice)
	if reflectSlice.Kind() != reflect.Slice {
		return def
	}

	//空指针
	if slice == nil {
		return def
	}

	_len := len(slice.([]interface{})) - 1
	if index > _len {
		return def
	}

	return slice.([]interface{})[index]
}

func (d defaultSlice) IntValue(slice []int, index, def int) int {
	return d.Value(slice, index, def).(int)
}

func (d defaultSlice) StringValue(slice []string, index int, def string) string {
	return d.Value(slice, index, def).(string)
}

func (d defaultSlice) columnsString() {

}
