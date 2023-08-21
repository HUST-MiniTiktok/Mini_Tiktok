package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/client"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

func ToKitexUserList(ctx context.Context, user_ids []int64) ([]*common.User, error) {
	// 协程补全
	kitex_users := make([]*common.User, 0, len(user_ids))
	errChan := make(chan error)
	userChan := make(chan *common.User)
	for _, user_id := range user_ids {
		go func(user_id int64) {
			userResp, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: user_id})
			if err != nil {
				errChan <- err
			} else {
				userChan <- userResp.User
			}
		}(user_id)
	}
	// 等待协程结束
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

func ToKitexFriendUserList(ctx context.Context, curr_user_token string, friend_user_ids []int64) ([]*common.FriendUser, error) {
	// 协程补全
	kitex_friend_users := make([]*common.FriendUser, 0, len(friend_user_ids))
	errChan := make(chan error)
	userChan := make(chan *common.FriendUser)
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
	// 等待协程结束
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
