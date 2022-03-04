package utils

import (
	"errors"
	"net"
)

var (
	Net = newDefaultNet()
)

//init 初始化
func init() {
	Net.init()
}

//defaultNet 工具类：网络
type defaultNet struct {
	ip  net.IP
	err error
}

//newDefaultNet 构造函数
func newDefaultNet() *defaultNet {
	return &defaultNet{}
}

//init 初始化网络
func (n *defaultNet) init() {
	defer n.deferGetIP()
	_, _ = n.getLocalIP()
}

//getLocalIP 获取本地ip
func (n *defaultNet) getLocalIP() (ip net.IP, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		n.err = err
		return
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ipv4 := ip.IP.To4(); ipv4 != nil {
				n.ip = ip.IP.To4()
				break
			}
		}
	}

	return n.ip, n.err
}

//LocalIP 获取本地ip。
func (n *defaultNet) LocalIP() (ip net.IP, err error) {
	return n.ip, n.err
}

//deferGetIP 延迟处理ip错误。
func (n *defaultNet) deferGetIP() {
	if n.ip == nil {
		n.ip = net.IPv4(127, 0, 0, 1)
		if n.err == nil {
			n.err = errors.New("fail to guess local ip")
		}
	}
}
