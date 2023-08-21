package service

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/client"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/pack"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils"
)

var (
	Jwt *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT()
}

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{ctx: ctx}
}

func (s *MessageService) MessageChat(request *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := user_claims.ID
	if err != nil {
		return pack.NewMessageChatResponse(errno.AuthorizationFailedErr), err
	}

	var db_messages []*db.Message
	db_messages, err = db.GetMessageByUserIdPair(s.ctx, curr_user_id, request.ToUserId, utils.MillTimeStampToTime(request.PreMsgTime))
	if err != nil {
		return pack.NewMessageChatResponse(err), err
	}

	kitex_messages := make([]*message.Message, 0, len(db_messages))
	for _, db_message := range db_messages {
		create_time := utils.TimeToMillTimeStamp(db_message.CreatedAt)
		kitex_messages = append(kitex_messages, &message.Message{
			Id:         db_message.ID,
			FromUserId: db_message.FromUserId,
			ToUserId:   db_message.ToUserId,
			Content:    db_message.Content,
			CreateTime: &create_time,
		})
	}
	resp = pack.NewMessageChatResponse(errno.Success)
	resp.MessageList = kitex_messages
	return resp, nil
}

func (s *MessageService) MessageAction(request *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	from_user_id := user_claims.ID
	if err != nil {
		return pack.NewMessageActionResponse(errno.AuthorizationFailedErr), err
	}

	if request.ActionType != 1 {
		return pack.NewMessageActionResponse(errno.ParamErr), errno.ParamErr
	}

	to_user_ck, err := client.UserRPC.CheckUserIsExist(s.ctx, &user.CheckUserIsExistRequest{UserId: request.ToUserId})
	if err != nil {
		return pack.NewMessageActionResponse(err), err
	}
	if !to_user_ck.IsExist {
		return pack.NewMessageActionResponse(errno.UserIsNotExistErr), errno.UserIsNotExistErr
	}

	db_message := db.Message{FromUserId: from_user_id, ToUserId: request.ToUserId, Content: request.Content}
	_, err = db.CreateMessage(s.ctx, &db_message)
	if err != nil {
		return pack.NewMessageActionResponse(err), err
	}

	return pack.NewMessageActionResponse(errno.Success), nil
}

func (s *MessageService) GetFriendLatestMsg(request *message.GetFriendLatestMsgRequest) (resp *message.GetFriendLatestMsgResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := user_claims.ID
	if err != nil {
		return pack.NewGetFriendLatestMsgResponse(errno.AuthorizationFailedErr), err
	}

	var db_message *db.Message
	db_message, err = db.GetLastestMsgByUserIdPair(s.ctx, curr_user_id, request.FriendUserId)
	if err != nil {
		return pack.NewGetFriendLatestMsgResponse(err), err
	}

	var msgType int64
	if db_message == nil { // db_message == nil means no message between friend users
		msgType = 2 // 2 means no message
		resp = pack.NewGetFriendLatestMsgResponse(errno.Success)
		msg := &db.Message{}
		resp.MsgType = &msgType
		resp.Message = &msg.Content
	} else if db_message.ToUserId == curr_user_id {
		msgType = 0 // 0 means friend user send message to curr user
		resp = pack.NewGetFriendLatestMsgResponse(errno.Success)
		resp.MsgType = &msgType
		resp.Message = &db_message.Content
	} else { // db_message.FromUserId == curr_user_id
		msgType = 1 // 1 means curr user send message to friend user
		resp = pack.NewGetFriendLatestMsgResponse(errno.Success)
		resp.MsgType = &msgType
		resp.Message = &db_message.Content
	}

	return resp, nil
}
