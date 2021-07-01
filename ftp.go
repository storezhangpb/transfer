package transfer

import (
	`io`
	`os`
	`path/filepath`
	`time`

	`github.com/jlaffaye/ftp`
)

var _ Transfer = (*Ftp)(nil)

// Ftp Ftp存储
type Ftp struct {
	// 地址
	Addr string `json:"addr" validate:"required"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 客户端
	client *ftp.ServerConn
	// 是否已经初始化
	initialized bool
}

// NewFtpFile 创建一个Ftp类型的文件
func NewFtpFile(filename string, ftp Ftp) File {
	return NewFile(FileTypeFtp, filename, ftp)
}

func (f Ftp) init() (err error) {
	if f.initialized {
		return
	}

	if f.client, err = ftp.Dial(f.Addr, ftp.DialWithTimeout(5*time.Second)); err != nil {
		return
	}
	err = f.client.Login(f.Username, f.Password)

	return
}

func (f Ftp) Upload(destFilename string, srcFilename string) (err error) {
	if err = f.init(); nil != err {
		return
	}

	// 打开文件
	var srcFile *os.File
	if srcFile, err = os.Open(srcFilename); nil != err {
		return
	}

	if err = f.client.Stor(destFilename, srcFile); nil != err {
		err = ErrorUpload
	}

	return
}

func (f Ftp) Download(srcFilename string, destFilename string) (err error) {
	if err = f.init(); nil != err {
		return
	}
	if err = f.client.ChangeDir(filepath.Dir(srcFilename)); nil != err {
		return
	}

	var rsp *ftp.Response
	if rsp, err = f.client.Retr(filepath.Ext(srcFilename)); nil != err {
		return
	}
	defer func() {
		_ = rsp.Close()
	}()

	var destFile *os.File
	if destFile, err = os.Create(destFilename); nil != err {
		return
	}
	_, err = io.Copy(destFile, rsp)

	return
}
