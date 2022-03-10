package mounts

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path"
	"runtime"
	
)

type Configs struct {
	Keys  Keys         `toml:"keys"`
	Jwt   JwtConfig    `toml:"jwt"`
	Log   x.LogConfig  `toml:"log"`
	Http  x.HttpConfig `toml:"http"`
	Mysql x.Mysql      `toml:"mysql"`
	Redis x.Redis      `toml:"redis"`
}

func (c Configs) Initialize() {
	c.Mysql.Init()
	c.Redis.Init()
	c.Log.Init()
	c.Jwt.Init()
}

func (c *Configs) Parse(dir string) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("current file unknown error")
	}

	baseDir := path.Dir(path.Dir(filename))
	file := fmt.Sprintf("%s/%s/%s.toml", baseDir, dir, x.Env())
	_, err := toml.DecodeFile(file, &c)
	return err
}

type Keys struct {
	PublicKey  string `toml:"public_key"`
	PrivateKey string `toml:"private_key"`
	SigKey     string `toml:"sig_key"`
	AesKey     string `toml:"aes_key"`
	IV         string `toml:"iv"`
}
