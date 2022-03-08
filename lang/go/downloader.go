package transfer

type downloader interface {
	Download(src string, dest string, base string) (err error)
}
