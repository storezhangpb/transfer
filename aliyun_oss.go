package transfer

import (
	`encoding/json`

	`github.com/aliyun/aliyun-oss-go-sdk/oss`
	log `github.com/sirupsen/logrus`
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

	if client, err = oss.New(ao.EndPoint, ao.AccessKey, ao.SecretKey); err != nil {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"error":    err,
		}).Error("创建Oss客户端失败")

		return
	}

	if bucket, err = client.Bucket(ao.Bucket); err != nil {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"error":    err,
		}).Error("创建Bucket失败")
	}

	return
}

func (ao AliyunOss) Upload(destFilename string, srcFilename string) (err error) {
	var bucket *oss.Bucket

	if bucket, err = ao.getBucket(); err != nil {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"fileKey":  destFilename,
			"filename": srcFilename,
			"error":    err,
		}).Error("创建Bucket对象失败")

		return
	}

	if err = bucket.PutObjectFromFile(destFilename, srcFilename); err != nil {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"fileKey":  destFilename,
			"filename": srcFilename,
			"error":    err,
		}).Error("上传文件失败")

		err = ErrorUpload
	} else {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"fileKey":  destFilename,
			"filename": srcFilename,
		}).Debug("上传文件成功")
	}

	return
}

func (ao AliyunOss) Download(srcFilename string, destFilename string) (err error) {
	var bucket *oss.Bucket

	if bucket, err = ao.getBucket(); err != nil {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"fileKey":  srcFilename,
			"filename": destFilename,
			"error":    err,
		}).Error("创建Bucket对象失败")

		return
	}

	if err = bucket.GetObjectToFile(srcFilename, destFilename); err != nil {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"fileKey":  srcFilename,
			"filename": destFilename,
			"error":    err,
		}).Error("下载文件失败")
	} else {
		log.WithFields(log.Fields{
			"endPoint": ao.EndPoint,
			"bucket":   ao.Bucket,
			"fileKey":  srcFilename,
			"filename": destFilename,
		}).Debug("下载文件成功")
	}

	return
}

func (ao AliyunOss) String() string {
	jsonBytes, _ := json.MarshalIndent(ao, "", "    ")

	return string(jsonBytes)
}
