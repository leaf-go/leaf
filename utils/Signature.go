package utils

import (
	"fmt"
	"time"
)

var (
	Signature = newSignature()
)

type signature struct {
	factor string
	expire int64
}

func newSignature() *signature {
	return &signature{
		factor: "@cRL@0xokn4qL5j2",
		expire: 180,
	}
}

func (s *signature) Make(data interface{}) (string, error) {
	json, err := Json.Marshal(data)
	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("%s@%s", json, s.factor)
	sign := Hash.Md5(str)

	return sign.String(), nil
}

func (s *signature) Valid(data map[string]interface{}, sign string) bool {
	t, ok := data["ts"]
	if !ok {
		return false
	}

	if !s.checkTs(t) {
		return false
	}

	sig, e := s.Make(data)
	if e != nil {
		return false
	}

	fmt.Println(sig ,sign)
	return sig == sign
}

func (s signature) checkTs(t interface{}) bool {
	var ts int64
	switch t.(type) {
	case int:
		ts = int64(t.(int))
		break
	case int64:
		ts = t.(int64)
	default:
		return false
	}

	now := time.Now().UnixNano() / 1e6
	return now-ts <= s.expire*1000
}
