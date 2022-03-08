package transfer

import (
	`google.golang.org/protobuf/proto`
)

func (t *Target) Upload(src string) error {
	return t.do(func(transfer transfer) error {
		return transfer.upload(t.Filename, src, t.Base)
	})
}

func (t *Target) Download(dest string) error {
	return t.do(func(transfer transfer) error {
		return transfer.download(t.Filename, dest, t.Base)
	})
}

func (t *Target) do(fun func(transfer transfer) (err error)) (err error) {
	var storage proto.Message

	switch t.Type {
	case Type_COS:
		storage = new(Cos)
		err = t.Storage.UnmarshalTo(storage)
	}

	if nil == err {
		err = fun(storage.(transfer))
	}

	return
}
