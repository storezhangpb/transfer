package transfer

import (
	`encoding/json`
	`os`

	`github.com/storezhang/gox`
)

const (
	// 存储类型
	// FileTypeHttp Http存储
	FileTypeHttp FileType = "http"
	// FileTypeAliyunOss 阿里云Oss
	FileTypeAliyunOss FileType = "oss"
	// FileTypeFtp Ftp存储
	FileTypeFtp FileType = "ftp"
	// FileTypeLocalFile 本地存储
	FileTypeLocalFile FileType = "local"
)

type (
	// FileType 文件类型
	FileType string

	// File 文件
	File struct {
		// Type 类型
		Type FileType `json:"type" validate:"required,oneof=http oss ftp local"`
		// Filename 文件名
		Filename string `json:"filename" validate:"required"`
		// Storage 存储
		Storage interface{} `json:"storage"`
	}
)

func NewFile(fileType FileType, filename string, storage interface{}) File {
	return File{
		Type:     fileType,
		Filename: filename,
		Storage:  storage,
	}
}

func (f *File) Upload(filename string) (err error) {
	return f.Storage.(Uploader).Upload(f.Filename, filename)
}

func (f *File) Download(filename string, force bool) (err error) {
	if force && gox.IsFileExist(filename) {
		if err = os.Remove(filename); nil != err {
			return
		}
	}
	err = f.Storage.(Downloader).Download(f.Filename, filename)

	return
}

func (f *File) UnmarshalJSON(data []byte) (err error) {
	type cloneType File

	rawMsg := json.RawMessage{}
	f.Storage = &rawMsg

	if err = json.Unmarshal(data, (*cloneType)(f)); nil != err {
		return
	}

	switch f.Type {
	case FileTypeHttp:
		cos := Http{}
		if err = json.Unmarshal(rawMsg, &cos); nil != err {
			return
		}
		f.Storage = cos
	case FileTypeAliyunOss:
		oss := AliyunOss{}
		if err = json.Unmarshal(rawMsg, &oss); nil != err {
			return
		}
		f.Storage = oss
	case FileTypeFtp:
		ftp := FTP{}
		if err = json.Unmarshal(rawMsg, &ftp); nil != err {
			return
		}
		f.Storage = ftp
	case FileTypeLocalFile:
		local := LocalFile{}
		if err = json.Unmarshal(rawMsg, &local); nil != err {
			return
		}
		f.Storage = local
	default:
		err = ErrorNotSupportStorage
	}

	return
}

func (f File) String() string {
	jsonBytes, _ := json.MarshalIndent(f, "", "    ")

	return string(jsonBytes)
}
