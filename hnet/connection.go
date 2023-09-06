package hnet

import (
	"fmt"
	"hin/hiface"
	"net"
)

type Connection struct {
	// 当前连接的 socket TCP 套接字
	Conn *net.TCPConn
	// 连接的 ID
	ConnID uint32
	// 当前连接的状态
	isClosed bool
	// 告知当前连接已经退出/停止的 channel
	ExitChan chan bool
	// 该连接的处理方法 API
	handleAPI hiface.HandleFunc
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi hiface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
		handleAPI: callbackApi,
	}
	return c
}

// Start 启动连接
func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	// 启动从当前连接的读数据的业务
	go c.StartReader()

	for {
		select {
		case <-c.ExitChan:
			// 代表 Reader 已经退出，此时 Writer 也要退出
			return
		}
	}
}

// Stop 停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//TODO: 如果用户注册了该连接的关闭回调业务，那么在此刻应该显示调用

	// 关闭 socket 连接
	c.Conn.Close()

	//	通知从缓冲队列读数据的业务，该连接已经关闭
	c.ExitChan <- true

	// 回收资源
	close(c.ExitChan)
}

// StartReader 处理 conn 读数据的 Goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到 buf 中，最大 512 字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err: ", err)
			c.ExitChan <- true
			continue
		}

		// 调用当前连接所绑定的 HandleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID ", c.ConnID, " handle is error")
			c.ExitChan <- true
			break
		}
	}
}

// GetTCPConnection 从当前连接获取原始的 socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接 ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端的 TCP 状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

