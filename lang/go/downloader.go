package transfer

// Downloader 下载
type Downloader interface {
	// Download 下载文件
	Download(srcFilename string, destFilename string) (err error)
}
