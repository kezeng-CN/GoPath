package znet

import "zinx/ziface"

//Request 数据请求结构
type Request struct {
	conn ziface.IConnection
	data []byte
}

//GetConnection 请求连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//GetData 请求数据
func (r *Request) GetData() []byte {
	return r.data
}
