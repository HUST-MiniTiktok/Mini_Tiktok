package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type CommentClient struct {
	client commentservice.Client
}

func NewCommentClient() (commentClient *CommentClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := commentservice.NewClient("comment", client.WithResolver(r))
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
