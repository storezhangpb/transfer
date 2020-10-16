package transfer

import (
	"github.com/storezhang/gox"
)

var (
	// ErrorNotSupportStorage 不支持的存储类型
	ErrorNotSupportStorage = &gox.CodeError{ErrorCode: 100, Msg: "不支持的存储类型"}
)
