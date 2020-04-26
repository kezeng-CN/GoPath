package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/utils"
	"zinx/ziface"
)

//Server 是对iServer的实现
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
}

//Start 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPacketSize)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
		}

		fmt.Println("start Zinx server ", s.Name, " succ, now listenning...")

		var cid uint32
		cid = 0

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

//Stop 停止服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name", s.Name)
}

//Serve 调用Start
func (s *Server) Serve() {
	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

//AddRouter 添加路由
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Addr Router succ! ")
}

//NewServer 返回实例
func NewServer() ziface.IServer {
	utils.GlobalObject.Reload()

	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TCPPort,
		Router:    nil,
	}
	return s
}
