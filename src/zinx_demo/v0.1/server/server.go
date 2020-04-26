package main

import (
	"zinx/znet"
)

func main() {
	s := znet.NewServer("name 0.1")
	s.Serve()
}
