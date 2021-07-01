package transfer

import (
	`github.com/storezhang/gox`
)

var _ Transfer = (*LocalFile)(nil)

// LocalFile 本地文件
type LocalFile struct{}

func NewLocalFile(filename string) File {
	return File{
		Type:     FileTypeLocalFile,
		Filename: filename,
		Storage:  LocalFile{},
	}
}

func (lf LocalFile) Upload(_ string, _ string) (err error) {
	return
}

func (lf LocalFile) Download(srcFilename string, destFilename string) (err error) {
	_, err = gox.CopyFile(srcFilename, destFilename)

	return
}
