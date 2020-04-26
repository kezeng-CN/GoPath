package ziface

//IRouter 处理业务方法
type IRouter interface {
	PreHandle(request IRequest)  //预处理
	Handle(request IRequest)     //处理
	PostHandle(request IRequest) //处理后
}
