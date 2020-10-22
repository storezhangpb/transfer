package transfer

import (
	`encoding/json`
	`io`
	`os`
	`path/filepath`
	`time`

	`github.com/jlaffaye/ftp`
	log `github.com/sirupsen/logrus`
)

// FTP Ftp存储
type FTP struct {
	// 地址
	Addr string `json:"addr"`
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
	// 客户端
	client *ftp.ServerConn
	// 是否已经初始化
	startup bool
}

func NewFTPFile(filename string, ftp FTP) File {
	return NewFile(FileTypeFtp, filename, ftp)
}

func (f FTP) init() (err error) {
	if !f.startup {
		if f.client, err = ftp.Dial(f.Addr, ftp.DialWithTimeout(5*time.Second)); err != nil {
			log.WithFields(log.Fields{
				"type":     "ftp",
				"addr":     f.Addr,
				"username": f.Username,
				"password": f.Password,
				"error":    err,
			}).Error("连接Ftp服务失败")
			return
		}

		if err = f.client.Login(f.Username, f.Password); err != nil {
			log.WithFields(log.Fields{
				"type":     "ftp",
				"addr":     f.Addr,
				"username": f.Username,
				"password": f.Password,
				"error":    err,
			}).Error("登录Ftp服务失败")
		}
	}

	return
}

func (f FTP) Upload(destFilename string, srcFilename string) (err error) {
	var srcFile *os.File

	if err = f.init(); nil != err {
		return
	}

	// 打开文件
	if srcFile, err = os.Open(srcFilename); nil != err {
		log.WithFields(log.Fields{
			"type":         "ftp",
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
			"error":        err,
		}).Error("打开上传文件失败")
		return
	}

	if err = f.client.Stor(destFilename, srcFile); nil != err {
		log.WithFields(log.Fields{
			"type":         "ftp",
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
			"error":        err,
		}).Error("上传文件失败")
	} else {
		log.WithFields(log.Fields{
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
		}).Error("上传文件成功")
	}

	return
}

func (f FTP) Download(srcFilename string, destFilename string) (err error) {
	var (
		rsp      *ftp.Response
		destFile *os.File
	)

	if err = f.init(); nil != err {
		return
	}

	if err = f.client.ChangeDir(filepath.Dir(srcFilename)); nil != err {
		return
	}

	if rsp, err = f.client.Retr(filepath.Ext(srcFilename)); err != nil {
		log.WithFields(log.Fields{
			"type":         "ftp",
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
			"error":        err,
		}).Error("下载文件失败")
		return
	} else {
		log.WithFields(log.Fields{
			"type":         "ftp",
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
		}).Debug("下载文件成功")
	}
	defer rsp.Close()

	if destFile, err = os.Create(destFilename); nil != err {
		log.WithFields(log.Fields{
			"type":         "ftp",
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
			"error":        err,
		}).Error("创建下载文件失败")
		return
	}

	if _, err = io.Copy(destFile, rsp); nil != err {
		log.WithFields(log.Fields{
			"type":         "ftp",
			"addr":         f.Addr,
			"username":     f.Username,
			"password":     f.Password,
			"srcFilename":  srcFilename,
			"destFilename": destFilename,
			"error":        err,
		}).Error("写入下载文件失败")
	}

	return
}

func (f FTP) String() string {
	jsonBytes, _ := json.MarshalIndent(f, "", "    ")

	return string(jsonBytes)
}
