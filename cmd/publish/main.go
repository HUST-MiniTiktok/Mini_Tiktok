package main

import (
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish/publishservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	dal.Init()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8883")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}

	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "publish"}),
		server.WithServiceAddr(addr),
		server.WithMiddleware(kitex.CommonMiddleware),
		server.WithMiddleware(kitex.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	if err := svr.Run(); err != nil {
		klog.Errorf("publish server stopped with error:", err)
	} else {
		klog.Infof("publish server stopped")
	}
}
