package utils

import "fmt"

var (
	Format = newDefaultFormat()
)

type defaultFormat struct {
}

func newDefaultFormat() *defaultFormat {
	return &defaultFormat{}
}

func (f *defaultFormat) YmdStringBirthday(year string, month string, day string) string {
	if len(month) < 2 {
		month = "0" + month
	}

	if len(day) < 2 {
		day = "0" + day
	}
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
