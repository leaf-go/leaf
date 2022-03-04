package utils

import (
	"regexp"
)

var (
	Validate = newValidate()
)

type validate struct{}

func newValidate() *validate {
	return &validate{}
}

func (v validate) Mobile(mobile string) bool {
	return Regexp.CheckMobile(mobile)
}

func (v validate)Version(version string) bool {
	ok, _ := regexp.MatchString(`^\d{1,3}\.\d{1,3}\.\d{1,3}$` , version)
	return ok
}

func (v validate) MobileWithTel(val string) bool  {
	ok, _ := regexp.MatchString(`^(0\d{2}-?[1-9]\d{7})|(0\d{3}-?[1-9]\d{6})|(1[1-9]\d{9})$` , val)
	return ok
}

func (v validate) BankCard(bankCard string) bool {
	odd := len(bankCard) & 1
	var sum int
	for i, c := range bankCard {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}

