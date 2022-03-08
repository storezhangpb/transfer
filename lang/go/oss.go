package transfer

import (
	`fmt`

	`github.com/aliyun/aliyun-oss-go-sdk/oss`
	`github.com/patrickmn/go-cache`
)

var ossCache = cache.New(expiration, purge)

func (o *Oss) upload(dest string, src string, _ string) error {
	return o.do(func(bucket *oss.Bucket) error {
		return bucket.PutObjectFromFile(dest, src)
	})
}

func (o *Oss) download(src string, dest string, _ string) error {
	return o.do(func(bucket *oss.Bucket) error {
		return bucket.GetObjectToFile(src, dest)
	})
}

func (o *Oss) do(fun func(bucket *oss.Bucket) (err error)) (err error) {
	if bucket, getErr := o.getBucket(); nil != err {
		err = getErr
	} else {
		err = fun(bucket)
	}

	return
}

func (o *Oss) getBucket() (bucket *oss.Bucket, err error) {
	key := o.key()
	if cached, ok := ossCache.Get(key); ok {
		bucket = cached.(*oss.Bucket)
	} else if bucket, err = o.newBucket(); nil == err {
		ossCache.Set(key, bucket, cache.DefaultExpiration)
	}

	return
}

func (o *Oss) newBucket() (bucket *oss.Bucket, err error) {
	if client, newErr := oss.New(o.Endpoint, o.Access, o.Secret); nil != newErr {
		err = newErr
	} else {
		bucket, err = client.Bucket(o.Bucket)
	}

	return
}

func (o *Oss) key() string {
	return fmt.Sprintf(`%s.%s`, o.Endpoint, o.Access)
}
