package transfer

import (
	`context`
	`crypto/tls`
	`fmt`
	`net/http`
	`net/url`

	`github.com/patrickmn/go-cache`

	`github.com/tencentyun/cos-go-sdk-v5`
)

var cosCache = cache.New(expiration, purge)

func (c *Cos) upload(dest string, src string, base string) error {
	return c.do(func(client *cos.Client) (err error) {
		ctx := context.Background()
		filepath := path(base, c.Separator, dest)
		_, _, err = client.Object.Upload(ctx, filepath, src, nil)

		return
	})
}

func (c *Cos) download(src string, dest string, base string) error {
	return c.do(func(client *cos.Client) (err error) {
		ctx := context.Background()
		filepath := path(base, c.Separator, src)
		_, err = client.Object.Download(ctx, filepath, dest, nil)

		return
	})
}

func (c *Cos) do(fun func(client *cos.Client) (err error)) (err error) {
	if client, getErr := c.getClient(); nil != err {
		err = getErr
	} else {
		err = fun(client)
	}

	return
}

func (c *Cos) getClient() (client *cos.Client, err error) {
	key := c.key()
	if cached, ok := cosCache.Get(key); ok {
		client = cached.(*cos.Client)
	} else if client, err = c.newClient(); nil == err {
		cosCache.Set(key, client, cache.DefaultExpiration)
	}

	return
}

func (c *Cos) newClient() (client *cos.Client, err error) {
	if bucketUrl, parseErr := url.Parse(c.Url); nil != parseErr {
		err = parseErr
	} else {
		client = cos.NewClient(&cos.BaseURL{BucketURL: bucketUrl}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:     c.Id,
				SecretKey:    c.Key,
				SessionToken: c.Token,
				// nolint:gosec
				Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			},
		})
	}

	return
}

func (c *Cos) key() string {
	return fmt.Sprintf(`%s.%s`, c.Url, c.Id)
}
