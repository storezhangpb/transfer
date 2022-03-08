package transfer

// Oss 不要删除此文件，只是占位
type Oss struct {
	message

	Endpoint string
	Bucket   string
	Access   string
	Secret   string

	Separator string
}
