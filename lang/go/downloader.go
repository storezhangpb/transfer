package transfer

// downloader 下载
type downloader interface {
	// Download 下载文件
	Download(src string, dest string, base string) (err error)
}
