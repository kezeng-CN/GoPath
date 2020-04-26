package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

//GlobalObj zinx全局参数
type GlobalObj struct {
	TCPServer     ziface.IServer
	Host          string
	TCPPort       int
	Name          string
	Version       string
	MaxPacketSize uint32
	MaxConn       int
}

//GlobalObject GlobalObj实例
var GlobalObject *GlobalObj

//Reload 读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TCPPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}

	GlobalObject.Reload()
}
