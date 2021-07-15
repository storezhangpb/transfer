package transfer

import (
	`context`
	`crypto/tls`
	`net/http`
	`net/url`

	`github.com/tencentyun/cos-go-sdk-v5`
	_ `github.com/tencentyun/cos-go-sdk-v5`
)

var _ Transfer = (*Cos)(nil)

// Cos 腾讯云对象存储
type Cos struct {
	// 通信地址
	Url string
	// 基础路径
	Base string
	// 授权，相当于用户名
	SecretId string
	// 授权，相当于密码
	SecretKey string
	//  临时密钥
	Token string
	// 分隔符
	Separator string `default:"/"`
}

// NewCosFile 创建一个腾讯云对象存储文件
func NewCosFile(filename string, cos Cos) File {
	return NewFile(FileTypeCos, filename, cos)
}

func (c Cos) Upload(destFilename string, srcFilename string) (err error) {
	var client *cos.Client

	if client, err = c.getClient(); nil != err {
		return
	}
	_, _, err = client.Object.Upload(context.Background(), path(c.Base, c.Separator, destFilename), srcFilename, nil)

	return
}

func (c *Cos) Download(srcFilename string, destFilename string) (err error) {
	var client *cos.Client

	if client, err = c.getClient(); nil != err {
		return
	}
	_, err = client.Object.Download(context.Background(), path(c.Base, c.Separator, srcFilename), destFilename, nil)

	return
}

func (c *Cos) getClient() (client *cos.Client, err error) {
	var bucketUrl *url.URL
	if bucketUrl, err = url.Parse(c.Url); nil != err {
		return
	}

	client = cos.NewClient(&cos.BaseURL{BucketURL: bucketUrl}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:     c.SecretId,
			SecretKey:    c.SecretKey,
			SessionToken: c.Token,
			// nolint:gosec
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	})

	return
}
