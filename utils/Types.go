package utils

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"html"
	"strconv"
	"strings"
)

var (
	xss   *bluemonday.Policy
	strip *bluemonday.Policy
)

func init() {
	xss = bluemonday.NewPolicy()
	xss.AllowAttrs("src").OnElements("img")
	xss.AllowElements("p", "b", "strong", "div", "span", "ul", "li", "font", "label")
	xss.AllowRelativeURLs(false)
	xss.AllowURLSchemes("http", "https")

	strip = bluemonday.StrictPolicy()
}

// BigInt int64
type BigInt int64

func (bi BigInt) ToByte() Byte {
	return Byte(bi)
}

// Integer int、int32
type Integer int

func (i Integer) ToByte() Byte {
	return Byte(i)
}

// SmallInt int8
type SmallInt int8

// Rune rune unit32
type Rune rune

// Runes []rune
type Runes []Rune

type RunesArray []Rune

// Byte byte uint8
type Byte byte

type IBytes interface {
	Origin() []byte
	String() string
}

type Bytes []byte

func (b Bytes) Origin() []byte {
	return b
}

func (b Bytes) String() string {
	return fmt.Sprintf("%s", b.Origin())
}

func getBytes(data interface{}) []byte {
	switch data.(type) {
	case []byte:
		return data.([]byte)
	case string:
		return []byte(data.(string))
	}

	panic("hash only supports bytes and string")
}

// String string
type String string

func (s String) Sub(start, length int) String {
	if length <= 0 {
		return s[start:]
	}

	return s[start:length]
}

func (s String) Bytes() []byte {
	return []byte(s)
}

func (s String) String() string {
	return string(s)
}

func (s String) Int() (int, error) {
	return strconv.Atoi(s.String())
}

type SafeString String

func (ss SafeString) Escape() string {
	s := ss.Utf8()
	return html.EscapeString(s)
}

func (ss SafeString) EscapeHTML() string {
	s := ss.Utf8()
	return xss.Sanitize(s)
}

func (ss SafeString) Utf8() string {
	s := strings.TrimSpace(string(ss))
	return strings.ToValidUTF8(s, "")
}
