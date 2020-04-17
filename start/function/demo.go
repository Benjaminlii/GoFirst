package main

import "fmt"

// 测试Go内置函数

func main() {
	num := new(int)
	fmt.Println("num =", num)
	fmt.Println("*num =", *num)
	fmt.Println("&num =", &num)
	fmt.Printf("num的类型 = %T\n", num)
}
