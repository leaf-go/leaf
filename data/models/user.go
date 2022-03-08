package models

type User struct {
	model
}

func NewUser() *User {
	return &User{model{tableName: "user"}}
}



