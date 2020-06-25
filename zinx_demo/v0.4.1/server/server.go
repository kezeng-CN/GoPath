package main

import (
	"fmt"
	"io"
	"net"
	"zinx/znet"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("server accept err:", err)
		}

		go func(conn net.Conn) {
			dp := znet.NewDataPack()
			for {
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)

				if err != nil {
					fmt.Println("read head error")
					break
				}

				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack err:", err)
					return
				}

				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())

					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack data err:", err)
						return
					}

					fmt.Println("==> Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
				}
			}
		}(conn)
	}
}
