package rpc

import (
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

var (
	UserRPC     *rpc.UserClient
	FavoriteRPC *rpc.FavoriteClient
	CommentRPC  *rpc.CommentClient
)

func init() {
	UserRPC = rpc.NewUserClient()
	FavoriteRPC = rpc.NewFavoriteClient()
	CommentRPC = rpc.NewCommentClient()
}
