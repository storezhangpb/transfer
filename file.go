package transfer

import (
	"encoding/json"
	"os"

	"github.com/storezhang/gox"
)

const (
	// 存储类型
	// StorageTypeTencentCos 腾讯云Cos
	StorageTypeTencentCos StorageType = "cos"
	// StorageTypeAliyunOss 阿里云Oss
	StorageTypeAliyunOss StorageType = "oss"
	// StorageTypeFTP FTP存储
	StorageTypeFTP StorageType = "ftp"
	// StorageTypeLocalFile 本地存储
	StorageTypeLocalFile StorageType = "local"
)

type (
	// StorageType 存储类型
	StorageType string

	// File 文件
	File struct {
		// StorageType 类型
		StorageType StorageType `json:"storageType" validate:"required,oneof=cos oss ftp"`
		// Filename 文件名
		Filename string `json:"filename" validate:"required"`
		// Storage 存储
		Storage interface{} `json:"storage"`
	}
)

func NewFile(storageType StorageType, filename string, storage interface{}) File {
	return File{
		StorageType: storageType,
		Filename:    filename,
		Storage:     storage,
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

	switch f.StorageType {
	case StorageTypeTencentCos:
		cos := TencentCos{}
		if err = json.Unmarshal(rawMsg, &cos); nil != err {
			return
		}
		f.Storage = cos
	case StorageTypeAliyunOss:
		oss := AliyunOss{}
		if err = json.Unmarshal(rawMsg, &oss); nil != err {
			return
		}
		f.Storage = oss
	case StorageTypeFTP:
		ftp := FTP{}
		if err = json.Unmarshal(rawMsg, &ftp); nil != err {
			return
		}
		f.Storage = ftp
	case StorageTypeLocalFile:
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
