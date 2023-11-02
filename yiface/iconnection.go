package yiface

import "net"

type IConnection interface {
	// Start 启动连接，让当前的链接开始工作
	Start()
	// Stop 停止连接，结束当前连接的工作
	Stop()

	// GetTCPConnection 返回当前连接
	GetTCPConnection() *net.TCPConn
	// GetConnID 返回连接id
	GetConnID() uint32
	// RemoteAddr 获取远端地址
	RemoteAddr() net.Addr

	// SendMsg 通过无缓冲管道发送数据方法
	SendMsg(msgId uint32, data []byte) error
	// SendBuffMsg 通过有缓冲管道发送数据方法
	SendBuffMsg(msgId uint32, data []byte) error

	// AddProperty 添加属性
	AddProperty(string, interface{})
	// RemoveProperty 删除属性
	RemoveProperty(string)
	// GetProperty 获得属性值
	GetProperty(string) (interface{}, error)
}

type HandleFunc func(*net.TCPConn, []byte, int) error
