// package main 是在学习Go语言过程中做的一些小练习
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// 学习基本数据类型

func main() {
	// int类型
	var numInt int = 100
	fmt.Println("numInt =", numInt)

	// byte类型相当于uint8
	var num1 byte = 225
	var num2 int8 = 127
	fmt.Println(num1, num2)

	// rune类型等价int32，刚好可以存储一个Unicode码
	var char rune = '李'
	fmt.Printf("%c\n", char)
	var num int = 124
	fmt.Printf("num的类型是%T, num占用的字节数是%d\n", num, unsafe.Sizeof(num))

	// 基本数据类型转string
	// 方式一：相当于是格式化输出，但是将要打印的字符串作为返回值返回
	str := fmt.Sprintf("%d.%d.%d %d:%d:%d", 2020, 4, 15, 18, 10, 00)
	fmt.Println(str)
	// 方式二：使用strconv包中的FormatXXXXX函数
	// FormatInt(要转换的值 int64，要转换的进制)
	fmt.Println(strconv.FormatInt(127, 2))
	// FormatFloat(要转换的值（float64），
	//             要转换的格式（byte类型，常用'f'十进制，'b'二进制，'E'指数），
	//             精度（小数点后的位数）,
	//             f的来源类型（32：float32、64：float64）) string
	fmt.Println(strconv.FormatFloat(127.123456789, 'f', 6, 32))
	// FormatBool(要转换的值（bool）)
	fmt.Println(strconv.FormatBool(true))

	// string转基本数据类型
	// 使用strconv包中的ParseXXXXXX函数来完成
	// ParseInt 参数：字符串，进制，整数的类型
	fmt.Println(strconv.ParseInt("123", 10, 64))
	// ParseFloat 参数：字符串，浮点数的类型
	fmt.Println(strconv.ParseFloat("123.456789", 32))
	// ParseBool 参数：字符串
	fmt.Println(strconv.ParseBool("true"))

}
