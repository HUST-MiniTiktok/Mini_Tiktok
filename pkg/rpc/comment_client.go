package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment/commentservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

type CommentClient struct {
	client commentservice.Client
}

func NewCommentClient() (commentClient *CommentClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := commentservice.NewClient("comment",
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
		klog.Fatalf("new comment client failed: %v", err)
	}
	commentClient = &CommentClient{client: c}
	return
}

func (c *CommentClient) CommentAction(context context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	resp, err = c.client.CommentAction(context, req)
	if err != nil {
		klog.Errorf("comment client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("comment client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("comment client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *CommentClient) CommentList(context context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp, err = c.client.CommentList(context, req)
	if err != nil {
		klog.Errorf("comment client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("comment client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("comment client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *CommentClient) GetVideoCommentCount(context context.Context, req *comment.GetVideoCommentCountRequest) (resp *comment.GetVideoCommentCountResponse, err error) {
	resp, err = c.client.GetVideoCommentCount(context, req)
	if err != nil {
		klog.Errorf("comment client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("comment client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("comment client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}
