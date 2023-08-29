package relation_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestRelationClient(t *testing.T) {
	client := rpc.NewRelationClient()
	if client == nil {
		t.Errorf("new relation client failed")
		return
	}
	t.Logf("new relation client success")
}