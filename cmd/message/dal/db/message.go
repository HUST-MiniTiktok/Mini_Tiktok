package db

import (
	"context"
	"time"
)

const MessageTableName = "message"

type Message struct {
	ID         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id"`
	FromUserId int64     `json:"from_user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Message) TableName() string {
	return MessageTableName
}

func CreateMessage(ctx context.Context, message *Message) (id int64, err error) {
	err = DB.WithContext(ctx).Create(message).Error
	if err != nil {
		return -1, err
	}
	id = message.ID
	return
}

func GetMessageById(ctx context.Context, id int64) (message *Message, err error) {
	err = DB.WithContext(ctx).First(&message, id).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetMessageByUserIdPair(ctx context.Context, toUserId, fromUserId int64, preMsgTime time.Time) (messages []*Message, err error) {
	//  preMsgTime 上次最新消息的时间
	err = DB.WithContext(ctx).Where("to_user_id = ? AND from_user_id = ? AND created_at > ?", toUserId, fromUserId, preMsgTime).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return
}
