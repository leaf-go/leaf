package utils

import (
	"os"
	"sync/atomic"
	"time"
	"unsafe"
)

type TraceID string

const (
	hexByteBits        = 4
	uint64Bytes        = 8
	localIPEnvironName = "TRACE_LOCAL_IP"
)

var (
	unixnano      int64 = time.Now().UnixNano()
	processID     int
	traceSequence int32
	traceSource   byte = 0xda
	hexBytes           = [1 << hexByteBits]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
	}
)

func init() {
	processID = os.Getpid() & (1<<16 - 1)
}

func format(dst []byte, value uint64, bytes int) []byte {
	data := (*[uint64Bytes]byte)(unsafe.Pointer(&value))[:bytes]

	for i := 0; i < bytes; i++ {
		v := data[bytes-i-1]
		dst[2*i] = hexBytes[v>>hexByteBits]
		dst[2*i+1] = hexBytes[v&(1<<hexByteBits-1)]
	}

	return dst[bytes*2:]
}

func makeString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func hexString(value uint64) string {
	buf := make([]byte, 16)
	format(buf, value, 8)
	return makeString(buf)
}

func MakeTraceID() string {
	ip, _ := Net.LocalIP()
	ts := (unixnano / int64(time.Second)) & (1<<32 - 1)
	reserved := 0
	seq := atomic.AddInt32(&traceSequence, 1) & (1<<24 - 1)

	buf := make([]byte, 32)
	b := buf
	b = format(b, uint64(ip[0]), 1)
	b = format(b, uint64(ip[1]), 1)
	b = format(b, uint64(ip[2]), 1)
	b = format(b, uint64(ip[3]), 1)
	b = format(b, uint64(ts), 4)
	b = format(b, uint64(reserved), 2)
	b = format(b, uint64(processID), 2)
	b = format(b, uint64(seq), 3)
	b = format(b, uint64(traceSource), 1)
	tid := makeString(buf)
	return tid
}

func MakeRequestId() string {
	return hexString(uint64(time.Now().UnixNano()))
}

func GetTraceID(traceID string) TraceID {
	return TraceID(traceID)
}
