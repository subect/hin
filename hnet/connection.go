package hnet

import (
	"fmt"
	"hin/hiface"
	"io"
	"net"
)

type Connection struct {
	// 当前连接的 socket TCP 套接字
	Conn *net.TCPConn
	// 连接的 ID
	ConnID uint32
	// 当前连接的状态
	isClosed bool

	// 该连接的处理方法 router
	Router hiface.IRouter

	// 告知当前连接已经退出/停止的 channel
	ExitChan chan bool
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router hiface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
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

		// 创建一个拆包解包的对象
		dp := NewDataPack()

		// 读取客户端的 Msg Head 二进制流 8 个字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read msg head error: ", err)
			c.ExitChan <- true
			continue
		}

		// 拆包，得到 msgID 和 msgDataLen 放在 msg 消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error: ", err)
			c.ExitChan <- true
			continue
		}

		// 根据 dataLen 再次读取 data，放在 msg.Data 中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read msg data error: ", err)
				c.ExitChan <- true
				continue
			}
		}
		msg.SetData(data)

		// 得到当前 conn 数据的 Request 请求数据

		req := Request{
			conn: c,
			msg:  msg,
		}

		// 从 router 中，找到注册绑定的 Conn 对应的 router 调用
		go func(request hiface.IRequest) {
			// 执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
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

// SendMsg 发送数据，将数据发送给远程的客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return fmt.Errorf("connection closed when send msg")
	}
	// 将 data 进行封包 MsgDataLen|MsgID|Data
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return fmt.Errorf("pack error msg id = %v", msgId)
	}

	// 将数据发送给客户端
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write msg id = ", msgId, " error: ", err)
		c.ExitChan <- true
		return fmt.Errorf("write msg id = %v, error: %v", msgId, err)
	}

	return nil
}
