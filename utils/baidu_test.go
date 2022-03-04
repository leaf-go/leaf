// @Description  描述
// @Author  姓名
// @Update  2020-10-27 20:16
package utils

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestBaidu_CheckText(t *testing.T) {
	baidu := Baidu{}
	baidu.CheckText(&gin.Context{}, "不要侮辱伟大")
}
