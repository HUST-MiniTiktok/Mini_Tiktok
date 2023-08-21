package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/pack"
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

var (
	Jwt *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT()
}

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{ctx: ctx}
}

func (s *RelationService) RelationAction(request *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := claim.ID
	if err != nil || claim.ID == 0 {
		return pack.NewRelationActionResponse(errno.AuthorizationFailedErr), err
	}

	if request.ActionType == 1 {
		// Action 1: follow
		db_follow := db.Follow{
			UserId:     request.ToUserId,
			FollowerId: curr_user_id,
		}
		_, err := db.CreateFollow(s.ctx, &db_follow)
		if err != nil {
			return pack.NewRelationActionResponse(err), err
		}
	} else if request.ActionType == 2 {
		// Action 2: unfollow
		db_follow := db.Follow{
			UserId:     request.ToUserId,
			FollowerId: curr_user_id,
		}
		_, err := db.DeleteFollow(s.ctx, &db_follow)
		if err != nil {
			return pack.NewRelationActionResponse(err), err
		}
	} else {
		return pack.NewRelationActionResponse(errno.ParamErr), errno.ParamErr
	}
	return pack.NewRelationActionResponse(errno.Success), nil
}

func (s *RelationService) RelationFollowList(request *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := claim.ID
	if err != nil || claim.ID == 0 {
		return pack.NewRelationFollowListResponse(errno.AuthorizationFailedErr), err
	}

	user_ids, err := db.GetFollowUserIdList(s.ctx, curr_user_id)
	if err != nil {
		return pack.NewRelationFollowListResponse(err), err
	}

	kitex_users, err := pack.ToKitexUserList(s.ctx, user_ids)
	if err != nil {
		return pack.NewRelationFollowListResponse(err), err
	}

	resp = pack.NewRelationFollowListResponse(errno.Success)
	resp.UserList = kitex_users
	return resp, nil
}

func (s *RelationService) RelationFollowerList(request *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := claim.ID
	if err != nil || claim.ID == 0 {
		return pack.NewRelationFollowerListResponse(errno.AuthorizationFailedErr), err
	}

	user_ids, err := db.GetFollowerUserIdList(s.ctx, curr_user_id)
	if err != nil {
		return pack.NewRelationFollowerListResponse(err), err
	}

	kitex_users, err := pack.ToKitexUserList(s.ctx, user_ids)
	if err != nil {
		return pack.NewRelationFollowerListResponse(err), err
	}

	resp = pack.NewRelationFollowerListResponse(errno.Success)
	resp.UserList = kitex_users
	return resp, nil
}

func (s *RelationService) RelationFriendList(request *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	curr_user_id := claim.ID
	if err != nil || claim.ID == 0 {
		return pack.NewRelationFriendListResponse(errno.AuthorizationFailedErr), err
	}

	friend_user_ids, err := db.GetFriendUserIdList(s.ctx, curr_user_id)
	if err != nil {
		return pack.NewRelationFriendListResponse(err), err
	}

	kitex_friend_users, err := pack.ToKitexFriendUserList(s.ctx, request.Token, friend_user_ids)
	if err != nil {
		return pack.NewRelationFriendListResponse(err), err
	}

	resp = pack.NewRelationFriendListResponse(errno.Success)
	resp.UserList = kitex_friend_users
	return resp, nil
}

func (s *RelationService) GetFollowInfo(request *relation.GetFollowInfoRequest) (resp *relation.GetFollowInfoResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = claim.ID
	}

	follow_count, err := db.GetFollowUserCount(s.ctx, request.ToUserId)
	if err != nil {
		return pack.NewGetFollowInfoResponse(err), err
	}

	follower_count, err := db.GetFollowerUserCount(s.ctx, request.ToUserId)
	if err != nil {
		return pack.NewGetFollowInfoResponse(err), err
	}

	is_follow, err := db.CheckFollow(s.ctx, request.ToUserId, curr_user_id)
	if err != nil {
		return pack.NewGetFollowInfoResponse(err), err
	}

	resp = pack.NewGetFollowInfoResponse(errno.Success)
	resp.FollowCount = follow_count
	resp.FollowerCount = follower_count
	resp.IsFollow = is_follow

	return resp, nil
}
