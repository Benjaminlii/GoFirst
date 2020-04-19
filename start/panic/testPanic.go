// package main
package main

import (
	"errors"
	"fmt"
)

// 测试Golang中的异常panic

// 运行时异常
func testPanic1() int {
	// 这个defer将一个异常处理语句压入defer栈
	// 发生异常后，当前函数会结束，并进入defer语句的执行阶段
	// 通过recover()内置函数可以捕获到之前退出的函数的异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err =", err)
		}
	}()
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	fmt.Println("num3 =", num3)
	return num3
}

// 可检查异常
func testPanic2(num int) (int, error) {
	// errors.New(str string)函数可以返回一个自定义的错误
	return num, errors.New("my error")
}

func main() {
	// 运行时异常
	testPanic1()
	fmt.Println("Run after panic!")

	// 可检查异常
	ans, err := testPanic2(1)
	fmt.Println("ans =", ans)
	if err != nil {
		// panic终止程序的进行，报错
		panic(err)
		// 当然也可以自行处理
		// fmt.Println(err)
	}
}
