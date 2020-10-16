package transfer

type (
	// Downloader 下载
	Downloader interface {
		// Download 下载文件
		Download(srcFilename string, destFilename string) (err error)
	}
)
