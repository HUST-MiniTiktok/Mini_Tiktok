package comment_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestCommentClient(t *testing.T) {
	client := rpc.NewCommentClient()
	if client == nil {
		t.Errorf("new comment client failed")
		return
	}
	t.Logf("new comment client success")
}
