package main

import (
	"leaf-go/boot"
	_ "leaf-go/boot"
	"x"
	//"leaf-go/leaf/x"
)

func main() {
	// 启动项加载
	boot.Application()

	// 这里需要做信号通知结束服务。因为有些服务本身就是协程的.
	x.NewService().Mounts("http.api").With("http.task").Boot()
}
