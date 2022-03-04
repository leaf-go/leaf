// @Description  描述
// @Author  姓名
// @Update  2020-10-27 19:49
package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

//todo 创建分布式ID 这里可以精读一下
var Snowflake *defaultSnowflake

func init() {
	ip, _ := Net.LocalIP()
	inode := int64(ip[2])<<8 + int64(ip[3])
	node, err := snowflake.NewNode(inode % 1024)
	if err != nil {
		panic(fmt.Sprintf("snowflake not created err:%s, %s", ip, err.Error()))
	}
	Snowflake = newDefaultSnowflake(node)
}

type defaultSnowflake struct {
	node *snowflake.Node
}

func newDefaultSnowflake(node *snowflake.Node) *defaultSnowflake {
	return &defaultSnowflake{node: node}
}

func (d *defaultSnowflake) NextID() int64 {
	return d.node.Generate().Int64()
}
