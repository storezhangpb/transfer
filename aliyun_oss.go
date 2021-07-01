package transfer

import (
	`github.com/aliyun/aliyun-oss-go-sdk/oss`
)

var _ Transfer = (*AliyunOss)(nil)

// AliyunOss 阿里云Oss存储
type AliyunOss struct {
	// 端点
	EndPoint string `json:"endPoint"`
	// 桶名称
	Bucket string `json:"bucket"`
	// 授权
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

// NewAliyunOssFile 创建一个阿里云文件
func NewAliyunOssFile(filename string, oss AliyunOss) File {
	return NewFile(FileTypeAliyunOss, filename, oss)
}

func (ao AliyunOss) getBucket() (bucket *oss.Bucket, err error) {
	var client *oss.Client

	if client, err = oss.New(ao.EndPoint, ao.AccessKey, ao.SecretKey); nil != err {
		return
	}
	bucket, err = client.Bucket(ao.Bucket)

	return
}

func (ao AliyunOss) Upload(destFilename string, srcFilename string) (err error) {
	var bucket *oss.Bucket

	if bucket, err = ao.getBucket(); nil != err {
		return
	}
	err = bucket.PutObjectFromFile(destFilename, srcFilename)

	return
}

func (ao AliyunOss) Download(srcFilename string, destFilename string) (err error) {
	var bucket *oss.Bucket

	if bucket, err = ao.getBucket(); nil != err {
		return
	}
	err = bucket.GetObjectToFile(srcFilename, destFilename)

	return
}
