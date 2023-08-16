package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/kitex_gen/comment"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/kitex_gen/comment/commentservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	commentClient commentservice.Client
)

func init() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	commentClient, err = commentservice.NewClient("comment", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new comment client failed: %v", err)
	}
}

func CommentAction(context context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	resp, err = commentClient.CommentAction(context, req)
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

func CommentList(context context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp, err = commentClient.CommentList(context, req)
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