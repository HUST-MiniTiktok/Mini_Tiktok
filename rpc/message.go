package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message/messageservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type MessageClient struct {
	client messageservice.Client
}

func NewMessageClient() (messageClient *MessageClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := messageservice.NewClient("message", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new message client failed: %v", err)
	}
	messageClient = &MessageClient{client: c}
	return
}

func (c *MessageClient) MessageAction(context context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	resp, err = c.client.MessageAction(context, req)
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

func (c *MessageClient) MessageChat(context context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	resp, err = c.client.MessageChat(context, req)
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