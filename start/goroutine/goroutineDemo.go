// package main
package main

import (
	"fmt"
	"time"
)

// 协程

func doInGoroutine() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
		time.Sleep(time.Second)
	}
}

func doMaster() {
	for i := 0; i < 10; i++ {
		fmt.Println("world", i)
		time.Sleep(time.Second)
	}
}

func main() {
	// Golang中在go关键字之后调用函数或者方法即可在新的一个协程中执行这个方法
	go doInGoroutine()
	doMaster()
}
