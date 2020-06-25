package znet

import "zinx/ziface"

//BaseRouter 基类
type BaseRouter struct{}

//PreHandle 预处理
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

//Handle 处理
func (br *BaseRouter) Handle(request ziface.IRequest) {}

//PostHandle 处理后
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
