package hnet

import "hin/hiface"

type BaseRouter struct{}

// PreHandle 处理业务之前的钩子方法
func (br *BaseRouter) PreHandle(request hiface.IRequest) {}

// Handle 处理业务的主方法
func (br *BaseRouter) Handle(request hiface.IRequest) {}

// PostHandle 处理业务之后的钩子方法
func (br *BaseRouter) PostHandle(request hiface.IRequest) {}
