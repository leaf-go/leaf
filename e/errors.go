package e

import (
	"github.com/leaf-go/x"
)

func init() {
	x.ErrorsInit(mappings)
}

const (
	retryMessage  = ",请稍后重试..."
	affirmMessage = ",请检查..."

	ParamsFailed x.InfoErrno = 801
	Executing    x.InfoErrno = 802
)

var mappings = x.Errors{
	ParamsFailed: "输入有误" + affirmMessage,
	Executing:    "请求正在执行中...",
}
