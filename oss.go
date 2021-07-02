package transfer

import (
	`github.com/aliyun/aliyun-oss-go-sdk/oss`
)

var _ Transfer = (*Oss)(nil)

// Oss 阿里云Oss存储
type Oss struct {
	// 端点
	Endpoint string `json:"endpoint"`
	// 桶名称
	Bucket string `json:"bucket"`
	// 授权
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

// NewOssFile 创建一个阿里云对象存储文件
func NewOssFile(filename string, oss Oss) File {
	return NewFile(FileTypeOss, filename, oss)
}

func (o Oss) Upload(destFilename string, srcFilename string) (err error) {
	var bucket *oss.Bucket

	if bucket, err = o.getBucket(); nil != err {
		return
	}
	err = bucket.PutObjectFromFile(destFilename, srcFilename)

	return
}

func (o Oss) Download(srcFilename string, destFilename string) (err error) {
	var bucket *oss.Bucket

	if bucket, err = o.getBucket(); nil != err {
		return
	}
	err = bucket.GetObjectToFile(srcFilename, destFilename)

	return
}

func (o Oss) getBucket() (bucket *oss.Bucket, err error) {
	var client *oss.Client

	if client, err = oss.New(o.Endpoint, o.AccessKey, o.SecretKey); nil != err {
		return
	}
	bucket, err = client.Bucket(o.Bucket)

	return
}
