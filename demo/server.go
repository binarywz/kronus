package main

import (
	"binary/wz/kronus/ziface"
	"binary/wz/kronus/znet"
	"fmt"
)

type PingRouter struct {
	znet.AbstractRouter
}

func (router *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("PingRouter PreHandle...")
	_, err := request.GetConn().GetConn().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("PingRouter PreHandle error:", err)
	}
}

func (router *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("PingRouter Handle...")
	_, err := request.GetConn().GetConn().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("PingRouter Handle error:", err)
	}
}

func (router *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("PingRouter PostHandle...")
	_, err := request.GetConn().GetConn().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("PingRouter PostHandle error:", err)
	}
}

func main() {
	// 创建一个Server句柄
	server := znet.NewServer()
	server.RegisterRouter(&PingRouter{})
	// 启动server
	server.Serve()
}
