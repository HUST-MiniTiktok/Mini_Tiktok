package conv

import (
	hertz_message "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/message"
	kitex_message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
)

func ToHertzMessage(message *kitex_message.Message) *hertz_message.Message {
	return &hertz_message.Message{
		ID:         message.Id,
		FromUserID: message.FromUserId,
		ToUserID:   message.ToUserId,
		Content:    message.Content,
		CreateTime: message.CreateTime,
	}
}

func ToHertzMessageList(message_list []*kitex_message.Message) []*hertz_message.Message {
	hertz_message_list := make([]*hertz_message.Message, len(message_list))
	for _, message := range message_list {
		hertz_message_list = append(hertz_message_list, ToHertzMessage(message))
	}
	return hertz_message_list
}

func ToKitexMessageChatRequest(req *hertz_message.MessageChatRequest) *kitex_message.MessageChatRequest {
	return &kitex_message.MessageChatRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		PreMsgTime: req.PreMsgTime,
	}
}

func ToHertzMessageChatResponse(resp *kitex_message.MessageChatResponse) *hertz_message.MessageChatResponse {
	return &hertz_message.MessageChatResponse{
		StatusCode:  resp.StatusCode,
		StatusMsg:   resp.StatusMsg,
		MessageList: ToHertzMessageList(resp.MessageList),
	}
}

func ToKitexMessageActionRequest(req *hertz_message.MessageActionRequest) *kitex_message.MessageActionRequest {
	return &kitex_message.MessageActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
	}
}

func ToHertzMessageActionResponse(resp *kitex_message.MessageActionResponse) *hertz_message.MessageActionResponse {
	return &hertz_message.MessageActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	}
}
