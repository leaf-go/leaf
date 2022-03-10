package models

import (
	"leaf/utils"
	"time"
	
)

var (
	db = x.MysqlDB
	aesKey  = []byte("7qRwF91vC2#W37e8")
	isCrypt = x.IsRelease()
)


type Model struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at;autoUpdateTime"`
}

func (m Model) EncodeField(content string) (result string, err error) {
	if !isCrypt {
		return content, nil
	}
	return utils.AES.Encrypt([]byte(content), aesKey)
}

func (m Model) DecodeField(content string) (result string, err error) {
	if !isCrypt {
		return content, nil
	}
	return utils.AES.Decrypt([]byte(content), aesKey)
}
