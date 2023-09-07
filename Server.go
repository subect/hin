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

	// 先读取客户端的数据，再回写 ping...ping...ping
	fmt.Println("recv from client: msgId = nil", ", data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
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
