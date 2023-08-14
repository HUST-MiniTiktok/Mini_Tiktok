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

func GetUserById(ctx context.Context, id int64) (user *User, err error) {
	user = new(User)
	err = DB.WithContext(ctx).First(user, id).Error
	return
}

func GetUserByUserName(ctx context.Context, userName string) (user *User, err error) {
	user = new(User)
	err = DB.WithContext(ctx).Where("user_name = ?", userName).First(user).Error
	return
}

func UpdateUser(ctx context.Context, user *User) (err error) {
	err = DB.WithContext(ctx).Save(user).Error
	return
}
