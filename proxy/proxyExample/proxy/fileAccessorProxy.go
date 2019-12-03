package proxy

import (
	"fmt"
)

type FileAccessorProxy struct {
	UserType string
}

func (f *FileAccessorProxy) ReadFile(filePath string) {
	if f.UserType == "user" {
		fmt.Printf("当前用户: %s \n", f.UserType)
		fileAccessor := fileAccessorImpl{}
		fileAccessor.ReadFile(filePath)
	} else if f.UserType == "guest" {
		fmt.Printf("当前用户: %s \n", f.UserType)
		fmt.Println("对不起，你没有访问权限")
	}
}
