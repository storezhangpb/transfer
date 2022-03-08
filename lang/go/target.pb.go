package transfer

import (
	`github.com/golang/protobuf/ptypes/any`
)

// Target 不要删除此文件，只是占位
type Target struct {
	Filename string
	Base     string
	Type     Type

	Storage *any.Any
}
