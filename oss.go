package transfer

import (
	`github.com/aliyun/aliyun-oss-go-sdk/oss`
)

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
