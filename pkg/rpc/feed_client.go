package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed/feedservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type FeedClient struct {
	client feedservice.Client
}

func NewFeedClient() (feedClient *FeedClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := feedservice.NewClient("feed", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new feed client failed: %v", err)
	}
	feedClient = &FeedClient{client: c}
	return
}

func (c *FeedClient) GetFeed(context context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp, err = c.client.GetFeed(context, req)
	if err != nil {
		klog.Errorf("feed client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("feed client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("feed client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}
