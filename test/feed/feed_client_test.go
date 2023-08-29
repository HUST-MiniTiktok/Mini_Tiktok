package feed_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestFeedClient(t *testing.T) {
	client := rpc.NewFeedClient()
	if client == nil {
		t.Errorf("new feed client failed")
		return
	}
	t.Logf("new feed client success")
}
