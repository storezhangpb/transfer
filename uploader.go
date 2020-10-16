package transfer

type (
	// Uploader 上传
	Uploader interface {
		// Upload 上传文件
		Upload(destFilename string, srcFilename string) (err error)
	}
)
