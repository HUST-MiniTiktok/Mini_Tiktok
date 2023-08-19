package main

import (
	"context"

	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/service"
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, request *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	relation_service := service.NewRelationService(ctx)
	resp, err = relation_service.RelationAction(request)
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, request *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	relation_service := service.NewRelationService(ctx)
	resp, err = relation_service.RelationFollowList(request)
	return
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, request *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	relation_service := service.NewRelationService(ctx)
	resp, err = relation_service.RelationFollowerList(request)
	return
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, request *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	relation_service := service.NewRelationService(ctx)
	resp, err = relation_service.RelationFriendList(request)
	return
}

// GetFollowInfo implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowInfo(ctx context.Context, request *relation.GetFollowInfoRequest) (resp *relation.GetFollowInfoResponse, err error) {
	relation_service := service.NewRelationService(ctx)
	resp, err = relation_service.GetFollowInfo(request)
	return
}
