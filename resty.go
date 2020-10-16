package transfer

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

// NewResty Resty客户端
func NewResty() *resty.Request {
	return resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).R()
}

// RestyStringBody 字符串形式的结果
func RestyStringBody(rsp *resty.Response) string {
	body := ""
	if nil != rsp {
		body = rsp.String()
	}

	return body
}
