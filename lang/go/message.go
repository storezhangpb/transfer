package transfer

import (
	`google.golang.org/protobuf/reflect/protoreflect`
)

type message struct{}

func (m *message) ProtoReflect() protoreflect.Message {
	panic(`占位，不要修改`)
}
