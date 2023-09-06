package utils

import (
	"encoding/json"
	"fmt"
	"hin/hiface"
	"io/ioutil"
)

// 存储一切有关 Zinx 框架的全局参数，供其他模块使用
// 一些参数是可以通过 zinx.json 由用户进行配置

var GlobalObject *GlobalObj

type GlobalObj struct {
	TcpServer  hiface.IServer // 全局 Server 对象
	Host       string         `json:"host"`       // 服务器主机监听的 IP
	TcpPort    int            `json:"port"`       // 服务器主机监听的端口号
	Name       string         `json:"name"`       // 服务器的名称
	Version    string         `json:"version"`    // 版本号
	MaxPackage uint32         `json:"maxPackage"` // 数据包的最大值
	MaxConn    int            `json:"maxConn"`    // 服务器主机允许的最大连接数
}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/hin.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:       "HServerApp",
		Version:    "V0.1",
		TcpPort:    7777,
		Host:       "",
		MaxPackage: 4096,
		MaxConn:    12000,
	}
}
