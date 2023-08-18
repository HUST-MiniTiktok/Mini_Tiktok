package rpc

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/rpc"
)

var (
	UserRPC *rpc.UserClient
	FavoriteRPC *rpc.FavoriteClient
)

func init() {
	UserRPC = rpc.NewUserClient()
	FavoriteRPC = rpc.NewFavoriteClient()
}