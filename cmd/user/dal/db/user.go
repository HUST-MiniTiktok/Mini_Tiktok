package db

import (
	"context"
)

const UserTableName = "user"

type User struct {
	ID              int64  `json:"id"`
	UserName        string `json:"user_name"`
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
	var db_user User
	err = DB.WithContext(ctx).Where("id = ?", id).Find(&db_user).Error
	if err != nil {
		return nil, err
	}
	if db_user.ID == 0 {
		return nil, nil
	}
	user = &db_user
	return user, nil
}

func GetUserByUserName(ctx context.Context, userName string) (user *User, err error) {
	users := make([]*User, 0)
	err = DB.WithContext(ctx).Where("user_name = ?", userName).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	user = users[0]
	return
}

func UpdateUser(ctx context.Context, user *User) (err error) {
	err = DB.WithContext(ctx).Save(user).Error
	return
}

func CheckUserById(ctx context.Context, userId int64) (exist bool, err error) {
	var user User
	err = DB.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
