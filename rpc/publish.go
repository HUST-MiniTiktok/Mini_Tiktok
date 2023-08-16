package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish/publishservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	publishClient publishservice.Client
)

func init() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	publishClient, err = publishservice.NewClient("publish", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new publish client failed: %v", err)
	}
}

func PublishAction(context context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp, err = publishClient.PublishAction(context, req)
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

func PublishList(context context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = publishClient.PublishList(context, req)
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

func GetVideoById(context context.Context, req *publish.GetVideoByIdRequest) (resp *publish.GetVideoByIdResponse, err error) {
	resp, err = publishClient.GetVideoById(context, req)
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

func GetVideoByIdList(context context.Context, req *publish.GetVideoByIdListRequest) (resp *publish.GetVideoByIdListResponse, err error) {
	resp, err = publishClient.GetVideoByIdList(context, req)
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
