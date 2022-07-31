package ziface

// IRequest 封装客户端请求
type IRequest interface {
	// GetConn 获取当前连接
	GetConn() IConn
	// GetParam 获取请求参数
	GetParam() []byte
}
