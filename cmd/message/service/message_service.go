package service

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/rpc"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/dal/db"
	_ "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/rpc"
	_ "github.com/HUST-MiniTiktok/mini_tiktok/conf"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
	"github.com/HUST-MiniTiktok/mini_tiktok/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
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
	klog.Infof("message_chat request: %v", *request)
	user_claims, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := user_claims.ID

	if err != nil {
		err_msg := err.Error()
		return &message.MessageChatResponse{StatusCode: int32(codes.PermissionDenied), StatusMsg: &err_msg}, err
	}

	var db_messages []*db.Message
	db_messages, err = db.GetMessageByUserIdPair(s.ctx, curr_user_id, request.ToUserId, util.MillTimeStampToTime(request.PreMsgTime))
	if err != nil {
		err_msg := err.Error()
		return &message.MessageChatResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
	}

	kitex_messages := make([]*message.Message, len(db_messages))
	for _, db_message := range db_messages {
		create_time := util.TimeToMillTimeStamp(db_message.CreatedAt)
		kitex_messages = append(kitex_messages, &message.Message{
			Id:         db_message.ID,
			FromUserId: db_message.FromUserId,
			ToUserId:   db_message.ToUserId,
			Content:    db_message.Content,
			CreateTime: &create_time,
		})
	}
	return &message.MessageChatResponse{StatusCode: int32(codes.OK), StatusMsg: nil, MessageList: kitex_messages}, nil
}

func (s *MessageService) MessageAction(request *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	klog.Infof("message_chat request: %v", *request)
	user_claims, err := Jwt.ExtractClaims(request.Token)
	from_user_id := user_claims.ID

	if err != nil {
		err_msg := err.Error()
		return &message.MessageActionResponse{StatusCode: int32(codes.PermissionDenied), StatusMsg: &err_msg}, err
	}

	if request.ActionType != 1 {
		err_msg := "invalid action_type"
		return &message.MessageActionResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &err_msg}, err
	}

	to_user_ck, err := rpc.UserRPC.CheckUserIsExist(s.ctx, &user.CheckUserIsExistRequest{UserId: request.ToUserId})
	if err != nil {
		err_msg := err.Error()
		return &message.MessageActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
	}
	if !to_user_ck.IsExist {
		err_msg := "to_user is not exist"
		return &message.MessageActionResponse{StatusCode: int32(codes.PermissionDenied), StatusMsg: &err_msg}, err
	}

	db_message := db.Message{FromUserId: from_user_id, ToUserId: request.ToUserId, Content: request.Content}
	_, err = db.CreateMessage(s.ctx, &db_message)
	if err != nil {
		err_msg := err.Error()
		return &message.MessageActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
	}

	return &message.MessageActionResponse{StatusCode: int32(codes.OK), StatusMsg: nil}, nil
}
