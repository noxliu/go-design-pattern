package template

import "fmt"

type Implement interface {
	download(url string)
	save(savePath string)
}

type Downloader interface {
	Download(url string, savePath string)
}

type Template struct {
	Implement Implement
}

func (t *Template) Download(url string, savePath string) {
	fmt.Println("prepare downloading")
	t.Implement.download(url)
	t.Implement.save(savePath)
	fmt.Println("finish downloading")
}
