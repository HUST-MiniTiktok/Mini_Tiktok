package relation_test

import (
	"testing"

	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
)

func TestRelationFollowAction(t *testing.T) {
	// user 0 follow user 1
	resp, err := RelationService.RelationAction(&relation.RelationActionRequest{
		Token:      DemoUserList[0].Token,
		ToUserId:   DemoUserList[1].Id,
		ActionType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follow action response: %v", resp)
	// user 1 follow user 0
	resp, err = RelationService.RelationAction(&relation.RelationActionRequest{
		Token:      DemoUserList[1].Token,
		ToUserId:   DemoUserList[0].Id,
		ActionType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follow action response: %v", resp)
	// user 2 follow user 0
	resp, err = RelationService.RelationAction(&relation.RelationActionRequest{
		Token:      DemoUserList[2].Token,
		ToUserId:   DemoUserList[0].Id,
		ActionType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation follow action response: %v", resp)
}

func TestRelationUnFollowAction(t *testing.T) {
	// user 0 unfollow user 1
	resp, err := RelationService.RelationAction(&relation.RelationActionRequest{
		Token:      DemoUserList[0].Token,
		ToUserId:   DemoUserList[1].Id,
		ActionType: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation unfollow action response: %v", resp)
	// user 1 unfollow user 0
	resp, err = RelationService.RelationAction(&relation.RelationActionRequest{
		Token:      DemoUserList[1].Token,
		ToUserId:   DemoUserList[0].Id,
		ActionType: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation unfollow action response: %v", resp)
	// user 2 unfollow user 0
	resp, err = RelationService.RelationAction(&relation.RelationActionRequest{
		Token:      DemoUserList[2].Token,
		ToUserId:   DemoUserList[0].Id,
		ActionType: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("relation unfollow action response: %v", resp)
}
