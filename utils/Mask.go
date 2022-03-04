package utils

import "strings"

var (
	Mask = newMask()
)

type defaultMask struct {
}

func newMask() *defaultMask {
	return &defaultMask{}
}

func (m *defaultMask) Mobile(phone string) string {
	count := len(phone)
	return phone[0:3] + "****" + phone[count-4:count]
}

func (m defaultMask) Name(name string) string {
	return string([]rune(name)[0:1]) + strings.Repeat("**", len([]rune(name)[1:]))
}

func (m defaultMask) IdCard(cardNo string) string {
	length := len(cardNo[0:])
	return cardNo[0:4] + strings.Repeat("*", length-8) + cardNo[len(cardNo[0:])-4:]
}
