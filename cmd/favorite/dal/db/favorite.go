package db

import (
	"context"
	"strconv"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

const (
	FavoriteTableName = "favorite"
)

type Favorite struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoincrement"`
	UserId    int64          `json:"user_id" gorm:"index:favorite_idx;index:favorite_user_idx"`
	VideoId   int64          `json:"video_id" gorm:"index:favorite_idx;index:favorite_video_idx"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (Favorite) TableName() string {
	return FavoriteTableName
}

// NewFavorite Create a new favorite record
// status 0:success 1:failed
func NewFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {

	favorite := Favorite{
		UserId:  user_id,
		VideoId: video_id}

	isFavorited, err := CheckFavorite(ctx, user_id, video_id)
	if err != nil {
		return 1, err
	}
	if isFavorited { // Repeat favorite
		return 1, nil
	}

	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Model(&Favorite{}).Create(&favorite).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 1, err
	}

	go RDIncrVideoFavoriteCount(video_id, 1)
	go RDIncrUserFavoriteCount(user_id, 1)

	go Filter.AddToBloomFilter(strconv.Itoa(int(video_id)))
	go Filter.AddToBloomFilter(strconv.Itoa(int(user_id)))

	return 0, nil
}

// Cancel a favorite record
func CancelFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {

	err = DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&Favorite{}).Error
	if err != nil {
		return 1, err
	}

	go RDIncrVideoFavoriteCount(video_id, -1)
	go RDIncrUserFavoriteCount(user_id, -1)

	return 0, nil
}

// CheckFavorite check if a user has favorited a video
func CheckFavorite(ctx context.Context, user_id int64, video_id int64) (status bool, err error) {
	var db_favorite Favorite
	err = DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", user_id, video_id).Limit(1).Find(&db_favorite).Error
	if err != nil {
		return false, err
	}
	// klog.Infof("check is favorited: %t", db_favorite != Favorite{})
	return db_favorite != Favorite{}, nil
}

// VideoFavoriteCount get the favorite count of a video
func VideoFavoriteCount(ctx context.Context, video_id int64) (count int64, err error) {

	if RDExistVideoFavoriteCount(video_id) {
		return RDGetVideoFavoriteCount(video_id), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(video_id, VideoFavoriteCountField) {
		defer RDClient.ReleaseLock(video_id, VideoFavoriteCountField)

		exist := Filter.TestBloom(strconv.Itoa(int(video_id)))
		if !exist { // video not exist
			return 0, nil
		}

		err = DB.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{VideoId: video_id}).Count(&count).Error
		if err != nil {
			return 0, err
		}

		go RDSetVideoFavoriteCount(video_id, count)

		return count, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return VideoFavoriteCount(ctx, video_id)
}

// UserFavoriteCount get the favorite count of a user
func UserFavoriteCount(ctx context.Context, user_id int64) (count int64, err error) {

	if RDExistUserFavoriteCount(user_id) {
		return RDGetUserFavoriteCount(user_id), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(user_id, UserFavoriteCountField) {
		defer RDClient.ReleaseLock(user_id, VideoFavoriteCountField)

		exist := Filter.TestBloom(strconv.Itoa(int(user_id)))
		if !exist { // user not exist
			return 0, nil
		}

		err = DB.WithContext(ctx).Model(&Favorite{}).Where(&Favorite{UserId: user_id}).Count(&count).Error
		if err != nil {
			return 0, err
		}

		go RDSetUserFavoriteCount(user_id, count)

		return count, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return UserFavoriteCount(ctx, user_id)
}

// GetFavoriteList get the favorite list of a user
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

func LoadFavoriteVideoIDToBloomFilter(ctx context.Context) error {
	var videoIdList []string

	err := DB.WithContext(ctx).Model(&Favorite{}).Pluck("VideoId", &videoIdList).Error
	if err != nil {
		klog.Errorf("Load Favorite VideoID To BloomFilter Failed: %v", err)
		return err
	}

	for _, video_id := range videoIdList {
		Filter.AddToBloomFilter(video_id)
	}
	return nil
}

func LoadFavoriteUserIDToBloomFilter(ctx context.Context) error {
	var userIdList []string

	err := DB.WithContext(ctx).Model(&Favorite{}).Pluck("UserId", &userIdList).Error
	if err != nil {
		klog.Errorf("Load Favorite UserID To BloomFilter Failed: %v", err)
		return err
	}

	for _, user_id := range userIdList {
		Filter.AddToBloomFilter(user_id)
	}
	return nil
}
