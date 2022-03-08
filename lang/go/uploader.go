package transfer

type uploader interface {
	Upload(dest string, src string, base string) (err error)
}
