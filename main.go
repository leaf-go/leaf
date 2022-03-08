package main

import (
	_ "leaf-go/boot"
	"x"
	//"leaf-go/leaf/x"
)

func main() {
	// env ??
	x.InitEnv("debug", "release")

	// 启动？ 怎么启动？
	x.NewService().Mounts("http.api").With("http.task").Boot()
}
