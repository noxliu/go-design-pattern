package proxy

import "fmt"

type fileAccessorImpl struct {
}

func (*fileAccessorImpl) ReadFile(filePath string) {
	fmt.Println("访问文件操作")
}
