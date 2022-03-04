package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"math/big"
)

var (
	AES = newDefaultAES()
)

type cryptHandler func(det, src []byte)

func newDefaultAES() *defaultAES {
	return &defaultAES{}
}

type defaultAES struct{}

func (a *defaultAES) Encrypt(content, key []byte) (signature string, err error) {
	if len(content) == 0 {
		return "", errors.New("src is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	content = a.PKCS5Padding(content, block.BlockSize())
	result := make([]byte, len(content))
	err = a.CryptBlocks(true, block, result, content)

	// 普通base64编码加密 区别于urlsafe base64
	return base64.StdEncoding.EncodeToString(result), err
}

func (a *defaultAES) Decrypt(crypted, key []byte) (resultString string, err error) {
	decrypted, err := base64.StdEncoding.DecodeString(string(crypted))
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	result := make([]byte, len(decrypted))
	err = a.CryptBlocks(false, block, result, decrypted)
	return string(a.PKCS5UnPadding(result)), err
}

func (a defaultAES) CryptBlocks(isEncrypt bool, block cipher.Block, result, src []byte) (err error) {
	length := len(src)
	size := block.BlockSize()
	if length%size != 0 {
		return errors.New("crypto/cipher: input not full blocks")
	}

	if len(result) < len(src) {
		errors.New("crypto/cipher: DefaultOutput smaller than input")
	}

	var handler cryptHandler
	if isEncrypt {
		handler = block.Encrypt
	} else {
		handler = block.Decrypt
	}

	for len(src) > 0 {
		handler(result, src[:size])
		src = src[size:]
		result = result[size:]
	}

	return
}

func (a *defaultAES) EncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(a.GenerateKeyECB(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func (a defaultAES) GenerateKey(l int) []byte {
	return Rand.RandBytes(l)
}

func (a *defaultAES) GenerateKeyECB(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// 支持NoPadding
func (a *defaultAES) Encrypt128(origData, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = a.PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData)) // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (a *defaultAES) Decrypt128(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = a.PKCS5UnPadding(origData)
	return origData, nil
}

func (a *defaultAES) NoPadding(cipherText []byte) []byte {
	c := new(big.Int).SetBytes(cipherText)
	privateKey := rsa.PrivateKey{}
	plainText := c.Exp(c, privateKey.D, privateKey.N).Bytes()
	return plainText
}

func (a *defaultAES) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (a *defaultAES) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
