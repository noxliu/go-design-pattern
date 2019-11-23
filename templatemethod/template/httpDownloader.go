package template

import "fmt"

type HttpDownloader struct {
}

func (*HttpDownloader) download(url string) {
	fmt.Println("http download from ", url)
}

func (*HttpDownloader) save(savePath string) {
	fmt.Println("http save file to ", savePath)
}
