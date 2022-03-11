package main

import (
	"github.com/leaf-go/x"
	"leaf/boot"
)

func main() {
	// 启动项加载
	boot.Application()

	x.NewService().Mounts("http.api").Boot()
	//x.NewService().Mounts("http.api").With("http.task").Boot()
	//x.NewService().Mounts("http.api").With("http.task").Boot()
}
