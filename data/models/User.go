package models

import (
	"gorm.io/gorm"
	"time"
)

func NewUser() *User {
	return &User{}
}

type Users []*User

func (u Users) Ids() []int {
	ids := make([]int, len(u))
	for i, v := range u {
		ids[i] = v.ID
	}

	return ids
}

type User struct {
	Model
	ID           int       `json:"id" gorm:"column:id"`
	Mobile       string    `json:"mobile" gorm:"column:mobile"`
	Name         string    `json:"name" gorm:"column:name"`
	Password     string    `json:"password" gorm:"column:password"`
	Token        string    `json:"token" gorm:"column:token"`
	Os           int       `json:"os" gorm:"column:os"`
	Did          string    `json:"did" gorm:"column:did"`
	RegisterAt   int       `json:"register_at" gorm:"column:register_at"`
	LastActiveAt int       `json:"last_active_at" gorm:"column:last_active_at"`
	Status       bool      `json:"status" gorm:"column:status"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	LastLoginIp  string    `json:"last_login_ip" gorm:"column:last_login_ip"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Table() *gorm.DB {
	return db.Table(u.TableName())
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}

func (u *User) IdGet(id int) error {
	return db.Where("id = ? AND status = 1", id).First(&u).Error
}

func (u *User) Create(mobile string, name string, password string, token string, os int, did string, registerAt int, lastActiveAt int, status bool, lastLoginIp string) error {
	u.Mobile = mobile
	u.Name = name
	u.Password = password
	u.Token = token
	u.Os = os
	u.Did = did
	u.RegisterAt = registerAt
	u.LastActiveAt = lastActiveAt
	u.Status = status
	u.LastLoginIp = lastLoginIp

	return db.Create(&u).Error
}

func (u *User) Edit(mobile string, name string, password string, token string, os int, did string, registerAt int, lastActiveAt int, status bool, lastLoginIp string) error {
	u.Mobile = mobile
	u.Name = name
	u.Password = password
	u.Token = token
	u.Os = os
	u.Did = did
	u.RegisterAt = registerAt
	u.LastActiveAt = lastActiveAt
	u.Status = status
	u.LastLoginIp = lastLoginIp

	return db.Save(&u).Error
}

func (u *User) Save() error {
	return db.Save(&u).Error
}
