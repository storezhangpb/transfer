package transfer

import (
	`google.golang.org/protobuf/proto`
)

func (f *File) Upload(src string) error {
	return f.do(func(transfer transfer) error {
		return transfer.upload(f.Name, src, f.Base)
	})
}

func (f *File) Download(dest string) error {
	return f.do(func(transfer transfer) error {
		return transfer.download(f.Name, dest, f.Base)
	})
}

func (f *File) do(fun func(transfer transfer) (err error)) (err error) {
	var storage proto.Message

	switch f.Type {
	case Type_COS:
		storage = new(Cos)
		err = f.Storage.UnmarshalTo(storage)
	}

	if nil == err {
		err = fun(storage.(transfer))
	}

	return
}
