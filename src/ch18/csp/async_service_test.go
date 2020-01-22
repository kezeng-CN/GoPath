package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	fmt.Println("service on")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("service off")
	return "Done"
}

func otherTask() {
	fmt.Println("something on")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("something off")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func asyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("async on")
		retCh <- ret
		fmt.Println("async off")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
}
