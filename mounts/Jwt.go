package mounts

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
)

const (
	KeyKindPublic KeyKind = iota + 1
	KeyKindPrivate
)

var (
	jwtConfig  *JwtConfig
	defaultJwt *Jwt
)

type KeyKind int

type HttpApIClaims struct {
	jwt.StandardClaims
	Mobile     string `json:"mobile,omitempty"`
	Token      string `json:"token,omitempty"`
	RealNameId int    `json:"rni,omitempty"`
	Ext        string `json:"ext,omitempty"`
}

func SingletonJWT() *Jwt {
	if defaultJwt == nil {
		defaultJwt = &Jwt{}
	}

	return defaultJwt
}


type UserClaims struct {
	jwt.StandardClaims
	Mobile string `json:"mobile"`
	Token  string `json:"token"`
}

type Jwt struct {
}

func (j Jwt) Generate(claims jwt.Claims) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(jwtConfig.parsedPrivateKey)
}

func (j Jwt) Valid(tokenString string, claims jwt.Claims) (err error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtConfig.parsedPublicKey, nil
	})

	if token == nil {
		return errors.New("token is invalid")
	}

	return
}

type JwtConfig struct {
	PublicKey  string `toml:"public_key"`
	PrivateKey string `toml:"private_key"`
	Issuer     string `toml:"issuer"`
	ExpiresAt  int    `toml:"expires_at"`

	parsedPublicKey  *ecdsa.PublicKey  `toml:"-"`
	parsedPrivateKey *ecdsa.PrivateKey `toml:"-"`
}

func (c *JwtConfig) Init() {
	var err error
	if err = c.publicKey(); err != nil {
		panic(err)
	}

	if err = c.privateKey(); err != nil {
		panic(err)
	}

	jwtConfig = c
}

func (c JwtConfig) Keys() (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	return c.parsedPublicKey, c.parsedPrivateKey
}

func (c *JwtConfig) publicKey() (err error) {
	key := c.formatKey(KeyKindPublic, c.PublicKey)
	c.parsedPublicKey, err = jwt.ParseECPublicKeyFromPEM(key)
	return
}

func (c *JwtConfig) privateKey() (err error) {
	key := c.formatKey(KeyKindPrivate, c.PrivateKey)
	c.parsedPrivateKey, err = jwt.ParseECPrivateKeyFromPEM(key)
	return
}

func (c JwtConfig) formatKey(kind KeyKind, key string) []byte {
	var wrap string
	if kind == KeyKindPrivate {
		wrap = "RSA PRIVATE"
	} else {
		wrap = "PUBLIC"
	}

	format := fmt.Sprintf("-----BEGIN %s KEY-----\n%s\n-----END %s KEY-----", wrap, key, wrap)
	return []byte(format)
}
