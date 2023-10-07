package db

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	VideoFavoriteCountField = ":videoFavoriteCount"
	UserFavoriteCountField  = ":userFavoriteCount"
	RDCacheExpire           = time.Hour
)

func RDExistVideoFavoriteCount(video_id int64) bool {
	video_id_str := strconv.FormatInt(video_id, 10)
	if RDClient.HExists(video_id_str, VideoFavoriteCountField) {
		RDClient.HExpire(video_id_str, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
		return true
	}
	return false
}

func RDExistUserFavoriteCount(user_id int64) bool {
	user_id_str := strconv.FormatInt(user_id, 10)
	if RDClient.HExists(user_id_str, UserFavoriteCountField) {
		RDClient.HExpire(user_id_str, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
		return true
	}
	return false
}

func RDIncrVideoFavoriteCount(video_id int64, val int64) {
	video_id_str := strconv.FormatInt(video_id, 10)
	if RDExistVideoFavoriteCount(video_id) {
		RDClient.HIncr(video_id_str, VideoFavoriteCountField, val)
	}
}

func RDIncrUserFavoriteCount(user_id int64, val int64) {
	user_id_str := strconv.FormatInt(user_id, 10)
	if RDExistUserFavoriteCount(user_id) {
		RDClient.HIncr(user_id_str, UserFavoriteCountField, val)
	}
}

func RDGetVideoFavoriteCount(video_id int64) int64 {
	video_id_str := strconv.FormatInt(video_id, 10)
	return RDClient.HGetInt(video_id_str, VideoFavoriteCountField)
}

func RDGetUserFavoriteCount(user_id int64) int64 {
	user_id_str := strconv.FormatInt(user_id, 10)
	return RDClient.HGetInt(user_id_str, UserFavoriteCountField)
}

func RDSetVideoFavoriteCount(video_id int64, val int64) {
	video_id_str := strconv.FormatInt(video_id, 10)
	RDClient.HSet(video_id_str, VideoFavoriteCountField, val, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}

func RDSetUserFavoriteCount(user_id int64, val int64) {
	user_id_str := strconv.FormatInt(user_id, 10)
	RDClient.HSet(user_id_str, UserFavoriteCountField, val, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}
