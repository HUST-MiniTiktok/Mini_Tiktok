package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/client"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	message "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
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

func ToKitexFriendUserList(ctx context.Context, curr_user_token string, friend_user_ids []int64) ([]*relation.FriendUser, error) {
	// 协程补全
	kitex_friend_users := make([]*relation.FriendUser, 0, len(friend_user_ids))
	errChan := make(chan error)
	userChan := make(chan *relation.FriendUser)
	for _, user_id := range friend_user_ids {
		go func(user_id int64) {
			userResp, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: user_id})
			if err != nil {
				errChan <- err
				return
			}
			friendMsgResp, err := client.MessageRPC.GetFriendLatestMsg(ctx, &message.GetFriendLatestMsgRequest{Token: curr_user_token, FriendUserId: user_id})
			if err != nil {
				errChan <- err
			} else {
				friend_user := &relation.FriendUser{
					User:    userResp.User,
					Message: friendMsgResp.Message,
					MsgType: friendMsgResp.MsgType,
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
