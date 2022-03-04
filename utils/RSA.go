package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const (
	RSA_CLASSES_PKCS1 = iota
	RSA_CLASSES_PKCS8
)

// 不需要HASH
const NO_HASH crypto.Hash = 0

var (
	RSA = newRsa()
)

type defaultRsa struct{}

func newRsa() *defaultRsa {
	return &defaultRsa{}
}

//FormatPublicKeyFromString 格式化公钥
func (r *defaultRsa) FormatPublicKeyFromString(publicKey string) []byte {
	format := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)
	return []byte(format)
}

//FormatPrivateKeyFromString 格式化私钥
func (r defaultRsa) FormatPrivateKeyFromString(privateKey string) []byte {
	format := fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", privateKey)
	return []byte(format)
}

//PubKeyEncrypt 公钥加密
func (r *defaultRsa) PubKeyEncrypt(data, pubKey []byte) (encrypt string, err error) {
	// 解析公钥
	p, _ := pem.Decode(pubKey)

	pubInterface, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		return "", err
	}

	pub := pubInterface.(*rsa.PublicKey)
	//加密
	v15Encode, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return "", err
	}

	base64Encode := base64.StdEncoding.EncodeToString(v15Encode)
	return base64Encode, nil
}

func (r *defaultRsa) PriKeyDecodeString(base64Key string, priKey string) (decript string, err error) {
	priBytes, _ := base64.StdEncoding.DecodeString(base64Key)
	localPriBytes := r.FormatPrivateKeyFromString(priKey)
	decodeBytes, err := r.priKeyDecode(priBytes, localPriBytes)
	decript = string(decodeBytes)
	return
}

func (r *defaultRsa) PriKeyDecode(data, priKey []byte) (decript []byte, err error) {
	return r.priKeyDecode(data, priKey)
}

// PriKeyDecode 私钥解密
func (r *defaultRsa) priKeyDecode(data, priKey []byte) (decript []byte, err error) {
	//解析私钥
	p, _ := pem.Decode(priKey)
	priv, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return
	}

	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

func (r *defaultRsa) PriKeyDecodePKCS8String(base64Key string, priKey string) (decript []byte, err error) {
	priBytes, _ := base64.StdEncoding.DecodeString(base64Key)
	localPriBytes := r.FormatPrivateKeyFromString(priKey)
	return r.priKeyDecodePKCS8(priBytes, localPriBytes)
}

func (r *defaultRsa) PriKeyDecodePKCS8(data, priKey []byte) (decript []byte, err error) {
	return r.priKeyDecodePKCS8(data, priKey)
}

//PriKeyDecode 私钥解密
func (r *defaultRsa) priKeyDecodePKCS8(data, priKey []byte) (decrypt []byte, err error) {
	//解析私钥
	p, _ := pem.Decode(priKey)
	key, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return
	}

	// 解密
	decrypt, err = rsa.DecryptPKCS1v15(rand.Reader, key.(*rsa.PrivateKey), data)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

//privateKey 解密
func (r *defaultRsa) privateKey(classes int, priKey []byte) (privateKey *rsa.PrivateKey, err error) {
	//解析私钥
	p, _ := pem.Decode(priKey)

	switch classes {
	case RSA_CLASSES_PKCS1:
		return x509.ParsePKCS1PrivateKey(p.Bytes)
	case RSA_CLASSES_PKCS8:
		pri, err := x509.ParsePKCS8PrivateKey(p.Bytes)
		if err == nil {
			return pri.(*rsa.PrivateKey), err
		}

		return nil, err
	}

	return privateKey, errors.New("_crypt.RSA.privateKey: classes not in (RSA_CLASSES_PKCS1, RSA_CLASSES_PKCS8)")
}

func (r defaultRsa) publicKey(classes int, pubKey []byte) (publicKey *rsa.PublicKey, err error) {
	p, _ := pem.Decode(pubKey)

	switch classes {
	case RSA_CLASSES_PKCS1:
		return x509.ParsePKCS1PublicKey(p.Bytes)
	case RSA_CLASSES_PKCS8:
		pub, err := x509.ParsePKIXPublicKey(p.Bytes)
		if err == nil {
			return pub.(*rsa.PublicKey), err
		}
		return nil, err
	}

	return publicKey, errors.New("_crypt.RSA.publicKey: classes not in (RSA_CLASSES_PKCS1, RSA_CLASSES_PKCS8)")
}

// hashed 生成hash
func (r defaultRsa) hashed(hash crypto.Hash, data []byte) (b []byte, err error) {
	if hash == NO_HASH {
		return data, nil
	}

	if hash == crypto.MD5 {
		return Hash.Md5(data), nil
	}

	if hash == crypto.SHA1 {
		return Hash.Sha1(data), nil
	}

	if hash == crypto.SHA256 {
		return Hash.Sha256(data), nil
	}

	return b, errors.New("_crypt.RSA.hashed not found " + hash.String())
}

//VerifyMd5WithPKCS1 md5和pkcs1 验签
func (r *defaultRsa) VerifyMd5WithPKCS1(data, pubKey []byte, sig string) (err error) {
	return r.verifySignature(RSA_CLASSES_PKCS1, crypto.MD5, data, pubKey, sig)
}

//VerifyMd5WithPKCS8 md5和pkcs8 验签
func (r *defaultRsa) VerifyMd5WithPKCS8(data, pubKey []byte, sig string) (err error) {
	return r.verifySignature(RSA_CLASSES_PKCS8, crypto.MD5, data, pubKey, sig)
}

//VerifySha1WithPKCS1 sha1和pkcs1 验签
func (r *defaultRsa) VerifySha1WithPKCS1(data, pubKey []byte, sig string) (err error) {
	return r.verifySignature(RSA_CLASSES_PKCS1, crypto.SHA1, data, pubKey, sig)
}

//VerifySha1WithPKCS8 sha1和pkcs1 验签
func (r *defaultRsa) VerifySha1WithPKCS8(data, pubKey []byte, sig string) (err error) {
	return r.verifySignature(RSA_CLASSES_PKCS8, crypto.SHA1, data, pubKey, sig)
}

//VerifySha256WithPKCS1 sha256和pkcs1 验签
func (r *defaultRsa) VerifySha256WithPKCS1(data, pubKey []byte, sig string) (err error) {
	return r.verifySignature(RSA_CLASSES_PKCS1, crypto.SHA256, data, pubKey, sig)
}

//VerifySha256WithPKCS8 sha256和pkcs8 解密
func (r *defaultRsa) VerifySha256WithPKCS8(data, pubKey []byte, sig string) (err error) {
	return r.verifySignature(RSA_CLASSES_PKCS8, crypto.SHA256, data, pubKey, sig)
}

// verifySignature 解签
func (r *defaultRsa) verifySignature(classes int, hash crypto.Hash, data []byte, pub []byte, sig string) (err error) {
	signature, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return
	}

	pubKey, err := r.publicKey(classes, pub)
	if err != nil {
		return
	}

	hashed, err := r.hashed(hash, data)
	if err != nil {
		return
	}

	return rsa.VerifyPKCS1v15(pubKey, hash, hashed, signature)
}

//SignMd5WithPKCS1 md5和pkcs1 加签
func (r *defaultRsa) SignMd5WithPKCS1(data, priKey []byte) (sign string, err error) {
	return r.addSignature(RSA_CLASSES_PKCS1, crypto.MD5, data, priKey)
}

//SignMd5WithPKCS8 md5和pkcs8 加签
func (r *defaultRsa) SignMd5WithPKCS8(data, priKey []byte) (sign string, err error) {
	return r.addSignature(RSA_CLASSES_PKCS8, crypto.MD5, data, priKey)
}

//SignSha1WithPKCS1 sha1和pkcs1 加签
func (r *defaultRsa) SignSha1WithPKCS1(data, priKey []byte) (sign string, err error) {
	return r.addSignature(RSA_CLASSES_PKCS1, crypto.SHA1, data, priKey)
}

//SignSha1WithPKCS8 sha1和pkcs1 加签
func (r *defaultRsa) SignSha1WithPKCS8(data, priKey []byte) (sign string, err error) {
	return r.addSignature(RSA_CLASSES_PKCS8, crypto.SHA1, data, priKey)
}

//SignSha256WithPKCS1 sha256和pkcs1 加签
func (r *defaultRsa) SignSha256WithPKCS1(data, priKey []byte) (sign string, err error) {
	return r.addSignature(RSA_CLASSES_PKCS1, crypto.SHA256, data, priKey)
}

//SignSha256WithPKCS8 sha256和pkcs8 加签
func (r *defaultRsa) SignSha256WithPKCS8(data, priKey []byte) (sign string, err error) {
	return r.addSignature(RSA_CLASSES_PKCS8, crypto.SHA256, data, priKey)
}

//addSignature 加签
func (r *defaultRsa) addSignature(classes int, hash crypto.Hash, data, priKey []byte) (sign string, err error) {
	pri, err := r.privateKey(classes, priKey)
	if err != nil {
		return
	}

	// defaultHash
	hashed, err := r.hashed(hash, data)
	if err != nil {
		return
	}

	b, err := rsa.SignPKCS1v15(rand.Reader, pri, hash, hashed)
	if err != nil {
		return
	}

	return base64.StdEncoding.EncodeToString(b), err
}
