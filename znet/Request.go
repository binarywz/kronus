package znet

import "binary/wz/kronus/ziface"

// Request 封装客户端请求
type Request struct {
	// 当前请求绑定的连接
	conn ziface.IConn
	// 客户端请求参数
	param []byte
}

// GetConn 获取请求绑定的连接
func (request *Request) GetConn() ziface.IConn {
	return request.conn
}

// GetParam 获取请求参数
func (request *Request) GetParam() []byte {
	return request.param
}
