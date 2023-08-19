package client

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

var (
	UserRPC    *rpc.UserClient
	MessageRPC *rpc.MessageClient
)

func init() {
	UserRPC = rpc.NewUserClient()
	MessageRPC = rpc.NewMessageClient()
}
