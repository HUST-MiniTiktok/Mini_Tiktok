package rpc

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

var (
	UserRPC *rpc.UserClient
)

func init() {
	UserRPC = rpc.NewUserClient()
}
