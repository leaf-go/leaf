package models

import (
	"gorm.io/gorm"
	"x"
)

var (
	db = x.MysqlDB
)

type modeler interface {
	IdGet()
}

type model struct {
	tableName string
}

func (m model) Table() *gorm.DB {
	return db.Table(m.tableName)
}

func (m model) TableName() string {
	return m.tableName
}

func (m *model) IdGet(id int) error {
	return m.Table().Where("id = ? AND state = 1", id).First(&m).Error
}
