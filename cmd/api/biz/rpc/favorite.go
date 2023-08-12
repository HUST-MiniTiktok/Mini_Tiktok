package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/kitex_gen/favorite"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	favoriteClient favoriteservice.Client
)

func init() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	favoriteClient, err = favoriteservice.NewClient("favorite", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new favorite client failed: %v", err)
	}
}

func FavoriteAction(context context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp, err = favoriteClient.FavoriteAction(context, req)
	if err != nil {
		klog.Errorf("favorite client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("favorite client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("favorite client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteList(context context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp, err = favoriteClient.FavoriteList(context, req)
	if err != nil {
		klog.Errorf("favorite client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("favorite client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("favorite client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}