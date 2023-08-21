package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation/relationservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type RelationClient struct {
	client relationservice.Client
}

func NewRelationClient() (relationClient *RelationClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := relationservice.NewClient("relation",
		client.WithResolver(r),
		client.WithMiddleware(kitex.CommonMiddleware),
		client.WithInstanceMW(kitex.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(5*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
	)
	if err != nil {
		klog.Fatalf("new relation client failed: %v", err)
	}
	relationClient = &RelationClient{client: c}
	return
}

func (c *RelationClient) RelationAction(context context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp, err = c.client.RelationAction(context, req)
	if err != nil {
		klog.Errorf("relation client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("relation client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("relation client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) RelationFollowList(context context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp, err = c.client.RelationFollowList(context, req)
	if err != nil {
		klog.Errorf("relation client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("relation client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("relation client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) RelationFollowerList(context context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp, err = c.client.RelationFollowerList(context, req)
	if err != nil {
		klog.Errorf("relation client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("relation client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("relation client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) RelationFriendList(context context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	resp, err = c.client.RelationFriendList(context, req)
	if err != nil {
		klog.Errorf("relation client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("relation client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("relation client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) GetFollowInfo(context context.Context, req *relation.GetFollowInfoRequest) (resp *relation.GetFollowInfoResponse, err error) {
	resp, err = c.client.GetFollowInfo(context, req)
	if err != nil {
		klog.Errorf("relation client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("relation client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("relation client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}
