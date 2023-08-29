package comment_test

import (
	"testing"

	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
)

func TestGetVideoCommentCount(t *testing.T) {
	resp, err := CommentService.GetVideoCommentCount(&comment.GetVideoCommentCountRequest{
		VideoId: DemoVideo.Id,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("Delete Comment response: %v", resp)
}
