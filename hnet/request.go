package hnet

import "hin/hiface"

type Request struct {
	// 已经和客户端建立好的连接
	conn hiface.IConnection
	// 客户端请求的数据
	data []byte
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() hiface.IConnection {
	return r.conn
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.data
}