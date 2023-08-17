package oss

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/mw/redis"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	OSSClient *minio.Client
	RDClient  *redis.RDClient
	err       error
)

func init() {
	RDClient = redis.NewRDClient()
	ctx := context.Background()
	OSSClient, err = minio.New(conf.GetConf().GetString("oss.endpoint"), &minio.Options{
		Creds: credentials.NewStaticV4(conf.GetConf().GetString("oss.accesskey"), conf.GetConf().GetString("oss.secretkey"), ""),
	})

	if err != nil {
		klog.Fatalf("init minio client failed: %v", err)
	}

	CreateBucket(ctx, conf.GetConf().GetString("oss.videobucket"))
	CreateBucket(ctx, conf.GetConf().GetString("oss.imagebucket"))
}
