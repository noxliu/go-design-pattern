package templateMethod

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
	fmt.Println("start download")
	t.Implement.download(url)
	t.Implement.save(savePath)
	fmt.Println("finish download")
}
