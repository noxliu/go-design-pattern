package main

import (
	. "go-design-pattern/templatemethod/template"
)

func main() {

	ftpDownloader := FtpDownloader{}
	template := Template{}
	template.Implement = &ftpDownloader
	template.Download("", "")

	httpDownloader := HttpDownloader{}
	template.Implement = &httpDownloader
	template.Download("", "")
}
