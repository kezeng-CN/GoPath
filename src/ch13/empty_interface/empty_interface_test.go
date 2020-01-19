package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomeThing(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// } else if s, ok := p.(string); ok {
	// 	fmt.Println("String", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomeThing(10)
	DoSomeThing("10")
}
