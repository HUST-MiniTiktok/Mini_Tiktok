package message_test

import (
	"testing"

	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
)

func TestGetFriendLatestMsg(t *testing.T) {
	resp, err := MessageService.GetFriendLatestMsg(&message.GetFriendLatestMsgRequest{
		Token:        DemoUserList[0].Token,
		FriendUserId: DemoUserList[1].Id,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
	if resp.GetMessage() != "reply again from user 1 to user 0" {
		t.Fatal("message not correct")
	}
}
