package db

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestInit(t *testing.T) {
	Init()
}

func TestCreateFollow(t *testing.T) {
	ctx := context.Background()
	Init()
	newfollow1 := Follow{
		UserId:     1,
		FollowerId: 2,
	}
	newfollow2 := Follow{
		UserId:     2,
		FollowerId: 1,
	}
	newfollow3 := Follow{
		UserId:     1,
		FollowerId: 3,
	}
	newfollow4 := Follow{
		UserId:     1,
		FollowerId: 4,
	}
	_, err := CreateFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow4)
	assert.Nil(t, err)
}

func TestDeleteFollow(t *testing.T) {
	ctx := context.Background()
	Init()
	newfollow1 := Follow{
		UserId:     1,
		FollowerId: 2,
	}
	newfollow2 := Follow{
		UserId:     2,
		FollowerId: 1,
	}
	newfollow3 := Follow{
		UserId:     1,
		FollowerId: 3,
	}
	newfollow4 := Follow{
		UserId:     1,
		FollowerId: 4,
	}
	_, err := CreateFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow4)
	assert.Nil(t, err)
	ok, err := DeleteFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow4)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
}

func TestCheckFollow(t *testing.T) {
	ctx := context.Background()
	Init()
	newfollow1 := Follow{
		UserId:     1,
		FollowerId: 2,
	}
	newfollow2 := Follow{
		UserId:     2,
		FollowerId: 1,
	}
	newfollow3 := Follow{
		UserId:     1,
		FollowerId: 3,
	}
	newfollow4 := Follow{
		UserId:     1,
		FollowerId: 4,
	}
	_, err := CreateFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow4)
	assert.Nil(t, err)
	ok, err := CheckFollow(ctx, 1, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = CheckFollow(ctx, 2, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = CheckFollow(ctx, 1, 3)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = CheckFollow(ctx, 1, 4)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = CheckFollow(ctx, 4, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, false, ok)
	ok, err = CheckFollow(ctx, 3, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, false, ok)
	ok, err = DeleteFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = CheckFollow(ctx, 1, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, false, ok)
	ok, err = CheckFollow(ctx, 2, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, false, ok)
}

func TestGetFollowUserIdList(t *testing.T) {
	ctx := context.Background()
	Init()
	newfollow1 := Follow{
		UserId:     1,
		FollowerId: 2,
	}
	newfollow2 := Follow{
		UserId:     2,
		FollowerId: 1,
	}
	newfollow3 := Follow{
		UserId:     3,
		FollowerId: 1,
	}
	newfollow4 := Follow{
		UserId:     4,
		FollowerId: 1,
	}
	_, err := CreateFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow4)
	assert.Nil(t, err)
	user_ids, err := GetFollowUserIdList(ctx, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{2, 3, 4}, user_ids)
	user_ids, err = GetFollowUserIdList(ctx, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{1}, user_ids)
	ok, err := DeleteFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	user_ids, err = GetFollowUserIdList(ctx, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{3, 4}, user_ids)
	user_ids, err = GetFollowUserIdList(ctx, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, 0, len(user_ids))
}

func TestGetFollowerUserIdList(t *testing.T) {
	ctx := context.Background()
	Init()
	newfollow1 := Follow{
		UserId:     1,
		FollowerId: 2,
	}
	newfollow2 := Follow{
		UserId:     2,
		FollowerId: 1,
	}
	newfollow3 := Follow{
		UserId:     1,
		FollowerId: 3,
	}
	newfollow4 := Follow{
		UserId:     1,
		FollowerId: 4,
	}
	_, err := CreateFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow4)
	assert.Nil(t, err)
	user_ids, err := GetFollowerUserIdList(ctx, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{2, 3, 4}, user_ids)
	user_ids, err = GetFollowerUserIdList(ctx, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{1}, user_ids)
	ok, err := DeleteFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	user_ids, err = GetFollowerUserIdList(ctx, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{3, 4}, user_ids)
	user_ids, err = GetFollowerUserIdList(ctx, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, 0, len(user_ids))
}

func TestGetFriendUserIdList(t *testing.T) {
	ctx := context.Background()
	Init()
	newfollow1 := Follow{
		UserId:     1,
		FollowerId: 2,
	}
	newfollow2 := Follow{
		UserId:     2,
		FollowerId: 1,
	}
	newfollow3 := Follow{
		UserId:     1,
		FollowerId: 3,
	}
	newfollow4 := Follow{
		UserId:     1,
		FollowerId: 4,
	}
	newfollow5 := Follow{
		UserId:     4,
		FollowerId: 1,
	}
	_, err := CreateFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow3)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow4)
	assert.Nil(t, err)
	_, err = CreateFollow(ctx, &newfollow5)
	assert.Nil(t, err)
	user_ids, err := GetFriendUserIdList(ctx, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{2, 4}, user_ids)
	user_ids, err = GetFriendUserIdList(ctx, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{1}, user_ids)
	user_ids, err = GetFriendUserIdList(ctx, 4)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{1}, user_ids)
	ok, err := DeleteFollow(ctx, &newfollow1)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	ok, err = DeleteFollow(ctx, &newfollow2)
	assert.Nil(t, err)
	assert.DeepEqual(t, true, ok)
	user_ids, err = GetFriendUserIdList(ctx, 1)
	assert.Nil(t, err)
	assert.DeepEqual(t, []int64{4}, user_ids)
	user_ids, err = GetFriendUserIdList(ctx, 2)
	assert.Nil(t, err)
	assert.DeepEqual(t, 0, len(user_ids))
}
