package main

import (
	"fmt"
	"hin/hiface"
	"hin/hnet"
)

type PingRouter struct {
	hnet.BaseRouter
}

func (pr *PingRouter) Handle(request hiface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("handle ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
		return
	}
}

//Server 模块的测试函数
func main() {
	s := hnet.NewServer()

	s.AddRouter(&PingRouter{})

	s.Serve()
}
