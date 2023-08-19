package pack

import (
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewRelationActionResponse(err error) *relation.RelationActionResponse {
	return conv.ToKitexBaseResponse(err, &relation.RelationActionResponse{}).(*relation.RelationActionResponse)
}

func NewRelationFollowListResponse(err error) *relation.RelationFollowListResponse {
	return conv.ToKitexBaseResponse(err, &relation.RelationFollowListResponse{}).(*relation.RelationFollowListResponse)
}

func NewRelationFollowerListResponse(err error) *relation.RelationFollowerListResponse {
	return conv.ToKitexBaseResponse(err, &relation.RelationFollowerListResponse{}).(*relation.RelationFollowerListResponse)
}

func NewRelationFriendListResponse(err error) *relation.RelationFriendListResponse {
	return conv.ToKitexBaseResponse(err, &relation.RelationFriendListResponse{}).(*relation.RelationFriendListResponse)
}

func NewGetFollowInfoResponse(err error) *relation.GetFollowInfoResponse {
	return conv.ToKitexBaseResponse(err, &relation.GetFollowInfoResponse{}).(*relation.GetFollowInfoResponse)
}
