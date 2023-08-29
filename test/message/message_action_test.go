package message_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/pack"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
)

func TestMessageAction(t *testing.T) {
	monkey.Patch(pack.IsExistUser, func(ctx context.Context, user_id int64) (bool, error) {
		return true, nil
	})
	defer monkey.Unpatch(pack.IsExistUser)

	resp, err := MessageService.MessageAction(&message.MessageActionRequest{
		Token:      DemoUserList[0].Token,
		ToUserId:   DemoUserList[1].Id,
		ActionType: 1,
		Content:    "hello from user 0 to user 1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
	resp, err = MessageService.MessageAction(&message.MessageActionRequest{
		Token:      DemoUserList[1].Token,
		ToUserId:   DemoUserList[0].Id,
		ActionType: 1,
		Content:    "reply from user 1 to user 0",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
	resp, err = MessageService.MessageAction(&message.MessageActionRequest{
		Token:      DemoUserList[1].Token,
		ToUserId:   DemoUserList[0].Id,
		ActionType: 1,
		Content:    "reply again from user 1 to user 0",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
}
