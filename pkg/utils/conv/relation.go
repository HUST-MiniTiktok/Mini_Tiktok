package conv

import (
	hertz_relation "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/relation"
	kitex_relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
)

func ToHertzFriendUser(user *kitex_relation.FriendUser) *hertz_relation.FriendUser {
	return &hertz_relation.FriendUser{
		User:    ToHertzUser(user.User),
		Message: user.Message,
		MsgType: user.MsgType,
	}
}

func ToHertzFriendUserList(user_list []*kitex_relation.FriendUser) []*hertz_relation.FriendUser {
	hertz_user_list := make([]*hertz_relation.FriendUser, 0, len(user_list))
	for _, user := range user_list {
		hertz_user_list = append(hertz_user_list, ToHertzFriendUser(user))
	}
	return hertz_user_list
}

func ToKitexRelationActionRequest(req *hertz_relation.RelationActionRequest) *kitex_relation.RelationActionRequest {
	return &kitex_relation.RelationActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	}
}

func ToHertzRelationActionResponse(resp *kitex_relation.RelationActionResponse) *hertz_relation.RelationActionResponse {
	return &hertz_relation.RelationActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	}
}

func ToKitexRelationFollowListRequest(req *hertz_relation.RelationFollowListRequest) *kitex_relation.RelationFollowListRequest {
	return &kitex_relation.RelationFollowListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
}

func ToHertzRelationFollowListResponse(resp *kitex_relation.RelationFollowListResponse) *hertz_relation.RelationFollowListResponse {
	return &hertz_relation.RelationFollowListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserList:   ToHertzUserList(resp.UserList),
	}
}

func ToKitexRelationFollowerListRequest(req *hertz_relation.RelationFollowerListRequest) *kitex_relation.RelationFollowerListRequest {
	return &kitex_relation.RelationFollowerListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
}

func ToHertzRelationFollowerListResponse(resp *kitex_relation.RelationFollowerListResponse) *hertz_relation.RelationFollowerListResponse {
	return &hertz_relation.RelationFollowerListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserList:   ToHertzUserList(resp.UserList),
	}
}

func ToKitexRelationFriendListRequest(req *hertz_relation.RelationFriendListRequest) *kitex_relation.RelationFriendListRequest {
	return &kitex_relation.RelationFriendListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
}

func ToHertzRelationFriendListResponse(resp *kitex_relation.RelationFriendListResponse) *hertz_relation.RelationFriendListResponse {
	return &hertz_relation.RelationFriendListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserList:   ToHertzFriendUserList(resp.UserList),
	}
}
