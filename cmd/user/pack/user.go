package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/client"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal/db"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
)

// ToKitexUser: convert db.User to common.User
func ToKitexUser(ctx context.Context, curr_user_token string, db_user *db.User) (kitex_user *common.User, err error) {
	// 设置默认信息
	default_avatar := oss.ToRealURL(ctx, "image/Avatar.png")
	default_background_image := oss.ToRealURL(ctx, "image/Background.png")
	default_signature := "This is my signature."
	kitex_user = &common.User{
		Id:              db_user.ID,
		Name:            db_user.UserName,
		Avatar:          &default_avatar,
		BackgroundImage: &default_background_image,
		Signature:       &default_signature,
	}

	errChan := make(chan error)
	publishInfoChan := make(chan *publish.GetPublishInfoByUserIdResponse)
	followInfoChan := make(chan *relation.GetFollowInfoResponse)
	favoriteInfoChan := make(chan *favorite.GetUserFavoriteInfoResponse)
	// create goroutines to get publish info, follow info and favorite info
	go func() {
		publish_info, err := client.PublishRPC.GetPublishInfoByUserId(ctx, &publish.GetPublishInfoByUserIdRequest{UserId: db_user.ID})
		if err != nil {
			errChan <- err
		} else {
			publishInfoChan <- publish_info
		}
	}()

	go func() {
		follow_info, err := client.RelationRPC.GetFollowInfo(ctx, &relation.GetFollowInfoRequest{Token: curr_user_token, ToUserId: db_user.ID})
		if err != nil {
			errChan <- err
		} else {
			followInfoChan <- follow_info
		}
	}()

	go func() {
		favorite_info, err := client.FavoriteRPC.GetUserFavoriteInfo(ctx, &favorite.GetUserFavoriteInfoRequest{UserId: db_user.ID})
		if err != nil {
			errChan <- err
		} else {
			favoriteInfoChan <- favorite_info
		}
	}()

	// wait for goroutines to finish
	for i := 0; i < 3; i++ {
		select {
		case err = <-errChan:
			return
		case publish_info := <-publishInfoChan:
			kitex_user.WorkCount = &publish_info.WorkCount
		case follow_info := <-followInfoChan:
			kitex_user.FollowCount = &follow_info.FollowCount
			kitex_user.FollowerCount = &follow_info.FollowerCount
			kitex_user.IsFollow = follow_info.IsFollow
		case favorite_info := <-favoriteInfoChan:
			kitex_user.TotalFavorited = &favorite_info.TotalFavorited
			kitex_user.FavoriteCount = &favorite_info.FavoriteCount
		}
	}

	return kitex_user, nil
}
