package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

const UserTableName = "user"

type User struct {
	gorm.Model
	UserName        string
	Password        string
	Avatar          string
	BackgroundImage string
	Signature       string
}

func (User) TableName() string {
	return UserTableName
}

func CreateUser(ctx context.Context, user *User) (id int64, err error) {
	err = DB.WithContext(ctx).Create(user).Error
	klog.Warnf("create user: %v", user)
	id = int64(user.ID)
	return
}

func GetUserById(ctx context.Context, id int64) (user *User, err error) {
	klog.Warnf("get user by id: %v", id)
	err = DB.WithContext(ctx).First(&user, id).Error
	return
}

func GetUserByUserName(ctx context.Context, userName string) (user *User, err error) {
	users := make([]*User, 0)
	klog.Warnf("get user by username: %v", userName)
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