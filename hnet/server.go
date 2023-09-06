package hnet

import (
	"fmt"
	"hin/hiface"
	"net"
)

// Server 服务器类
type Server struct {
	Name      string // 服务器名称
	IPVersion string // 服务器绑定的IP版本
	IP        string // 服务器监听的IP
	Port      int    // 服务器监听的端口
}

// NewServer 创建一个服务器句柄
func NewServer(name string) hiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
	return s
}

func (s *Server) Start() {
	fmt.Println("[Start] Server name: ", s.Name, ", listenner at IP: ", s.IP, ", Port: ", s.Port)

	//开启一个 goroutine 去处理 Linster 业务
	go func() {
		// 1.获取一个 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		//2.监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " error: ", err)
			return
		}

		fmt.Println("start HServer succ, ", s.Name, " succ, Listening...")

		//3.启动 server 网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}

			//TODO:设置服务器最大连接限制，如果超过最大连接数，则关闭新连接

			//TODO：处理新连接请求的业务方法，此时 conn 和 handle 应该是 一一绑定的

			//客户端建立连接，做一些业务，做一个最基本的最大512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err: ", err)
						continue
					}

					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err: ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[Stop] HServer name: ", s.Name)
	//TODO:将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或者回收
}

func (s *Server) Serve() {
	s.Start()
	//TODO:启动服务的时候初始化一些事情

	//阻塞，否则主 Go 会退出，listener 的 go将会退出
	select {}
}
