package feed_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/pack"
	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal/db"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
)

func TestFeed(t *testing.T) {
	monkey.Patch(pack.ToKitexVideo, func(ctx context.Context, curr_user_id int64, curr_user_token string, db_video *db.Video) (*common.Video, error) {
		return &common.Video{
			Id:       db_video.ID,
			PlayUrl:  oss.ToRealURL(ctx, db_video.PlayURL),
			CoverUrl: oss.ToRealURL(ctx, db_video.CoverURL),
			Title:    db_video.Title,
		}, nil
	})
	defer monkey.Unpatch(pack.ToKitexVideo)

	resp, err := FeedService.GetFeed(&feed.FeedRequest{Token: &DemoUser.Token})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("feed response: %v", resp)
}
