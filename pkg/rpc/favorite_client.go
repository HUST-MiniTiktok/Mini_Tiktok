package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite/favoriteservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

type FavoriteClient struct {
	client favoriteservice.Client
}

func NewFavoriteClient() (favoriteClient *FavoriteClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := favoriteservice.NewClient("favorite",
		client.WithResolver(r),
		client.WithSuite(opentracing.NewDefaultClientSuite()),
		client.WithMiddleware(kitex.CommonMiddleware),
		client.WithInstanceMW(kitex.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(5*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
	)
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

func (c *FavoriteClient) GetVideoFavoriteInfo(context context.Context, req *favorite.GetVideoFavoriteInfoRequest) (resp *favorite.GetVideoFavoriteInfoResponse, err error) {
	resp, err = c.client.GetVideoFavoriteInfo(context, req)
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

func (c *FavoriteClient) GetUserFavoriteInfo(context context.Context, req *favorite.GetUserFavoriteInfoRequest) (resp *favorite.GetUserFavoriteInfoResponse, err error) {
	resp, err = c.client.GetUserFavoriteInfo(context, req)
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
