package main

import (
	"github.com/leaf-go/x"
	"leaf/boot"
)

func main() {
	// 启动项加载
	boot.Application()

	// 这里需要做信号通知结束服务。因为有些服务本身就是协程的.
	x.NewService().Mounts("http.api").Boot()
	//x.NewService().Mounts("http.api").With("http.task").Boot()
	//x.NewService().Mounts("http.api").With("http.task").Boot()
}
