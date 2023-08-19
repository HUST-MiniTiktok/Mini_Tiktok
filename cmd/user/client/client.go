package client

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

var (
	PublishRPC  *rpc.PublishClient
	RelationRPC *rpc.RelationClient
	FavoriteRPC *rpc.FavoriteClient
)

func init() {
	PublishRPC = rpc.NewPublishClient()
	RelationRPC = rpc.NewRelationClient()
	FavoriteRPC = rpc.NewFavoriteClient()
}
