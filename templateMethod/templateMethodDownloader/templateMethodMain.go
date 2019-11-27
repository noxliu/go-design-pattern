package main

import (
	. "go-design-pattern/templateMethod/templateMethodDownloader/templateMethod"
)

func main() {
	template := Template{}

	ftpDownloader := FtpDownloader{}
	template.Implement = &ftpDownloader
	template.Download("http://www.xxx.com/a.zip", "/opt/local")

	httpDownloader := HttpDownloader{}
	template.Implement = &httpDownloader
	template.Download("http://www.xxx.com/a.zip", "/opt/local")
}
