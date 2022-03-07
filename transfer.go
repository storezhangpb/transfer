package transfer

// Transfer 数据交换
type Transfer interface {
	Uploader
	Downloader
}
