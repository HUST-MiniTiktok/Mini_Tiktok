package rpc

import (
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/mw/rpc"
)

var (
	UserRPC *rpc.UserClient
)

func init() {
	UserRPC = rpc.NewUserClient()
}
