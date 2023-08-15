package oss

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)


func MakeBucket(ctx context.Context, bucketName string) {
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

func PutToBucketByBuffer(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = Client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return
}

func PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = Client.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return
}
