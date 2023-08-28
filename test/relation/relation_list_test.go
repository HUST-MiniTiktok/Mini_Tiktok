package relation_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/pack"
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"

	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
)

func TestRelationFollowList(t *testing.T) {
	monkey.Patch(pack.ToKitexUserList, PatchedToKitexUserList)
	defer monkey.Unpatch(pack.ToKitexUserList)

	// user 0 follow list
	resp, err := RelationService.RelationFollowList(&relation.RelationFollowListRequest{
		UserId: DemoUserList[0].Id,
		Token:  DemoUserList[0].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follow list response: %v", resp)
	// user 1 follow list
	resp, err = RelationService.RelationFollowList(&relation.RelationFollowListRequest{
		UserId: DemoUserList[1].Id,
		Token:  DemoUserList[1].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follow list response: %v", resp)
	// user 2 follow list
	resp, err = RelationService.RelationFollowList(&relation.RelationFollowListRequest{
		UserId: DemoUserList[2].Id,
		Token:  DemoUserList[2].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follow list response: %v", resp)
}

func TestRelationFollowerList(t *testing.T) {
	monkey.Patch(pack.ToKitexUserList, PatchedToKitexUserList)
	defer monkey.Unpatch(pack.ToKitexUserList)

	// user 0 follower list
	resp, err := RelationService.RelationFollowerList(&relation.RelationFollowerListRequest{
		UserId: DemoUserList[0].Id,
		Token:  DemoUserList[0].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follower list response: %v", resp)
	// user 1 follower list
	resp, err = RelationService.RelationFollowerList(&relation.RelationFollowerListRequest{
		UserId: DemoUserList[1].Id,
		Token:  DemoUserList[1].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follower list response: %v", resp)
	// user 2 follower list
	resp, err = RelationService.RelationFollowerList(&relation.RelationFollowerListRequest{
		UserId: DemoUserList[2].Id,
		Token:  DemoUserList[2].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follower list response: %v", resp)
}

func TestRelationFriendList(t *testing.T) {
	monkey.Patch(pack.ToKitexFriendUserList, PatchedToKitexFriendUserList)
	defer monkey.Unpatch(pack.ToKitexFriendUserList)

	// user 0 friend list
	resp, err := RelationService.RelationFriendList(&relation.RelationFriendListRequest{
		UserId: DemoUserList[0].Id,
		Token:  DemoUserList[0].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation friend list response: %v", resp)
	// user 1 friend list
	resp, err = RelationService.RelationFriendList(&relation.RelationFriendListRequest{
		UserId: DemoUserList[1].Id,
		Token:  DemoUserList[1].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation friend list response: %v", resp)
	// user 2 friend list
	resp, err = RelationService.RelationFriendList(&relation.RelationFriendListRequest{
		UserId: DemoUserList[2].Id,
		Token:  DemoUserList[2].Token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation friend list response: %v", resp)
}

func PatchedToKitexUserList(ctx context.Context, curr_user_token string, user_ids []int64) ([]*common.User, error) {
	kitex_users := make([]*common.User, 0, len(user_ids))
	for _, user_id := range user_ids {
		kitex_users = append(kitex_users, &common.User{
			Id: user_id,
		})
	}
	return kitex_users, nil
}

func PatchedToKitexFriendUserList(ctx context.Context, curr_user_token string, friend_user_ids []int64) ([]*common.FriendUser, error) {
	kitex_users := make([]*common.FriendUser, 0, len(friend_user_ids))
	for _, friend_user_id := range friend_user_ids {
		kitex_users = append(kitex_users, &common.FriendUser{
			Id: friend_user_id,
		})
	}
	return kitex_users, nil
}
