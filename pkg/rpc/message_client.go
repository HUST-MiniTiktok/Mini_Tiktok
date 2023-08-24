package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/message/messageservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

type MessageClient struct {
	client messageservice.Client
}

func NewMessageClient() (messageClient *MessageClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	c, err := messageservice.NewClient("message",
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

func (c *MessageClient) GetFriendLatestMsg(context context.Context, req *message.GetFriendLatestMsgRequest) (resp *message.GetFriendLatestMsgResponse, err error) {
	resp, err = c.client.GetFriendLatestMsg(context, req)
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
