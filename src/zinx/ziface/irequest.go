package ziface

//IRequest 请求链接 + 请求数据
type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
