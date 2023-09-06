package hiface

import "net"

// IConnection 定义连接接口
type IConnection interface {
	// Start 启动连接，让当前连接开始工作
	Start()
	// Stop 停止连接，结束当前连接状态
	Stop()
	// GetTCPConnection 获取当前连接绑定的 socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前连接模块的连接 ID
	GetConnID() uint32
	// RemoteAddr 获取远程客户端的 TCP 状态 IP port
	RemoteAddr() net.Addr
}

// HandleFunc 定义一个统一处理连接业务的接口
type HandleFunc func(*net.TCPConn, []byte, int) error
