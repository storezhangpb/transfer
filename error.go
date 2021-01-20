package transfer

import (
	`github.com/storezhang/gox`
)

var (
	// ErrorNotSupportStorage 不支持的存储类型
	ErrorNotSupportStorage = &gox.CodeError{ErrorCode: 100, Message: "不支持的存储类型"}
	// ErrorDownload
	ErrorDownload = &gox.CodeError{ErrorCode: 101, Message: "文件下载失败"}
	// ErrorUpload
	ErrorUpload = &gox.CodeError{ErrorCode: 102, Message: "文件上传失败"}
)
