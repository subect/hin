package hiface

// IServer 定义服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)
}
