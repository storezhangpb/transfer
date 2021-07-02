package transfer

const (
	// FileTypeHttp Http存储
	FileTypeHttp FileType = "http"
	// FileTypeOss 阿里云对象存储
	FileTypeOss FileType = "oss"
	// FileTypeCos 腾讯云对象存储
	FileTypeCos FileType = "cos"
	// FileTypeFtp Ftp存储
	FileTypeFtp FileType = "ftp"
	// FileTypeLocalFile 本地存储
	FileTypeLocalFile FileType = "local"
)

// FileType 文件类型
type FileType string
