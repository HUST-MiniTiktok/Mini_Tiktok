package message_test

import (
	"testing"
	"time"

	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	utils "github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils"
)

func TestMessageChat(t *testing.T) {
	resp, err := MessageService.MessageChat(&message.MessageChatRequest{
		Token:      DemoUserList[0].Token,
		ToUserId:   DemoUserList[1].Id,
		PreMsgTime: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
	if resp.MessageList == nil || len(resp.MessageList) != 3 {
		t.Fatal("message list not correct")
	}
	resp, err = MessageService.MessageChat(&message.MessageChatRequest{
		Token:      DemoUserList[1].Token,
		ToUserId:   DemoUserList[0].Id,
		PreMsgTime: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
	if resp.MessageList == nil || len(resp.MessageList) != 3 {
		t.Fatal("message list not correct")
	}
	resp, err = MessageService.MessageChat(&message.MessageChatRequest{
		Token:      DemoUserList[2].Token,
		ToUserId:   DemoUserList[0].Id,
		PreMsgTime: utils.TimeToMillTimeStamp(time.Now()),
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("message response: %v", resp)
	if resp.MessageList != nil && len(resp.MessageList) != 0 {
		t.Fatal("message list not correct")
	}
}
