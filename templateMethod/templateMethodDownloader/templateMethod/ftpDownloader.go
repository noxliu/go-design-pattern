package templateMethod

import "fmt"

type FtpDownloader struct {
}

func (*FtpDownloader) download(url string) {
	fmt.Println("ftp download from ", url)
}

func (*FtpDownloader) save(savePath string) {
	fmt.Println("ftp save file to ", savePath)
}
