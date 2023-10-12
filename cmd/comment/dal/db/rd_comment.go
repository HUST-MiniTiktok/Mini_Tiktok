package db

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	CommentCountField = ":commentCount"
	RDCacheExpire     = time.Hour
)

func RDExistCommentCount(video_id int64) bool {
	video_id_str := strconv.FormatInt(video_id, 10)
	if RDClient.HExists(video_id_str, CommentCountField) {
		RDClient.HExpire(video_id_str, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
		return true
	}
	return false
}

func RDIncrCommentCount(video_id int64, val int64) {
	video_id_str := strconv.FormatInt(video_id, 10)
	if RDExistCommentCount(video_id) {
		RDClient.HIncr(video_id_str, CommentCountField, val)
	}
}

func RDSetCommentCount(video_id int64, val int64) {
	video_id_str := strconv.FormatInt(video_id, 10)
	RDClient.HSet(video_id_str, CommentCountField, val, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}

func RDGetCommentCount(video_id int64) int64 {
	video_id_str := strconv.FormatInt(video_id, 10)
	return RDClient.HGetInt(video_id_str, CommentCountField)
}
