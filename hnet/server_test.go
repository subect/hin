package hnet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

//模拟客户端
func ClientTest() {
	fmt.Println("ClientTest...start")
	// 3s 后 发起连接
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client tart err,exit!")
		return
	}
	for {
		_, err := conn.Write([]byte("hello hnet"))
		if err != nil {
			fmt.Println("write conn err:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			return
		}

		fmt.Println("server call back:", string(buf[:cnt]))
		time.Sleep(1 * time.Second)
	}
}

//Server 模块的测试
func TestServer(t *testing.T) {
	//1.创建一个 server 句柄
	s := NewServer("[HServer V0.1]")

	//客户端测试
	go ClientTest()

	//2.启动 server
	s.Serve()
}