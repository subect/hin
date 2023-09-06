package hiface

// IRequest 把客户端请求的连接信息和请求数据包装到一个 Request 中
type IRequest interface {
	// GetConnection 得到当前连接
	GetConnection() IConnection
	// GetData 获取请求体的数据
	GetData() []byte
}
