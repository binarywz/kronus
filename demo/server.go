package main

import "binary/wz/kronus/znet"

func main() {
	// 创建一个Server句柄
	server := znet.NewServer("[kronus v0.1]")
	// 启动server
	server.Serve()
}