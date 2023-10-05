package main

import (
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment/commentservice"
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
	if conf.GetConf().GetBool("tracer.enabled") {
		tracer.InitJaeger("comment")
	}

	dal.Init()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8885")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}

	opts := []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "comment"}),
		server.WithServiceAddr(addr),
		server.WithMiddleware(kitex.CommonMiddleware),
		server.WithMiddleware(kitex.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithLimit(&limit.Option{MaxConnections: 1000000, MaxQPS: 100000000}),
		server.WithRegistry(r),
	}
	if conf.GetConf().GetBool("tracer.enabled") {
		opts = append(opts, server.WithSuite(opentracing.NewDefaultServerSuite()))
	}

	svr := comment.NewServer(new(CommentServiceImpl), opts...)

	if err := svr.Run(); err != nil {
		klog.Errorf("comment server stopped with error:", err)
	} else {
		klog.Infof("comment server stopped")
	}
}
