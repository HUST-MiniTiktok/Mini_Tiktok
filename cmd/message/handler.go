package main

import (
	"context"

	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/service"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, request *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	message_service := service.NewMessageService(ctx)
	resp, err = message_service.MessageChat(request)
	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, request *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	message_service := service.NewMessageService(ctx)
	resp, err = message_service.MessageAction(request)
	return
}

// GetFriendLatestMsg implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetFriendLatestMsg(ctx context.Context, request *message.GetFriendLatestMsgRequest) (resp *message.GetFriendLatestMsgResponse, err error) {
	message_service := service.NewMessageService(ctx)
	resp, err = message_service.GetFriendLatestMsg(request)
	return
}
