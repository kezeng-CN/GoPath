package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should not be less than 2 ")
var GreaterThanHunderedError = errors.New("n should not be larger than 100 ")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, GreaterThanHunderedError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(2); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number")
		}
		if err == GreaterThanHunderedError {
			t.Error("Need a smaller number")
		}
	} else {
		t.Log(v)
	}
	GetFibonacci1("100")
	GetFibonacci2("100")
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("error", err)
		}
	} else {
		fmt.Println("error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(list)
	return
}
