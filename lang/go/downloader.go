package transfer

type downloader interface {
	download(src string, dest string, base string) (err error)
}
