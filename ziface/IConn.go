package ziface

import "net"

// IConn 定义连接模块的抽象层
type IConn interface {
	// Start 启动连接，当前连接准备开始工作
	Start()
	// Stop 停止连接，结束当前连接的工作
	Stop()
	// GetConn 获取当前连接绑定的socket conn
	GetConn() *net.TCPConn
	// GetConnId 获取当前连接的ID
	GetConnId() uint32
	// RemoteAddr 获取远程客户端的TCP状态、IP、port
	RemoteAddr() net.Addr
	// Send 发送数据，将数据发送给远程的客户端
	Send(buf []byte) error
}

// EventHandler 连接所绑定的处理业务的函数类型
type EventHandler func(*net.TCPConn, []byte, int) error
