package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite/favoriteservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type FavoriteClient struct {
	client favoriteservice.Client
}

func NewFavoriteClient() (favoriteClient *FavoriteClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := favoriteservice.NewClient("favorite", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new favorite client failed: %v", err)
	}
	favoriteClient = &FavoriteClient{client: c}
	return
}

func (c *FavoriteClient) FavoriteAction(context context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp, err = c.client.FavoriteAction(context, req)
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

func (c *FavoriteClient) FavoriteList(context context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp, err = c.client.FavoriteList(context, req)
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

func (c *FavoriteClient) GetVideoFavoriteCount(context context.Context, req *favorite.GetVideoFavoriteCountRequest) (resp *favorite.GetVideoFavoriteCountResponse, err error) {
	resp, err = c.client.GetVideoFavoriteCount(context, req)
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

func (c *FavoriteClient) CheckIsFavorite(context context.Context, req *favorite.CheckIsFavoriteRequest) (resp *favorite.CheckIsFavoriteResponse, err error) {
	resp, err = c.client.CheckIsFavorite(context, req)
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