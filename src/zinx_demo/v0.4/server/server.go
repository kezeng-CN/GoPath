package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

//PingRouter 测试
type PingRouter struct {
	znet.BaseRouter
}

//Handle 处理
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
