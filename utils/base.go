package utils

import (
	"fmt"
	"strings"
)

func DD(args ...interface{}) {
	repeat := strings.Repeat("%+v ", len(args)) + "\n"
	fmt.Printf(repeat, args...)
}
