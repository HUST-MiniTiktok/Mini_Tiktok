package favorite_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestFavoriteClient(t *testing.T) {
	client := rpc.NewFavoriteClient()
	if client == nil {
		t.Errorf("new favorite client failed")
		return
	}
	t.Logf("new favorite client success")
}
