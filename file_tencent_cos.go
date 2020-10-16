package transfer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storezhang/gox"
)

// TencentCos 腾讯云Cos存储
type TencentCos struct {
	// 预签名地址
	PreassignedURL string `json:"preassignedURL"`
	// 方法
	Method gox.HttpMethod `json:"method"`
}

func NewTencentCosFile(filename string, cos TencentCos) File {
	return NewFile(StorageTypeTencentCos, filename, cos)
}

func (tc TencentCos) Upload(destFilename string, srcFilename string) (err error) {
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
		rsp, err = req.Put(tc.PreassignedURL)
	case gox.HttpMethodGet:
		rsp, err = req.Get(tc.PreassignedURL)
	default:
		rsp, err = req.Put(tc.PreassignedURL)
	}

	if nil != err {
		log.WithFields(log.Fields{
			"preassignedURL": tc.PreassignedURL,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("上传文件出错")

		return
	}

	if http.StatusOK != rsp.StatusCode() {
		log.WithFields(log.Fields{
			"preassignedURL": tc.PreassignedURL,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("上传文件失败")
	} else {
		log.WithFields(log.Fields{
			"preassignedURL": tc.PreassignedURL,
			"fileKey":        destFilename,
			"filename":       srcFilename,
		}).Debug("上传文件成功")
	}

	return
}

func (tc TencentCos) Download(srcFilename string, destFilename string) (err error) {
	var (
		req *resty.Request
		rsp *resty.Response
	)

	req = NewResty().SetOutput(destFilename)
	switch tc.Method {
	case gox.HttpMethodPut:
		rsp, err = req.Put(tc.PreassignedURL)
	case gox.HttpMethodGet:
		rsp, err = req.Get(tc.PreassignedURL)
	default:
		rsp, err = req.Get(tc.PreassignedURL)
	}

	if nil != err {
		log.WithFields(log.Fields{
			"preassignedURL": tc.PreassignedURL,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("下载文件出错")

		return
	}

	if http.StatusOK != rsp.StatusCode() {
		log.WithFields(log.Fields{
			"preassignedURL": tc.PreassignedURL,
			"fileKey":        destFilename,
			"filename":       srcFilename,
			"error":          err,
		}).Error("下载文件失败")
	} else {
		log.WithFields(log.Fields{
			"preassignedURL": tc.PreassignedURL,
			"fileKey":        destFilename,
			"filename":       srcFilename,
		}).Debug("下载文件成功")
	}

	return
}

func (tc TencentCos) String() string {
	jsonBytes, _ := json.MarshalIndent(tc, "", "    ")

	return string(jsonBytes)
}
