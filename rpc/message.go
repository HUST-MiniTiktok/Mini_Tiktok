package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/kitex_gen/message/messageservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	messageClient messageservice.Client
)

func init() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	messageClient, err = messageservice.NewClient("message", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new message client failed: %v", err)
	}
}

func MessageAction(context context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	resp, err = messageClient.MessageAction(context, req)
	if err != nil {
		klog.Errorf("message client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("message client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("message client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func MessageChat(context context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	resp, err = messageClient.MessageChat(context, req)
	if err != nil {
		klog.Errorf("message client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("message client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("message client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}