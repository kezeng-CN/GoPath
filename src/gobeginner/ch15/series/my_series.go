package series

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// Square is .
func Square(n int) int {
	return n * n
}

// GetFabonacciSerie is .
func GetFabonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1]+ret[i-2])
	}
	return ret
}
