package znet

import (
	"binary/wz/kronus/ziface"
	"fmt"
	"net"
)

// Conn 连接
type Conn struct {
	// Socket连接字
	Conn *net.TCPConn

	// 连接ID
	ConnId uint32

	// 当前连接的状态
	isClosed bool

	// 当前连接绑定的业务处理函数
	handler ziface.EventHandler

	// 通知连接状态channel
	ExitChan chan bool
}

func NewConn(conn *net.TCPConn, connId uint32, handler ziface.EventHandler) *Conn {
	c := &Conn{
		Conn:     conn,
		ConnId:   connId,
		handler:  handler,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (conn *Conn) Start() {
	fmt.Println("Conn Start(), ConnId:", conn.ConnId)
	// 启动当前连接读数据的业务
	go conn.read()
}

func (conn *Conn) read() {
	fmt.Println("Conn read event start...")
	defer fmt.Println("read end, connId:", conn.ConnId, "remote addr:", conn.RemoteAddr())
	defer conn.Stop()

	for {
		// 读取客户端数据
		buf := make([]byte, 512)
		length, readErr := conn.Conn.Read(buf)
		if readErr != nil {
			fmt.Println("recv buf error:", readErr)
			break
		}
		// 调用当前连接绑定的EventHandler
		if handlerErr := conn.handler(conn.Conn, buf, length); handlerErr != nil {
			fmt.Println("handle occur error:", handlerErr, ", connId:", conn.ConnId)
			break
		}
	}
}

func (conn *Conn) Stop() {
	fmt.Println("Conn Stop(), ConnId:", conn.ConnId)
	if conn.isClosed == true {
		return
	}
	// 关闭Socket连接
	conn.Conn.Close()
	// 回收channel
	close(conn.ExitChan)
	conn.isClosed = true
}

func (conn *Conn) GetConn() *net.TCPConn {
	return conn.Conn
}

func (conn *Conn) GetConnId() uint32 {
	return conn.ConnId
}

func (conn *Conn) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

func (conn *Conn) Send(buf []byte) error {
	_, err := conn.Conn.Write(buf)
	return err
}
