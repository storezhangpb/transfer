package transfer

import (
	`io`
	`os`
	`path/filepath`
	`time`

	`github.com/jlaffaye/ftp`
)

func (f *Ftp) init() (err error) {
	if f.initialized {
		return
	}

	if f.client, err = ftp.Dial(f.Addr, ftp.DialWithTimeout(5*time.Second)); err != nil {
		return
	}
	err = f.client.Login(f.Username, f.Password)

	return
}

func (f *Ftp) Upload(destFilename string, srcFilename string) (err error) {
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

func (f *Ftp) Download(srcFilename string, destFilename string) (err error) {
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
