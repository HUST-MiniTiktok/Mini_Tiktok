package oss

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)

func CreateBucket(ctx context.Context, bucketName string) {
	exists, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		klog.Fatalf("check bucket exists failed: %v", err)
		return
	}
	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			klog.Fatalf("create bucket failed: %v", err)
			return
		}
		klog.Infof("create bucket %s success", bucketName)
	}
}

func PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	defer fileObj.Close()
	info, err = Client.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	return
}

func GetObjectURL(ctx context.Context, bucketName, filename string) (obj_url *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	obj_url, err = Client.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return
}

func PutToBucketWithBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = Client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return
}

func ToRealURL(ctx context.Context, db_url string) (real_url string) {
	names := strings.Split(db_url, "/")
	bucket_name := names[0]
	file_name := names[1]
	real_url_, err := GetObjectURL(ctx, bucket_name, file_name)
	if err != nil {
		klog.Errorf("get object url failed: %v", err)
	} else {
		real_url_.Host = conf.GetConf().GetString("oss.endpoint")
		real_url = real_url_.String()
	}
	return
}

func ToDbURL(ctx context.Context, bucket_name string, file_name string) (db_url string) {
	db_url = bucket_name + "/" + file_name
	return
}
