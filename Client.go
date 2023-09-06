package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
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
