package transfer

// Ftp 文件传送服务
type Ftp struct {
	message

	Addr     string
	Username string
	Password string
}
