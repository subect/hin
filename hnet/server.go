package hnet

import (
	"errors"
	"fmt"
	"hin/hiface"
	"hin/utils"
	"net"
)

// Server 服务器类
type Server struct {
	Name      string // 服务器名称
	IPVersion string // 服务器绑定的IP版本
	IP        string // 服务器监听的IP
	Port      int    // 服务器监听的端口
	Router    hiface.IRouter
}

// NewServer 创建一个服务器句柄
func NewServer() hiface.IServer {
	utils.GlobalObject.Reload()
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
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

		var cid uint32
		cid = 0

		//3.启动 server 网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}

			//TODO:设置服务器最大连接限制，如果超过最大连接数，则关闭新连接

			//3.2 处理新连接请求的业务方法，此时 conn 和 handle 应该是 一一绑定的
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前连接的处理业务
			go dealConn.Start()
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

func (s *Server) AddRouter(router hiface.IRouter) {
	s.Router = router
	fmt.Println("Add Router succ!")
}

// CallBackToClient 定义当前客户端连接所绑定的 handle api（目前这个 handle 是写死的，以后优化成可配置的）
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err: ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}
