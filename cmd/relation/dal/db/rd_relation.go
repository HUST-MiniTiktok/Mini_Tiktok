package db

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	FollowSuffix   = ":follow"
	FollowerSuffix = ":follower"
	RDCacheExpire  = time.Hour
)

func RDAddFollow(user_id int64, follower_id int64) {
	RDClient.SAdd(strconv.FormatInt(follower_id, 10)+FollowSuffix, user_id, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}

func RDAddFollower(user_id int64, follower_id int64) {
	RDClient.SAdd(strconv.FormatInt(user_id, 10)+FollowerSuffix, follower_id, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
}

func RDDelFollow(user_id int64, follower_id int64) {
	key := strconv.FormatInt(follower_id, 10) + FollowSuffix
	if RDClient.SExistKey(key) {
		RDClient.SRem(key, user_id)
	}
}

func RDDelFollower(user_id int64, follower_id int64) {
	key := strconv.FormatInt(user_id, 10) + FollowerSuffix
	if RDClient.SExistKey(key) {
		RDClient.SRem(key, follower_id)
	}
}

func RDExistFollowKey(follower_id int64) bool {
	key := strconv.FormatInt(follower_id, 10) + FollowSuffix
	RDClient.SExpire(key, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
	return RDClient.SExistKey(key)
}

func RDExistFollowerKey(user_id int64) bool {
	key := strconv.FormatInt(user_id, 10) + FollowerSuffix
	RDClient.SExpire(key, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
	return RDClient.SExistKey(key)
}

func RDExistFollowValue(user_id int64, follower_id int64) bool {
	key := strconv.FormatInt(follower_id, 10) + FollowSuffix
	RDClient.SExpire(key, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
	return RDClient.SExistValue(key, user_id)
}

func RDExistFollowerValue(user_id int64, follower_id int64) bool {
	key := strconv.FormatInt(user_id, 10) + FollowerSuffix
	RDClient.SExpire(key, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute)
	return RDClient.SExistValue(key, follower_id)
}

func RDCountFollow(follower_id int64) int64 {
	return RDClient.SCount(strconv.FormatInt(follower_id, 10) + FollowSuffix)
}

func RDCountFollower(user_id int64) int64 {
	return RDClient.SCount(strconv.FormatInt(user_id, 10) + FollowerSuffix)
}

func RDCountFriend(user_id int64) int64 {
	return RDClient.SCountInter(strconv.FormatInt(user_id, 10)+FollowSuffix, strconv.FormatInt(user_id, 10)+FollowerSuffix)
}

func RDGetFollowList(follower_id int64) []int64 {
	return RDClient.SGetIntSlice(strconv.FormatInt(follower_id, 10) + FollowSuffix)
}

func RDGetFollowerList(user_id int64) []int64 {
	return RDClient.SGetIntSlice(strconv.FormatInt(user_id, 10) + FollowerSuffix)
}

func RDGetFriendList(user_id int64) []int64 {
	return RDClient.SGetInterIntSlice(strconv.FormatInt(user_id, 10)+FollowSuffix, strconv.FormatInt(user_id, 10)+FollowerSuffix)
}
