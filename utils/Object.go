package utils

import (
	"reflect"
)

var (
	Object = &object{}
)

type object struct{}

func (o *object) NilJsonObject(v interface{}) interface{} {
	if reflect.ValueOf(v).IsNil() {
		return struct{}{}
	}

	return v
}

func (o *object) NilObject() struct{}  {
	return struct{}{}
}
