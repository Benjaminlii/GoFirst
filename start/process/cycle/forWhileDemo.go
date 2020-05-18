// Package main
package main

import (
	"fmt"
)

//循环结构

func main() {
	// for循环
	// 打印10句hello world
	for i := 0; i < 10; i++ {
		fmt.Println(i, "hello world!")
	}

	// go语言中的while循环也是使用for关键字实现的
	var a = 0
	for a < 100 {
		fmt.Println(a)
		a++
	}

	// 接下来是go语言中的for-range循环，类似Java中的for-each循环
	var str = "hello world"
	for index, ch := range str {
		fmt.Printf("str[%d] = %c\n", index, ch)
	}
	fmt.Println("---------------------------------------")
	// for each循环的map遍历
	var mapA = map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}
	for key, val := range mapA {
		fmt.Printf("key = %v, val = %v\n", key, val)
	}

}
