package utils

import (
	"regexp"
)

var (
	Regexp = newRegexp()
)

type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
}

func (a *Address) GetProvince(hasSuffix bool) string {
	if hasSuffix {
		return a.Province
	}

	reg := regexp.MustCompile(`(省|自治区|市)`)
	return reg.ReplaceAllString(a.Province, "")
}

func (a Address) GetCity(hasSuffix bool) string {
	if hasSuffix {
		return a.City
	}

	reg := regexp.MustCompile(`市|自治州`)
	return reg.ReplaceAllString(a.City, "")
}

func NewAddress(province string, city string, county string) *Address {
	return &Address{Province: province, City: city, County: county}
}

type _regexp struct {
}

func newRegexp() *_regexp {
	return &_regexp{}
}

func (r *_regexp) Match(search, value string) ( *regexp.Regexp, bool) {
	compile, _ := regexp.Compile(search)
	return compile,compile.MatchString(value)
}

func (r _regexp) ReplaceAllString(search string, replace string, value string) string {
	if compile , ok := r.Match(search , value);ok {
		value = compile.ReplaceAllString(value, replace)
	}

	return value
}

func (r *_regexp) IsInteger(str string) bool {
	ok, _ := regexp.MatchString(`^[0-9][0-9]*$`, str)
	return ok
}

func (r _regexp) CheckMobile(mobile string) bool {
	ok, _ := regexp.MatchString(`^1[1-9]\d{9}$`, mobile)
	return ok
}

func (r _regexp) CheckIP(ip string) bool {
	ok, _ := regexp.MatchString(`^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$`, ip)
	return ok
}

func (r _regexp) ParseAddress(addr string) *Address {
	reg := regexp.MustCompile(`([^省]+省|.+自治区|上海市?|北京市?|天津市?|重庆市?)([^市]+市|.+自治州)?([^县]+县|.+区|.+镇|.+局)?`)
	find := reg.FindAllStringSubmatch(addr, -1)
	if len(find) == 0 {
		return nil
	}

	length := len(find[0])
	switch length {
	case 1:
		return nil
	case 2:
		return &Address{
			Province: find[0][1],
		}
	case 3:
		return &Address{
			Province: find[0][1],
			City:     find[0][2],
		}
	case 4:
		return &Address{
			Province: find[0][1],
			City:     find[0][2],
			County:   find[0][3],
		}
	}
	return nil
}
