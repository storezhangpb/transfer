package transfer

import (
	`io/ioutil`
	`net/http`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

var _ Transfer = (*Http)(nil)

// Http Http存储
type Http struct {
	// 访问地址
	Url string `json:"url"`
	// 方法
	Method gox.HttpMethod `json:"method"`
}

// NewHttpFile 创建Http存储文件
func NewHttpFile(filename string, url string, method gox.HttpMethod) File {
	return NewFile(FileTypeHttp, filename, Http{
		Url:    url,
		Method: method,
	})
}

func (h Http) Upload(_ string, srcFilename string) (err error) {
	var (
		req *resty.Request
		rsp *resty.Response

		contentType string
		fileBytes   []byte
	)

	if contentType, err = gox.ContentType(srcFilename); nil != err {
		return
	}
	if fileBytes, err = ioutil.ReadFile(srcFilename); nil != err {
		return
	}

	req = NewResty().SetBody(fileBytes).SetHeader(gox.HeaderContentType, contentType)
	switch h.Method {
	case gox.HttpMethodPut:
		rsp, err = req.Put(h.Url)
	case gox.HttpMethodGet:
		rsp, err = req.Get(h.Url)
	default:
		rsp, err = req.Put(h.Url)
	}

	if nil != err {
		return
	}

	if http.StatusOK != rsp.StatusCode() {
		err = ErrorUpload
	}

	return
}

func (h Http) Download(_ string, destFilename string) (err error) {
	var (
		req *resty.Request
		rsp *resty.Response
	)

	req = NewResty().SetOutput(destFilename)
	switch h.Method {
	case gox.HttpMethodPut:
		rsp, err = req.Put(h.Url)
	case gox.HttpMethodGet:
		rsp, err = req.Get(h.Url)
	default:
		rsp, err = req.Get(h.Url)
	}

	if nil != err {
		return
	}

	if http.StatusOK != rsp.StatusCode() {
		err = ErrorDownload
	}

	return
}
