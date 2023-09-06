package main

import (
	"fmt"
	"hin/hiface"
	"hin/hnet"
)

type PingRouter struct {
	hnet.BaseRouter
}

func (pr *PingRouter) PreHandle(request hiface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
		return
	}
}

func (pr *PingRouter) Handle(request hiface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("handle ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
		return
	}
}

func (pr *PingRouter) PostHandle(request hiface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
		return
	}
}

//Server 模块的测试函数
func main() {
	s := hnet.NewServer("[HServer V0.1]")

	s.AddRouter(&PingRouter{})

	s.Serve()
}
