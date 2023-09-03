package main

import (
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish/publishservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	limit "github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

func main() {
	tracer.InitJaeger("publish")
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
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 100000}),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)

	if err := svr.Run(); err != nil {
		klog.Errorf("publish server stopped with error:", err)
	} else {
		klog.Infof("publish server stopped")
	}
}
