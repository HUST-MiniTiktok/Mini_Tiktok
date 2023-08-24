package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user/userservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

type UserClient struct {
	client userservice.Client
}

func NewUserClient() (userClient *UserClient) {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err.Error())
	}
	c, err := userservice.NewClient("user",
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
		klog.Fatalf("new user client failed: %v", err.Error())
	}
	userClient = &UserClient{client: c}
	return
}

func (c *UserClient) User(context context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp, err = c.client.User(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err.Error())
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserClient) Login(context context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	klog.Warnf("login req: %v", req)
	resp, err = c.client.Login(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err.Error())
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserClient) Register(context context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	klog.Warnf("register req: %v", req)
	resp, err = c.client.Register(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err.Error())
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserClient) CheckUserIsExist(context context.Context, req *user.CheckUserIsExistRequest) (resp *user.CheckUserIsExistResponse, err error) {
	klog.Warnf("check user is exist req: %v", req)
	resp, err = c.client.CheckUserIsExist(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err.Error())
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}
