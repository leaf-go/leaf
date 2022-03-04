package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
)

var (
	File     = newDefaultFile()
	APP_PATH string

	FileExistsError    = errors.New("file already exists")
	FileNotExistsError = errors.New("file not exists")
)

type defaultFile struct{}

func newDefaultFile() *defaultFile {
	return &defaultFile{}
}

func (f *defaultFile) LoadNet(url string) (file []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	return ioutil.ReadAll(resp.Body)
}

// AppPath ./../../
func (f defaultFile) AppPath() string {
	if 0 == len(APP_PATH) {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("current file unknown error")
		}

		APP_PATH = path.Dir(path.Dir(filename))
	}
	return APP_PATH
}

// Load 加载内容
func (f *defaultFile) Load(filepath string) (content []byte, err error) {
	return ioutil.ReadFile(filepath)
}

// Exists 判断文件是否存在
func (f *defaultFile) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// Create 创建文件
func (f defaultFile) Create(filepath string, content []byte) (err error) {
	exists, _ := f.Exists(filepath)
	if exists {
		return FileExistsError
	}

	file, err := os.Create(filepath)
	if err != nil {
		return
	}

	defer file.Close()
	return ioutil.WriteFile(filepath, content, 0644)
}

func (f defaultFile) Append(filepath string, content []byte) (err error) {
	exists, _ := f.Exists(filepath)
	if !exists {
		return FileNotExistsError
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	n, _ := file.Seek(0, 2)
	_, err = file.WriteAt(content, n)
	return
}

func (f defaultFile) Delete(filepath string) (err error) {
	return os.Remove(filepath)
}
