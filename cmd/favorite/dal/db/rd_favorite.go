package db

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	VFavoriteCount          = "favoriteCount"
	UserFavoriteCountSuffix = ":Ufavorite"
	RDCacheExpire           = time.Hour
)

func RDExistVideoFavoriteCount(video_id int64) bool {
	video_id_str := strconv.FormatInt(video_id, 10)
	if RDClient.HExists(video_id_str, VFavoriteCount) {
		RDClient.HExpire(video_id_str, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
		return true
	}
	return false
}

func RDExistUserFavoriteCount(user_id int64) bool {
	user_id_str := strconv.FormatInt(user_id, 10) + UserFavoriteCountSuffix
	if RDClient.Exists(user_id_str) {
		RDClient.Expire(user_id_str, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
		return true
	}
	return false
}

func RDIncrVideoFavoriteCount(video_id int64, val int64) {
	video_id_str := strconv.FormatInt(video_id, 10)
	if RDExistVideoFavoriteCount(video_id) {
		RDClient.HIncr(video_id_str, VFavoriteCount, val)
	}
}

func RDIncrUserFavoriteCount(user_id int64, val int64) {
	user_id_str := strconv.FormatInt(user_id, 10) + UserFavoriteCountSuffix
	if RDExistUserFavoriteCount(user_id) {
		RDClient.IncrBy(user_id_str, val)
	}
}

func RDGetVideoFavoriteCount(video_id int64) int64 {
	video_id_str := strconv.FormatInt(video_id, 10)
	return RDClient.HGetInt(video_id_str, VFavoriteCount)
}

func RDGetUserFavoriteCount(user_id int64) int64 {
	user_id_str := strconv.FormatInt(user_id, 10) + UserFavoriteCountSuffix
	return RDClient.GetInt(user_id_str)
}

func RDSetVideoFavoriteCount(video_id int64, val int64) {
	video_id_str := strconv.FormatInt(video_id, 10)
	RDClient.HSet(video_id_str, VFavoriteCount, val, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}

func RDSetUserFavoriteCount(user_id int64, val int64) {
	user_id_str := strconv.FormatInt(user_id, 10) + UserFavoriteCountSuffix
	RDClient.Set(user_id_str, val, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}
