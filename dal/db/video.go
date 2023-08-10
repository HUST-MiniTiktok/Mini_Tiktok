package db

import (
	"context"
	model "github.com/HUST-MiniTiktok/mini_tiktok/dal/db/model"
)

func GetVideoFeed(ctx context.Context, latest_time int64) ([]*model.Video, error) {
	result := make([]*model.Video, 0)
	if err := DB.WithContext(ctx).
		Where("created_at < ?", latest_time).
		Order("created_at desc").
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}

//TODO：其他的数据库操作，比如查询、更新等