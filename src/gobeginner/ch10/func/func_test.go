package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(int) int) func(int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret = ret + op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resource.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("err") // defer仍会执行
}
