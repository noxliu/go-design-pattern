package main

import (
	"./templateMethod"
)

func main() {
	template := templateMethod.Template{}

	ftpDownloader := templateMethod.FtpDownloader{
		UserName: "testUser",
		PassWord: "password",
	}
	template.Implement = &ftpDownloader
	template.Download("ftp://www.xxx.com/a.zip", "/opt/local")

	httpDownloader := templateMethod.HttpDownloader{}
	template.Implement = &httpDownloader
	template.Download("http://www.xxx.com/a.zip", "/opt/local")
}
