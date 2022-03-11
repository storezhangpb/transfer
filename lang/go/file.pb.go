package transfer

import (
	`github.com/golang/protobuf/ptypes/any`
)

// File 不要删除此文件，只是占位
type File struct {
	Name string
	Base string
	Type Type

	Storage *any.Any
}
