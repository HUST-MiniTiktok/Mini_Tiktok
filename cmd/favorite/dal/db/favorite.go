package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const FavoriteTableName = "favorite"

type Favorite struct {
	ID        int64          `json:"id"`
	UserId    int64          `json:"user_id"`
	VideoId   int64          `json:"video_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (Favorite) TableName() string {
	return FavoriteTableName
}

// status 返回0——成功，返回1——失败
// err 返回nil——成功，返回其他——失败原因
func NewFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {

	// 创建一条favorite数据
	favorite := Favorite{
		UserId:  user_id,
		VideoId: video_id}

	//新建喜欢、新增喜欢为同一事务
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.WithContext(ctx).Model(&Favorite{}).Where(&favorite).Count(&count).Error; err != nil {
			return err
		}
		if count != 0 { // 重复点赞
			return nil
		}

		if err := tx.WithContext(ctx).Model(&Favorite{}).Create(&favorite).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 1, err
	}

	return 0, nil
}

func CancelFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {
	//先根据user_id和video_id寻找到id，再根据id软删除
	var favorite Favorite

	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//通过user_id和video_id找到要删除的favorite记录
		if err := tx.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{UserId: user_id, VideoId: video_id}).Scan(&favorite).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&Favorite{}).Delete(&favorite).Error; err != nil { //删除这条favorite记录
			return err
		}
		return nil
	})

	if err != nil {
		return 1, err
	}

	return 0, nil
}

func CheckFavorite(ctx context.Context, user_id int64, video_id int64) (status bool, err error) {

	var count int64
	err = DB.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{UserId: user_id, VideoId: video_id}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func VideoFavoriteCount(ctx context.Context, video_id int64) (count int64, err error) {

	err = DB.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{VideoId: video_id}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UserFavoriteCount(ctx context.Context, user_id int64) (count int64, err error) {

	err = DB.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{UserId: user_id}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetFavoriteList(ctx context.Context, user_id int64) (status int32, videoIDList []int64, err error) { //仅返回videoID

	var favoriteList []Favorite
	//根据user_id，在favorite表中找到他的所有favorite记录
	if err = DB.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{UserId: user_id}).Find(&favoriteList).Error; err != nil {
		return 1, nil, err
	}

	for _, favoritelog := range favoriteList {
		videoIDList = append(videoIDList, favoritelog.VideoId)
	}

	return 0, videoIDList, err
}
