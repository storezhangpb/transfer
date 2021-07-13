package transfer

import (
	`github.com/storezhang/gox`
)

var _ Transfer = (*Local)(nil)

// Local 本地文件
type Local struct{}

func NewLocal(filename string) File {
	return File{
		Type:     FileTypeLocalFile,
		Filename: filename,
		Storage:  Local{},
	}
}

func (l *Local) Upload(_ string, _ string) (err error) {
	return
}

func (l *Local) Download(srcFilename string, destFilename string) (err error) {
	_, err = gox.CopyFile(srcFilename, destFilename)

	return
}
