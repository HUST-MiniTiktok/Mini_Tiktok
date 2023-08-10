package db

import (
	"gorm.io/gorm"
)

const UserTableName = "user"

type User struct {
	gorm.Model
	UserName        string `json:"username"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return UserTableName
}