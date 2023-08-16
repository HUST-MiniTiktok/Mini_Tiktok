package db

import (
	"context"

	"gorm.io/gorm"
)

// status 返回0——成功，返回1——失败
// err 返回nil——成功，返回其他——失败原因
func NewFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {

	// 创建一条favorite数据
	favorite := Favorite{
		//TODO:ID这里不是逐主键
		UserId:  user_id,
		VideoId: video_id}

	//新建喜欢、新增喜欢为同一事务
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favorite).Error; err != nil {
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
		if err := tx.Select("*").First(&Favorite{UserId: user_id, VideoId: video_id}).Scan(&favorite).Error; err != nil {
			return err
		}
		//删除这条favorite记录
		if err := tx.Delete(&favorite).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 1, err
	}

	return 0, nil
}

// func GetFavoriteList(user_id int64) (status int32, videoList []Video, err error) { //返回video数组
// 	var favoriteList []Favorite

// 	//根据user_id，在favorite表中找到他的所有favorite记录
// 	if err = DB.Select("video_id").Where(&Favorite{UserId: user_id}).Find(&favoriteList).Error; err != nil {
// 		return 1, nil, err
// 	}
// 	//根据favorite记录找到所有对应video_id的video
// 	var video Video
// 	for _, favoritelog := range favoriteList {
// 		if err = DB.Where(&Video{ID: favoritelog.VideoId}).First(&video).Error; err != nil {
// 			return 1, videoList, errors.New("failed finding all the favorite videos")
// 		}
// 		videoList = append(videoList, video)
// 	}

//		return 0, videoList, err
//	}
func GetFavoriteList(ctx context.Context, user_id int64) (status int32, videoIDList []int64, err error) { //仅返回videoID
	var favoriteList []Favorite

	//根据user_id，在favorite表中找到他的所有favorite记录
	if err = DB.WithContext(ctx).Select("video_id").Where(&Favorite{UserId: user_id}).Find(&favoriteList).Error; err != nil {
		return 1, nil, err
	}

	for _, favoritelog := range favoriteList {
		videoIDList = append(videoIDList, favoritelog.VideoId)
	}

	return 0, videoIDList, err
}
