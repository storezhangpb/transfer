package transfer

import (
	`encoding/json`
	`os`

	`github.com/storezhang/gox`
)

// File 文件
type File struct {
	// Type 类型
	Type FileType `json:"type" validate:"required,oneof=http oss cos ftp local"`
	// Filename 文件名
	Filename string `json:"filename" validate:"required"`
	// Checksum 文件校验
	Checksum gox.Checksum `json:"checksum" validate:"omitempty,structonly"`
	// Storage 存储
	Storage interface{} `json:"storage"`
}

func NewFile(fileType FileType, filename string, storage interface{}, checksums ...gox.Checksum) File {
	file := File{
		Type:     fileType,
		Filename: filename,
		Storage:  storage,
	}
	if 1 == len(checksums) {
		file.Checksum = checksums[0]
	}

	return file
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
	case FileTypeOss:
		oss := Oss{}
		if err = json.Unmarshal(rawMsg, &oss); nil != err {
			return
		}
		f.Storage = oss
	case FileTypeFtp:
		ftp := Ftp{}
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
