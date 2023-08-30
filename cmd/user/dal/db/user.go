package db

import (
	"context"
)

const UserTableName = "user"

type User struct {
	ID              int64  `json:"id" gorm:"primaryKey;autoincrement"`
	UserName        string `json:"user_name" gorm:"uniqueIndex:user_name_idx"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return UserTableName
}

// CreateUser: create a new user record
func CreateUser(ctx context.Context, user *User) (id int64, err error) {
	err = DB.WithContext(ctx).Create(user).Error
	id = int64(user.ID)
	return
}

// GetUserById: get a user by user id
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

// GetUserByUserName: get a user by user name
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

// UpdateUser: update a user record
func UpdateUser(ctx context.Context, user *User) (err error) {
	err = DB.WithContext(ctx).Save(user).Error
	return
}

// CheckUserById: check if a user exists by user id
func CheckUserById(ctx context.Context, userId int64) (exist bool, err error) {
	var user User
	err = DB.Where("id = ?", userId).Limit(1).Find(&user).Error
	if err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
