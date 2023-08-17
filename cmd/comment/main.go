package main

import (
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	dal.Init()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8885")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "comment"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatalf("run Comment rpc server failed: %v", err)
	}
}
