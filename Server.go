package main

import "hin/hnet"

//Server 模块的测试函数
func main() {
	s := hnet.NewServer("[HServer V0.1]")
	s.Serve()
}
