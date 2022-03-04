package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	Rand        = NewRand()
	letters     = []byte("abcdefghijkmnpqrstuvwxyz123456789")
	longLetters = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type _rand struct {
}

func NewRand() *_rand {
	return &_rand{}
}

func (r _rand) Rand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func (r _rand) Default() int {
	return r.Rand(1e10, 1<<15-1)
}

func (r _rand) RandLowerString(n int) string {
	return string(r.randString(n, 31, letters))
}

func (r _rand) RandString(n int) string {
	return string(r.randString(n, 62, longLetters))
}

func (r _rand) RandBytes(n int) []byte {
	return r.randString(n, 63, longLetters)
}

func (r _rand) randString(n int, bit byte, rangeBytes []byte) []byte {
	if n <= 0 {
		return []byte{}
	}

	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}

	arc := uint8(0)
	for i, x := range b {
		arc = x & bit
		if arc > 0 {
			arc -= 1
		}
		b[i] = rangeBytes[arc]
	}

	return b
}

func (r _rand) UniqString() string {
	timeString := strconv.FormatInt(time.Now().UnixNano(),10)
	randString := r.RandString(32)
	s := fmt.Sprintf("%s_%s" , randString , timeString)
	return Hash.Md5String(s)
}
