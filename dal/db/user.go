package db

import (
	"context"
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

func CreateUser(ctx context.Context, user *User) (id int64, err error) {
	err = DB.WithContext(ctx).Create(user).Error
	id = int64(user.ID)
	return
}

//TODO：其他的数据库操作，比如查询、更新等
