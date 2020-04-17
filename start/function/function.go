// pkg main
package main

import "fmt"

//函数相关知识练习

// func 函数名 (行参列表)(返回值类型列表){
//     /* 代码块 */
//     return 返回值列表
// }

// 单个返回值的函数
func getSum(num1 int, num2 int) (sum int) {
	return num1 + num2
}

// 多个返回值的函数
func getSumAndSub(num1 int, num2 int) (sum int, sub int) {
	return num1 + num2, num1 - num2
}

// 函数的递归调用
// 斐波那契数列的第n项
func getFib(index int) int {
	if index == 1 || index == 2 {
		return 1
	}
	return getFib(index-1) + getFib(index-2)
}

// 函数的返回值列表
func getSomeNum() (int, int, int, int) {
	return 1, 2, 3, 4
}

// 函数作为形参进行传入调用
func function1(aFunc func(int) int, num int) int {
	return aFunc(num)
}

// 可变参数
// 求一到多个参数的和
func getSumFromNums(num1 int, arg ...int) (sum int) {
	sum = num1
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

// init函数
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println(getSum(1, 1))

	sum, sub := getSumAndSub(100, 50)
	fmt.Println("sum =", sum, ",sub =", sub)

	fmt.Println(getFib(6))

	fmt.Println(getSomeNum())
	num1, num2, num3, num4 := getSomeNum()
	fmt.Println(num1, num2, num3, num4)

	fmt.Println(function1(getFib, 6))

	fmt.Println(getSumFromNums(1, 2, 3, 4, 5))

	// 匿名函数

}
