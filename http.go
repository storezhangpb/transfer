package transfer

import (
	`encoding/json`
	`io/ioutil`
	`net/http`

	`github.com/go-resty/resty/v2`
	log `github.com/sirupsen/logrus`
	`github.com/storezhang/gox`
)

var _ Transfer = (*Http)(nil)

// Http Http存储
type Http struct {
	// Url 访问地址
	Url string `json:"url"`
	// Method 方法
	Method gox.HttpMethod `json:"method"`
}

// NewHttpFile 创建Http存储文件
func NewHttpFile(filename string, url string, method gox.HttpMethod) File {
	return NewFile(FileTypeHttp, filename, Http{
		Url:    url,
		Method: method,
	})
}

func (tc Http) Upload(destFilename string, srcFilename string) (err error) {
	var (
		req *resty.Request
		rsp *resty.Response

		contentType string
		fileBytes   []byte
	)

	if contentType, err = gox.GetContentType(srcFilename); nil != err {
		return
	}
	if fileBytes, err = ioutil.ReadFile(srcFilename); nil != err {
		return
	}

	req = NewResty().SetBody(fileBytes).SetHeader(gox.HeaderContentType, contentType)
	switch tc.Method {
	case gox.HttpMethodPut:
		rsp, err = req.Put(tc.Url)
	case gox.HttpMethodGet:
		rsp, err = req.Get(tc.Url)
	default:
		rsp, err = req.Put(tc.Url)
	}

	if nil != err {
		log.WithFields(log.Fields{
			"preassignedURL": tc.Url,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("上传文件出错")

		return
	}

	if http.StatusOK != rsp.StatusCode() {
		log.WithFields(log.Fields{
			"preassignedURL": tc.Url,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("上传文件失败")

		err = ErrorUpload
	} else {
		log.WithFields(log.Fields{
			"preassignedURL": tc.Url,
			"fileKey":        destFilename,
			"filename":       srcFilename,
		}).Debug("上传文件成功")
	}

	return
}

func (tc Http) Download(srcFilename string, destFilename string) (err error) {
	var (
		req *resty.Request
		rsp *resty.Response
	)

	req = NewResty().SetOutput(destFilename)
	switch tc.Method {
	case gox.HttpMethodPut:
		rsp, err = req.Put(tc.Url)
	case gox.HttpMethodGet:
		rsp, err = req.Get(tc.Url)
	default:
		rsp, err = req.Get(tc.Url)
	}

	if nil != err {
		log.WithFields(log.Fields{
			"preassignedURL": tc.Url,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("下载文件出错")

		return
	}

	if http.StatusOK != rsp.StatusCode() {
		log.WithFields(log.Fields{
			"preassignedURL": tc.Url,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("下载文件失败")

		err = ErrorDownload
	} else {
		log.WithFields(log.Fields{
			"preassignedURL": tc.Url,
			"fileKey":        destFilename,
			"filename":       srcFilename,
		}).Debug("下载文件成功")
	}

	return
}

func (tc Http) String() string {
	jsonBytes, _ := json.MarshalIndent(tc, "", "    ")

	return string(jsonBytes)
}
