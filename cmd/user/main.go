package main

import (
	"log"
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal"
	user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user/userservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	dal.Init()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8881")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}
	
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
