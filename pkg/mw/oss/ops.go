package oss

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)

func CreateBucket(ctx context.Context, bucketName string) {
	exists, err := OSSClient.BucketExists(ctx, bucketName)
	if err != nil {
		klog.Fatalf("check bucket exists failed: %v", err)
		return
	}
	if !exists {
		err = OSSClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
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
	info, err = OSSClient.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	return
}

func GetObjectURL(ctx context.Context, bucketName, filename string) (obj_url *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	obj_url, err = OSSClient.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return
}

func PutToBucketWithBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = OSSClient.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return
}

func PutToBucketWithFilePath(ctx context.Context, bucketName, filename, filePath string) (info minio.UploadInfo, err error) {
	info, err = OSSClient.FPutObject(ctx, bucketName, filename, filePath, minio.PutObjectOptions{})
	return
}

func ToRealURL(ctx context.Context, db_url string) (real_url string) {
	if RDClient.Exists(db_url) {
		return RDClient.Get(db_url)
	}
	names := strings.Split(db_url, "/")
	bucket_name := names[0]
	file_name := names[1]
	real_url_, err := GetObjectURL(ctx, bucket_name, file_name)
	go func() {
		RDClient.Set(db_url, real_url_, time.Hour*24)
	}()
	if err != nil {
		klog.Errorf("get object url failed: %v", err)
	} else {
		real_url_.Host = MinioHost
		real_url = real_url_.String()
	}
	return
}

func ToDbURL(bucket_name string, file_name string) (db_url string) {
	db_url = bucket_name + "/" + file_name
	return
}

func LoadDefaultImageData(ctx context.Context, image_bucket string) (err error) {
	_, filename, _, _ := runtime.Caller(0)
	default_image_dir := fmt.Sprintf("%s/default_data/%s/", path.Dir(filename), image_bucket)
	// 上传默认头像
	_, err = PutToBucketWithFilePath(ctx, "image", "Avatar.png", default_image_dir+"Avatar.png")
	if err != nil {
		klog.Fatalf("upload default avatar failed: %v", err)
		return
	}
	klog.Infof("upload default avatar success: %v")
	// 上传默认背景图
	_, err = PutToBucketWithFilePath(ctx, "image", "Background.png", image_bucket+"Background.png")
	if err != nil {
		klog.Fatalf("upload default background image failed: %v", err)
		return
	}
	klog.Infof("upload default background image success: %v")
	return
}
