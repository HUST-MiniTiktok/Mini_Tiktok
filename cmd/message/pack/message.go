package pack

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/dal/db"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils"
)

func ToKitexMessage(ctx context.Context, db_message *db.Message) *message.Message {
	create_time := utils.TimeToMillTimeStamp(db_message.CreatedAt)
	return &message.Message{
		Id:         db_message.ID,
		FromUserId: db_message.FromUserId,
		ToUserId:   db_message.ToUserId,
		Content:    db_message.Content,
		CreateTime: &create_time,
	}
}

func ToKitexMessageList(ctx context.Context, db_messages []*db.Message) []*message.Message {
	kitex_messages := make([]*message.Message, 0, len(db_messages))
	for _, db_message := range db_messages {
		kitex_message := ToKitexMessage(ctx, db_message)
		kitex_messages = append(kitex_messages, kitex_message)
	}
	return kitex_messages
}
