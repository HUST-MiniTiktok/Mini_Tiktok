package client

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

var (
	PublishRPC *rpc.PublishClient
)

func init() {
	PublishRPC = rpc.NewPublishClient()
}