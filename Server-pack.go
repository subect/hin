package main

import (
	"fmt"
	"hin/hnet"
	"io"
	"net"
)

func main() {
	// 创建 socket TCP 客户端
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// 创建服务器goroutine ，负责从客户端读取黏包数据，然后进行解析
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
		}

		go func(conn net.Conn) {
			// 创建封包拆包对象 dp
			dp := hnet.NewDataPack()
			for {
				// 先读出流的 head 部分
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Printf("server unpack err:%v", err.Error())
					return
				}
				// 将 headData字节流 拆包到 msg中
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Printf("server unPack err:%v", err.Error())
				}

				// 根据 headData 中的 dataLen 再次读取 data，放在 msg.Data 中
				if msgHead.GetDataLen() > 0 {
					// msg 有 data 数据，需要再次读取 data 数据
					msg := msgHead.(*hnet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Printf("server unpack data err:%v", err.Error())
						return
					}
					fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
				}
			}
		}(conn)
	}

}
