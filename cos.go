package transfer

import (
	`context`
	`crypto/tls`
	`net/http`
	`net/url`

	`github.com/tencentyun/cos-go-sdk-v5`
)

func (c *Cos) Upload(destFilename string, srcFilename string) (err error) {
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
			SecretID:     c.Id,
			SecretKey:    c.Key,
			SessionToken: c.Token,
			// nolint:gosec
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	})

	return
}
