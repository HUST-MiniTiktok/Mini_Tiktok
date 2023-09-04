package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish/publishservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

type PublishClient struct {
	client publishservice.Client
}

func NewPublishClient() (publishClient *PublishClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}

	opts := []client.Option{
		client.WithResolver(r),
		client.WithMiddleware(kitex.CommonMiddleware),
		client.WithInstanceMW(kitex.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithConnectTimeout(100 * time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
	}
	if conf.GetConf().GetBool("tracer.enabled") {
		opts = append(opts, client.WithSuite(opentracing.NewDefaultClientSuite()))
	}

	c, err := publishservice.NewClient("publish", opts...)

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

func (c *PublishClient) GetPublishInfoByUserId(context context.Context, req *publish.GetPublishInfoByUserIdRequest) (resp *publish.GetPublishInfoByUserIdResponse, err error) {
	resp, err = c.client.GetPublishInfoByUserId(context, req)
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
