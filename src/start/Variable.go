package main

import "fmt"

/**
 * @File:  Variable
 * @Description:
 * 练习Go语言中的变量的用法
 *
 * @Author: bytedance
 * @Date: 2020/4/14 15:55
 */

// 全局变量的声明
// 单个全局变量的声明
var globalNum1 = 100

// 多个全局变量的声明
var globalNum2, globalNum3 = 1, 2
var (
	globalNum4 = 4
	globalNum5 = 5
)

func main() {
	// 指定变量类型，如果不直接初始化，那么使用默认值
	var num1 int
	num1 += 100
	fmt.Println("num1 =", num1)

	// 根据值自动判定变量类型
	var num2 = 200
	fmt.Println("num2 =", num2)

	// 使用:=来省略var关键字
	num3 := 300
	fmt.Println("num3 =", num3)

	// 同时声明多个变量
	// 方法一
	var num4, num5, num6 int
	fmt.Println("num4 =", num4, ", num5 =", num5, ", num6 =", num6)
	// 方法二
	var num7, num8, num9 = 1, 2, 3
	fmt.Println("num7 =", num7, ", num8 =", num8, ", num9 =", num9)
	// 方法三
	num10, num11, num12 := 1, 2, 3
	fmt.Println("num10 =", num10, ", num11 =", num11, ", num12 =", num12)

	// 全局变量的声明测试
	fmt.Println("globalNum1 =", globalNum1,
		"\nglobalNum2 =", globalNum2,
		"\nglobalNum3 =", globalNum3,
		"\nglobalNum4 =", globalNum4,
		"\nglobalNum5 =", globalNum5)

	// 变量的变化范围
	var n1 int
	n1 = 10
	// 这里如果尝试进行类型转换就会报错
	// n1 = 10.1
	fmt.Println("n1 =", n1)
}
