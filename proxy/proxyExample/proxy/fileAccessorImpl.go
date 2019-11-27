package proxy

import "fmt"

type fileAccessorImpl struct {
}

func (*fileAccessorImpl) ReadFile(filePath string) {
	fmt.Println("access file")
}
