package main

import "fmt"

// 管道测试

var channel = make(chan int, 10)

func main() {
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
