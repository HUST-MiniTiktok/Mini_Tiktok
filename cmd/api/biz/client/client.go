package client

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

var (
	UserRPC *rpc.UserClient
	PublishRPC *rpc.PublishClient
	FeedRPC *rpc.FeedClient
	FavoriteRPC *rpc.FavoriteClient
	CommentRPC *rpc.CommentClient
	RelationRPC *rpc.RelationClient
	MessageRPC *rpc.MessageClient
)

func init() {
	UserRPC = rpc.NewUserClient()
	PublishRPC = rpc.NewPublishClient()
	FeedRPC = rpc.NewFeedClient()
	FavoriteRPC = rpc.NewFavoriteClient()
	CommentRPC = rpc.NewCommentClient()
	RelationRPC = rpc.NewRelationClient()
	MessageRPC = rpc.NewMessageClient()
}

