package main

import (
	"log"

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
	r, err := etcd.NewEtcdRegistry(conf.GetConf().GetStringSlice("registry.address"))
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}


	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
