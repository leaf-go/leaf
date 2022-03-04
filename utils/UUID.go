package utils

import (
	"fmt"
	"time"
)

const (
	twepoch      int64 = 1474992000000
	workerIdBits int64 = 10
	sequenceBits int64 = 12
	maxWorderId  int64 = -1 ^ (-1 << workerIdBits)
	sequenceMask int64 = -1 ^ (-1 << sequenceBits)
)

var (
	lastTs   int64 = -1
	sequence int64 = 0
)

type intUUID struct {
	wordId int64
}

func NewIntUUID(wordId int64) *intUUID {

	if wordId > maxWorderId || wordId < 0 {
		panic(fmt.Sprintf("workerId can't be greater than %d or less than 0", maxWorderId))
	}

	return &intUUID{wordId: wordId}
}

func (u intUUID) NextID() (int64, error) {
	ts := u.milliseconds()
	lts := lastTs

	if ts < lts {
		return 0, fmt.Errorf("Clock moved backwards.  Refusing to generate id for %d milliseconds", lts-ts)
	}

	if lts == ts {
		sequence = (sequence + 1) & sequenceMask
		if sequence == 0 {
			ts = u.tilNextMillis(lts)
		}

	} else {
		sequence = 0
	}

	lastTs = ts
	tsLeftShift := sequenceBits + workerIdBits
	wordIdShift := sequenceBits

	return ((ts - twepoch) << tsLeftShift) | (u.wordId << wordIdShift) | sequence, nil
}

func (u intUUID) milliseconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (u intUUID) tilNextMillis(lastTs int64) int64 {
	ts := u.milliseconds()
	for ts <= lastTs {
		ts = u.milliseconds()
	}

	return ts
}
