// package main
package main

import "fmt"

// 用于练习流程控制中的分支结构

func main() {
	// if语句
	// 输入a，b中的一个，返回不同的语句
	var ch byte
	_, _ = fmt.Scanf("%c", &ch)
	if ch == 'a' {
		fmt.Println("you input a！")
	} else if ch == 'b' {
		fmt.Println("you input b！")
	} else {
		fmt.Println("Undefined！")
	}

	// switch语句
	// 输入
	var char byte
	_, _ = fmt.Scanf("%c", &char)
	switch char {
	case 'a', 'b':
		fmt.Println("you input a or b！")
	case 'c', 'd':
		fmt.Println("you input c or d！")
	default:
		fmt.Println("Undefined！")
	}

}
