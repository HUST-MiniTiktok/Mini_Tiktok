package db

import  (
	"context"
	model "github.com/HUST-MiniTiktok/mini_tiktok/dal/db/model"
)

func CreateUser(ctx context.Context, user *model.User) (id int64, err error) {
	err = DB.WithContext(ctx).Create(user).Error
	id = int64(user.ID)
	return
}

//TODO：其他的数据库操作，比如查询、更新等
