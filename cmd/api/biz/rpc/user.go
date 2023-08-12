package rpc

import (
	"context"
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user/userservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	userClient userservice.Client
)

func init() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err)
	}
	userClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("new user client failed: %v", err)
	}
}

func User(context context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp, err = userClient.User(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func Login(context context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse,err error) {
	resp, err = userClient.Login(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

func Register(context context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse,err error) {
	resp, err = userClient.Register(context, req)
	if err != nil {
		klog.Errorf("user client failed: %v", err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		klog.Errorf("user client failed: %v -> %v", resp.StatusCode, resp.StatusMsg)
		return nil, fmt.Errorf("user client failed: %v", resp.StatusMsg)
	}
	return resp, nil
}

