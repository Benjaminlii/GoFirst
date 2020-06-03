// package main
package main

// 测试Go中的时间和日期处理包

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(time.Now().Unix())
	// 测试time包中的时间格式化
	// 方式一：使用fmt.Sprintf("格式"，参数)来进行格式化填充
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	// 方式二：使用time.Time对象的Format方法进行格式化
	// 其中 2006，1，2，15，4，5这些数组都是固定代表某一个时间属性的，类比Java中的YYYY等等
	timeStr := now.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)
}
