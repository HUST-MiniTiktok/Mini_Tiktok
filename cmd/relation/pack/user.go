package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/client"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

// ToKitexUserList: get []*common.User from []user_ids
func ToKitexUserList(ctx context.Context, curr_user_token string, user_ids []int64) ([]*common.User, error) {
	kitex_users := make([]*common.User, 0, len(user_ids))
	errChan := make(chan error)
	userChan := make(chan *common.User)
	// create goroutines to get user info
	for _, user_id := range user_ids {
		go func(user_id int64) {
			userResp, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: user_id, Token: curr_user_token})
			if err != nil {
				errChan <- err
			} else {
				userChan <- userResp.User
			}
		}(user_id)
	}
	// wait for goroutines to finish
	for i := 0; i < len(user_ids); i++ {
		select {
		case err := <-errChan:
			return nil, err
		case user := <-userChan:
			kitex_users = append(kitex_users, user)
		}
	}
	return kitex_users, nil
}

// ToKitexFriendUserList: get []*common.FriendUser from []friend_user_ids
func ToKitexFriendUserList(ctx context.Context, curr_user_token string, friend_user_ids []int64) ([]*common.FriendUser, error) {
	kitex_friend_users := make([]*common.FriendUser, 0, len(friend_user_ids))
	errChan := make(chan error)
	userChan := make(chan *common.FriendUser)
	// create goroutines to get user info
	for _, user_id := range friend_user_ids {
		go func(user_id int64) {
			userResp, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: user_id, Token: curr_user_token})
			if err != nil {
				errChan <- err
				return
			}
			friendMsgResp, err := client.MessageRPC.GetFriendLatestMsg(ctx, &message.GetFriendLatestMsgRequest{FriendUserId: user_id, Token: curr_user_token})
			if err != nil {
				errChan <- err
			} else {
				friend_user := &common.FriendUser{
					Id:              userResp.User.Id,
					Name:            userResp.User.Name,
					FollowCount:     userResp.User.FollowCount,
					FollowerCount:   userResp.User.FollowerCount,
					IsFollow:        userResp.User.IsFollow,
					Avatar:          userResp.User.Avatar,
					BackgroundImage: userResp.User.BackgroundImage,
					Signature:       userResp.User.Signature,
					TotalFavorited:  userResp.User.TotalFavorited,
					WorkCount:       userResp.User.WorkCount,
					FavoriteCount:   userResp.User.FavoriteCount,
					Message:         friendMsgResp.Message,
					MsgType:         friendMsgResp.MsgType,
				}
				userChan <- friend_user
			}
		}(user_id)
	}
	// wait for goroutines to finish
	for i := 0; i < len(friend_user_ids); i++ {
		select {
		case err := <-errChan:
			return nil, err
		case user := <-userChan:
			kitex_friend_users = append(kitex_friend_users, user)
		}
	}
	return kitex_friend_users, nil
}
