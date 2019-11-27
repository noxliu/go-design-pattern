package main

import (
	. "go-design-pattern/proxy/proxyExample/proxy"
)

func main() {
	//fileAccessorImpl := fileAccessorImpl{} 用户不能访问到 fileAccessorImpl, 利用包权限隔离，只能通过下面的代理实现访问
	userFileAccessProxy := FileAccessorProxy{UserType: "user"}
	guestFileAccessProxy := FileAccessorProxy{UserType: "guest"}
	userFileAccessProxy.ReadFile("/opt/user/file")
	guestFileAccessProxy.ReadFile("/opt/user/file")
}
