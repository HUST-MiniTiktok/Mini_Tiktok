package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/client"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
)

// 补全视频信息
func ToKitexVideo(ctx context.Context, db_video *db.Video, user_id int64) (*common.Video, error) {
	// 复制视频信息
	kitex_video := &common.Video{
		Id:       db_video.ID,
		PlayUrl:  oss.ToRealURL(ctx, db_video.PlayURL),
		CoverUrl: oss.ToRealURL(ctx, db_video.CoverURL),
		Title:    db_video.Title,
	}

	// 协程补全
	errChan := make(chan error)
	authorChan := make(chan *user.UserResponse)
	favoriteInfoChan := make(chan *favorite.GetVideoFavoriteInfoResponse)
	commentCountChan := make(chan *comment.GetVideoCommentCountResponse)

	go func() {
		author, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: db_video.AuthorID})
		if err != nil {
			errChan <- err
		} else {
			authorChan <- author
		}
	}()

	go func() {
		comment_count, err := client.CommentRPC.GetVideoCommentCount(ctx, &comment.GetVideoCommentCountRequest{VideoId: db_video.ID})
		if err != nil {
			errChan <- err
		} else {
			commentCountChan <- comment_count
		}
	}()

	go func() {
		favorite_info, err := client.FavoriteRPC.GetVideoFavoriteInfo(ctx, &favorite.GetVideoFavoriteInfoRequest{UserId: user_id, VideoId: db_video.ID})
		if err != nil {
			errChan <- err
		} else {
			favoriteInfoChan <- favorite_info
		}
	}()

	// 等待协程结束
	for i := 0; i < 3; i++ {
		select {
		case err := <-errChan:
			return nil, err
		case author := <-authorChan:
			kitex_video.Author = author.User
		case comment_count := <-commentCountChan:
			kitex_video.CommentCount = comment_count.CommentCount
		case favorite_info := <-favoriteInfoChan:
			kitex_video.FavoriteCount = favorite_info.FavoriteCount
			kitex_video.IsFavorite = favorite_info.IsFavorite
		}
	}

	return kitex_video, nil
}
