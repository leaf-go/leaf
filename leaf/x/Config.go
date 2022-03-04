package x

import (
	"github.com/BurntSushi/toml"
	"os"
)

type TomlConfig struct {
	Path string
}

func NewTomlConfig(path string) *TomlConfig {
	t := &TomlConfig{Path: path}
	if exists, err := t.fileExists(); !exists {
		if err == nil {
			panic("toml config file not exists")
		}

		panic(err)
	}

	return t
}

func (c *TomlConfig) Parse(config interface{}) (err error) {
	_, err = toml.DecodeFile(c.Path, &config)
	return
}

func (c *TomlConfig) fileExists() (bool, error) {
	// 获取文件信息
	_, err := os.Stat(c.Path)
	if err == nil {
		return true, nil
	}

	// 如果出现错误, 判断错误是不是文件不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
