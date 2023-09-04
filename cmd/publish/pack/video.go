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

// ToKitexVideo: convert db.Video to common.Video
func ToKitexVideo(ctx context.Context, curr_user_id int64, curr_user_token string, db_video *db.Video) (*common.Video, error) {
	kitex_video := &common.Video{
		Id:       db_video.ID,
		PlayUrl:  oss.ToRealURL(ctx, db_video.PlayURL),
		CoverUrl: oss.ToRealURL(ctx, db_video.CoverURL),
		Title:    db_video.Title,
	}

	errChan := make(chan error)
	authorChan := make(chan *user.UserResponse)
	favoriteInfoChan := make(chan *favorite.GetVideoFavoriteInfoResponse)
	commentCountChan := make(chan *comment.GetVideoCommentCountResponse)
	// create goroutines to get author, comment count and favorite info
	go func() {
		author, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: db_video.AuthorID, Token: curr_user_token})
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
		favorite_info, err := client.FavoriteRPC.GetVideoFavoriteInfo(ctx, &favorite.GetVideoFavoriteInfoRequest{UserId: curr_user_id, VideoId: db_video.ID})
		if err != nil {
			errChan <- err
		} else {
			favoriteInfoChan <- favorite_info
		}
	}()

	// wait for goroutines to finish
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

// ToKitexVideoList converts []*db.Video to []*common.Video
func ToKitexVideoList(ctx context.Context, curr_user_id int64, curr_user_token string, db_videos []*db.Video) ([]*common.Video, error) {
	kitex_videos := make([]*common.Video, 0, len(db_videos))
	err_chan := make(chan error)
	video_chan := make(chan *common.Video)
	for _, db_video := range db_videos {
		go func(db_video *db.Video) {
			kitex_video, err := ToKitexVideo(ctx, curr_user_id, curr_user_token, db_video)
			if err != nil {
				err_chan <- err
			} else {
				video_chan <- kitex_video
			}
		}(db_video)
	}
	for i := 0; i < len(db_videos); i++ {
		select {
		case err := <-err_chan:
			return nil, err
		case kitex_video := <-video_chan:
			kitex_videos = append(kitex_videos, kitex_video)
		}
	}
	return kitex_videos, nil
}
