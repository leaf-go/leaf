package utils

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/sha3"
	"hash"
	"hash/crc32"
)

var (
	Hash = newHash()
)

type HashBytes []byte

func (h HashBytes) Origin() []byte {
	return h
}

func (h HashBytes) String() string {
	return fmt.Sprintf("%x", h.Origin())
}

type defaultHash struct{}

func newHash() *defaultHash {
	return &defaultHash{}
}

func (h *defaultHash) Hmac(fn func() hash.Hash, key []byte, data []byte) []byte {
	hs := hmac.New(fn, key)
	hs.Write(data)
	return hs.Sum(nil)
}

func (h defaultHash) Object(hash crypto.Hash) hash.Hash {
	switch hash {
	case crypto.MD5:
		return md5.New()
	case crypto.SHA1:
		return sha1.New()
	case crypto.SHA256:
		return sha256.New()
	case crypto.SHA384:
		return sha3.New384()
	case crypto.BLAKE2b_512:
		return sha3.New512()
	}

	panic("failed to hash type")
}

func (h defaultHash) hash(hash crypto.Hash, data []byte) HashBytes {
	object := h.Object(hash)
	object.Write(data)
	return object.Sum(nil)
}

func (h *defaultHash) Md5(data interface{}) HashBytes {
	return h.hash(crypto.MD5, getBytes(data))
}

func (h *defaultHash) Md5String(data string) string {
	return h.Md5(data).String()
}

func (h *defaultHash) Sha1(data interface{}) HashBytes {
	return h.hash(crypto.SHA1, getBytes(data))
}

func (h *defaultHash) Sha1String(data string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(data))
	return fmt.Sprintf("%x", sha1Hash.Sum(nil))
}

func (h *defaultHash) Sha256(data interface{}) HashBytes {
	return h.hash(crypto.SHA256, getBytes(data))
}

func (h *defaultHash) Sha256String(data string) string {
	return h.Sha256(data).String()
}

func (h *defaultHash) CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
