package transfer

type uploader interface {
	upload(dest string, src string, base string) (err error)
}
