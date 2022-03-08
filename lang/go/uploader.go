package transfer

// uploader 上传
type uploader interface {
	// Upload 上传文件
	Upload(dest string, src string, base string) (err error)
}
