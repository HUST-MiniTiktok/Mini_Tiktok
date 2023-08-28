package relation_test

import (
	"testing"

	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
)

func TestGetFollowInfoByUser(t *testing.T) {
	resp, err := RelationService.GetFollowInfo(&relation.GetFollowInfoRequest{
		Token:    DemoUserList[0].Token,
		ToUserId: DemoUserList[1].Id,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("get follow info by user response: %v", resp)
}
