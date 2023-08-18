package main

import (
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed/feedservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	dal.Init()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8882")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}

	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "feed"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	if err := svr.Run(); err != nil {
		klog.Errorf("feed server stopped with error:", err)
	} else {
		klog.Infof("feed server stopped")
	}
}
