package main

import (
	"net"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/dal"
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation/relationservice"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/kitex"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	dal.Init()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8887")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}

	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "relation"}),
		server.WithServiceAddr(addr),
		server.WithMiddleware(kitex.CommonMiddleware),
		server.WithMiddleware(kitex.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	if err := svr.Run(); err != nil {
		klog.Errorf("relation server stopped with error:", err)
	} else {
		klog.Infof("relation server stopped")
	}
}
