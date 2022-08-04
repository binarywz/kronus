package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// 连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("connect server error:", err)
		return
	}
	// 调用Write 写数据
	for {
		if _, err := conn.Write([]byte("Hello Golang.")); err != nil {
			fmt.Println("conn write error:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("recv data:%s, len:%d\n", buf[:cnt], cnt)
		time.Sleep(1 * time.Second)
	}
}
