package main

import (
	"fmt"
	"hin/hnet"
	"net"
)

// 客户端goroutine ，模拟粘包的数据，然后进行发送
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	// 创建一个封包对象 dp
	dp := hnet.NewDataPack()

	// 封装一个msg1包
	msg1 := &hnet.Message{
		Id:      1,
		DataLen: 3,
		Data:    []byte{'h', 'i', 'n'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
		return
	}

	// 封装一个msg2包
	msg2 := &hnet.Message{
		Id:      2,
		DataLen: 11,
		Data:    []byte{'h', 'e', 'l', 'l', 'o', ',', 'w', 'o', 'r', 'l', 'd'},
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err:", err)
		return
	}

	// 将sendData1，和 sendData2 拼接一起，组成粘包
	sendData1 = append(sendData1, sendData2...)

	// 一次性发送给服务端
	conn.Write(sendData1)

	// 客户端阻塞
	select {}

}
