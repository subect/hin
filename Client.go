package main

import (
	"fmt"
	"hin/hnet"
	"io"
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

		// 发送封包的 msg 消息
		dp := hnet.NewDataPack()
		msg, _ := dp.Pack(hnet.NewMsgPackage(1, []byte("hin v0.5 Client Test Message")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		// 先读出流中的 head 部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head error")
			break
		}

		// 将 headData 字节流 拆包到 msg 中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			// msg 有 data 数据，需要再次读取 data 数据
			msg := msgHead.(*hnet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			// 根据 dataLen 从 io 中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}
}
