package transfer

// Uploader 上传
type Uploader interface {
	// Upload 上传文件
	Upload(destFilename string, srcFilename string) (err error)
}
