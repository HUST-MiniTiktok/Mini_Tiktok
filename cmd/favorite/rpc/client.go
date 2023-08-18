package rpc

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/rpc"
)

var (
	PublishRPC *rpc.PublishClient
)

func init() {
	PublishRPC = rpc.NewPublishClient()
}