package znet

import (
	"fmt"
	"net"

	"binary/wz/kronus/ziface"
)

// Server IServer的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的IP版本
	IpVersion string
	// 服务器绑定的IP
	IP string
	// 服务器监听的端口
	Port int
	// 服务器路由
	Router ziface.IRouter
}

func (server *Server) Start() {
	fmt.Printf("[Start] Server listen at IP: %s, Port: %d, start...\n", server.IP, server.Port)
	// 1.获取一个TCP的addr
	addr, err := net.ResolveTCPAddr(server.IpVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}
	// 2.监听服务器的地址
	listener, err := net.ListenTCP(server.IpVersion, addr)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	fmt.Println("Server start success...")
	var connId uint32 = 0

	// 3.阻塞的等待客户端连接，处理客户端连接业务
	for {
		// 若客户端连接，阻塞会返回
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}
		handleConn := NewConn(conn, connId, server.Router)
		handleConn.Start()
		connId++
	}
}

func (server *Server) Stop() {

}

func (server *Server) Serve() {
	// 启动server的服务功能
	server.Start()
	// 阻塞状态
	select {}
}

func (server *Server) RegisterRouter(router ziface.IRouter) {
	server.Router = router
	fmt.Println("Register router success...")
}

// NewServer 初始化Server
// golang中接口类型是引用
func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      name,
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return server
}
