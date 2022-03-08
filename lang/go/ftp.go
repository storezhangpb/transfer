package transfer

import (
	`fmt`
	`io`
	`os`
	`path/filepath`
	`time`

	`github.com/jlaffaye/ftp`
	`github.com/patrickmn/go-cache`
)

var ftpCache = cache.New(expiration, purge)

func (f *Ftp) upload(dest string, src string, _ string) error {
	return f.do(func(client *ftp.ServerConn) (err error) {
		if srcFile, openErr := os.Open(src); nil != openErr {
			err = openErr
		} else {
			err = client.Stor(dest, srcFile)
		}

		return
	})
}

func (f *Ftp) download(src string, dest string, _ string) error {
	return f.do(func(client *ftp.ServerConn) (err error) {
		if err = client.ChangeDir(filepath.Dir(src)); nil != err {
			return
		}

		var rsp *ftp.Response
		if rsp, err = client.Retr(filepath.Ext(src)); nil != err {
			return
		}
		defer func() {
			_ = rsp.Close()
		}()

		var destFile *os.File
		if destFile, err = os.Create(dest); nil != err {
			return
		}
		_, err = io.Copy(destFile, rsp)

		return
	})
}

func (f *Ftp) do(fun func(client *ftp.ServerConn) (err error)) (err error) {
	if client, getErr := f.getClient(); nil != err {
		err = getErr
	} else {
		err = fun(client)
	}

	return
}

func (f *Ftp) getClient() (client *ftp.ServerConn, err error) {
	key := f.key()
	if cached, ok := ftpCache.Get(key); ok {
		client = cached.(*ftp.ServerConn)
	} else if client, err = f.newClient(); nil == err {
		ossCache.Set(key, client, cache.DefaultExpiration)
	}

	return
}

func (f *Ftp) newClient() (client *ftp.ServerConn, err error) {
	if client, err = ftp.Dial(f.Addr, ftp.DialWithTimeout(5*time.Second)); nil == err {
		err = client.Login(f.Username, f.Password)
	}

	return
}

func (f *Ftp) key() string {
	return fmt.Sprintf(`%s.%s`, f.Addr, f.Username)
}
