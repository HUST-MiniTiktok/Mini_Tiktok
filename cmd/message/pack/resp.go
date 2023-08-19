package pack

import (
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewMessageChatResponse(err error) *message.MessageChatResponse {
	return conv.ToKitexBaseResponse(err, &message.MessageChatResponse{}).(*message.MessageChatResponse)
}

func NewMessageActionResponse(err error) *message.MessageActionResponse {
	return conv.ToKitexBaseResponse(err, &message.MessageActionResponse{}).(*message.MessageActionResponse)
}

func NewGetFriendLatestMsgResponse(err error) *message.GetFriendLatestMsgResponse {
	return conv.ToKitexBaseResponse(err, &message.GetFriendLatestMsgResponse{}).(*message.GetFriendLatestMsgResponse)
}
