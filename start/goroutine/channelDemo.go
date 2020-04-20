package main

import "fmt"

// 管道测试

// 管道是引用类型，且必须有类型
// 管道需要先进行分配内存，即make后才能使用
// 管道的长度是不可变的
// 其使用价值在于一边放，一边取
var channel chan int = make(chan int, 10)

func produce() {

}

func consumption() {

}

func main() {
	fmt.Println(len(channel), cap(channel))
	channel <- 100
	channel <- 99
	channel <- 98
	channel <- 97
	channel <- 96
	channel <- 95
	channel <- 94
	fmt.Println(len(channel), cap(channel))
	close(channel)
	for val := range channel {
		fmt.Println(val)
	}
	fmt.Println(len(channel), cap(channel))
}
