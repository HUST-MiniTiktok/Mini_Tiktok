package rpc

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/rpc"
)

var (
	PublishRPC *rpc.PublishClient
)

func init() {
	PublishRPC = rpc.NewPublishClient()
}