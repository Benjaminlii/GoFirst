// package main
package main

import "fmt"

// 学习复杂数据类型

func main() {

	// 指针
	// 声明一个指针
	var num int = 100
	var p *int = &num
	// 简介访问这个指针的内容
	fmt.Println(*p)
	*p = 200
	fmt.Println(*p)

}
