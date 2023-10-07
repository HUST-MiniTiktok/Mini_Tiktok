package db

import (
	"context"
	"strconv"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

const FollowTableName = "follow"

type Follow struct {
	ID         int64          `json:"id" gorm:"primaryKey;autoincrement"`
	UserId     int64          `json:"user_id" gorm:"index:follow_idx;index:follow_user_idx"`
	FollowerId int64          `json:"follower_id" gorm:"index:follow_idx;index:follow_follower_idx"`
	CreatedAt  time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (Follow) TableName() string {
	return FollowTableName
}

// CreateFollow: create a new follow record
func CreateFollow(ctx context.Context, follow *Follow) (id int64, err error) {
	// 如果关注关系已经存在，那么就不需要再次创建
	ok, err := CheckFollow(ctx, follow.UserId, follow.FollowerId)
	if err != nil {
		return -1, err
	}
	if ok {
		return 0, nil
	}
	err = DB.WithContext(ctx).Create(follow).Error
	if err != nil {
		return -1, err
	}

	go RDAddFollow(follow.UserId, follow.FollowerId)
	go RDAddFollower(follow.UserId, follow.FollowerId)

	go Filter.AddToBloomFilter(strconv.Itoa(int(follow.UserId)))
	go Filter.AddToBloomFilter(strconv.Itoa(int(follow.FollowerId)))

	return follow.ID, nil
}

// DeleteFollow: delete a follow record
func DeleteFollow(ctx context.Context, follow *Follow) (ok bool, err error) {
	err = DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", follow.UserId, follow.FollowerId).Delete(follow).Error
	if err != nil {
		return false, err
	}

	go RDDelFollow(follow.UserId, follow.FollowerId)
	go RDDelFollower(follow.UserId, follow.FollowerId)

	return true, nil
}

// CheckFollow: check if a follow record exists
func CheckFollow(ctx context.Context, user_id int64, follower_id int64) (ok bool, err error) {
	if RDExistFollowKey(follower_id) {
		return RDExistFollowValue(user_id, follower_id), nil
	} else if RDExistFollowerKey(user_id) {
		return RDExistFollowerValue(user_id, follower_id), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(user_id, FollowSuffix) {
		defer RDClient.ReleaseLock(user_id, FollowSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(user_id)))
		if !exist { // user not exist
			return false, nil
		}

		exist = Filter.TestBloom(strconv.Itoa(int(follower_id)))
		if !exist { // user not exist
			return false, nil
		}

		var db_follow Follow
		err = DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", user_id, follower_id).Limit(1).Find(&db_follow).Error
		if err != nil {
			return false, err
		}
		result := db_follow != (Follow{})
		if result {
			go func() {
				if !RDExistFollowKey(follower_id) || !RDExistFollowValue(user_id, follower_id) {
					RDAddFollow(user_id, follower_id)
				}
				if !RDExistFollowerKey(user_id) || !RDExistFollowerValue(user_id, follower_id) {
					RDAddFollower(user_id, follower_id)
				}
			}()
		}
		return result, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return CheckFollow(ctx, user_id, follower_id)
}

// GetFollowById: get a follow user id list by current user id
func GetFollowUserIdList(ctx context.Context, userId int64) (user_ids []int64, err error) {
	if RDExistFollowKey(userId) {
		return RDGetFollowList(userId), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(userId, FollowSuffix) {
		defer RDClient.ReleaseLock(userId, FollowSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(userId)))
		if !exist { // user not exist
			return nil, nil
		}

		var follows []Follow
		err = DB.WithContext(ctx).Where("follower_id = ?", userId).Order("user_id asc").Find(&follows).Error
		if err != nil {
			return nil, err
		}
		for _, follow := range follows {
			user_ids = append(user_ids, follow.UserId)
		}
		return user_ids, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return GetFollowUserIdList(ctx, userId)
}

// GetFollowerUserIdList: get a follower user id list by current user id
func GetFollowerUserIdList(ctx context.Context, userId int64) (user_ids []int64, err error) {
	if RDExistFollowerKey(userId) {
		return RDGetFollowerList(userId), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(userId, FollowerSuffix) {
		defer RDClient.ReleaseLock(userId, FollowerSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(userId)))
		if !exist { // user not exist
			return nil, nil
		}

		var follows []Follow
		err = DB.WithContext(ctx).Where("user_id = ?", userId).Order("user_id asc").Find(&follows).Error
		if err != nil {
			return nil, err
		}
		for _, follow := range follows {
			user_ids = append(user_ids, follow.FollowerId)
		}
		return user_ids, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return GetFollowerUserIdList(ctx, userId)
}

// GetFriendUserIdList: get a friend user id list by current user id
func GetFriendUserIdList(ctx context.Context, userId int64) (user_ids []int64, err error) {
	// Friend 要求 A关注B 且 B关注A
	// user_id 表示被关注者A, follower_id 表示关注者B
	if RDExistFollowKey(userId) && RDExistFollowerKey(userId) {
		return RDGetFriendList(userId), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(userId, FollowSuffix+FollowerSuffix) {
		defer RDClient.ReleaseLock(userId, FollowSuffix+FollowerSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(userId)))
		if !exist { // user not exist
			return nil, nil
		}

		var follows []Follow
		err = DB.WithContext(ctx).Where("user_id = ?", userId).Where("follower_id IN (SELECT user_id FROM follow WHERE follower_id = ?)", userId).Order("user_id asc").Find(&follows).Error
		if err != nil {
			return nil, err
		}
		for _, follow := range follows {
			user_ids = append(user_ids, follow.FollowerId)
		}
		return user_ids, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return GetFriendUserIdList(ctx, userId)
}

// GetFollowUserCount: get a follow user count number of current user id
func GetFollowUserCount(ctx context.Context, userId int64) (count int64, err error) {
	if RDExistFollowKey(userId) {
		return RDCountFollow(userId), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(userId, FollowSuffix) {
		defer RDClient.ReleaseLock(userId, FollowSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(userId)))
		if !exist { // user not exist
			return 0, nil
		}

		err = DB.WithContext(ctx).Model(&Follow{}).Where("follower_id = ?", userId).Count(&count).Error
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return GetFollowUserCount(ctx, userId)
}

// GetFollowerUserCount: get a follower user count number of current user id
func GetFollowerUserCount(ctx context.Context, userId int64) (count int64, err error) {
	if RDExistFollowerKey(userId) {
		return RDCountFollower(userId), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(userId, FollowSuffix+FollowerSuffix) {
		defer RDClient.ReleaseLock(userId, FollowSuffix+FollowerSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(userId)))
		if !exist { // user not exist
			return 0, nil
		}

		err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ?", userId).Count(&count).Error
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return GetFollowerUserCount(ctx, userId)
}

// GetFriendUserCount: get a friend user count number of current user id
func GetFriendUserCount(ctx context.Context, userId int64) (count int64, err error) {
	// Friend 要求 A关注B 且 B关注A
	// user_id 表示被关注者A, follower_id 表示关注者B
	if RDExistFollowKey(userId) && RDExistFollowerKey(userId) {
		return RDCountFriend(userId), nil
	}
	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(userId, FollowSuffix+FollowerSuffix) {
		defer RDClient.ReleaseLock(userId, FollowSuffix+FollowerSuffix)

		exist := Filter.TestBloom(strconv.Itoa(int(userId)))
		if !exist { // user not exist
			return 0, nil
		}

		err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ?", userId).Where("follower_id IN (SELECT user_id FROM follow WHERE follower_id = ?)", userId).Count(&count).Error
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	time.Sleep(redis.RetryTime) // delay and retry
	return GetFriendUserCount(ctx, userId)
}

func LoadRelationUserIDToBloomFilter(ctx context.Context) error {
	var userIdList []string

	err := DB.WithContext(ctx).Model(&Follow{}).Pluck("UserId", &userIdList).Error
	if err != nil {
		klog.Errorf("Load Relation UserID To BloomFilter Failed: %v", err)
		return err
	}

	for _, user_id := range userIdList {
		Filter.AddToBloomFilter(user_id)
	}
	return nil
}

func LoadRelationFollowerIDToBloomFilter(ctx context.Context) error {
	var userIdList []string

	err := DB.WithContext(ctx).Model(&Follow{}).Pluck("FollowerId", &userIdList).Error
	if err != nil {
		klog.Errorf("Load Relation FollowerID To BloomFilter Failed: %v", err)
		return err
	}

	for _, user_id := range userIdList {
		Filter.AddToBloomFilter(user_id)
	}
	return nil
}
