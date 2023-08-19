package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish/publishservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type PublishClient struct {
	client publishservice.Client
}

func NewPublishClient() (publishClient *PublishClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := publishservice.NewClient("publish", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new publish client failed: %v", err)
	}
	publishClient = &PublishClient{client: c}
	return
}

func (c *PublishClient) PublishAction(context context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp, err = c.client.PublishAction(context, req)
	if err != nil {
		klog.Errorf("publish client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("publish client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("publish client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *PublishClient) PublishList(context context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = c.client.PublishList(context, req)
	if err != nil {
		klog.Errorf("publish client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("publish client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("publish client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *PublishClient) GetVideoById(context context.Context, req *publish.GetVideoByIdRequest) (resp *publish.GetVideoByIdResponse, err error) {
	resp, err = c.client.GetVideoById(context, req)
	if err != nil {
		klog.Errorf("publish client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("publish client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("publish client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *PublishClient) GetVideoByIdList(context context.Context, req *publish.GetVideoByIdListRequest) (resp *publish.GetVideoByIdListResponse, err error) {
	resp, err = c.client.GetVideoByIdList(context, req)
	if err != nil {
		klog.Errorf("publish client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("publish client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("publish client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}
