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

// 函数的闭包
// 累加器
// 这里返回的是一个函数，其具体的定义是内部的匿名函数
// 但在这个匿名函数内引用了外部的变量count，就形成了闭包
// 也就是说，匿名函数可以引用其外部的变量，并和这些外部变量共同组成一个闭包
// 可以理解为整个闭包是一个对象，这个函数是其中的方法，变量则担任成员变量的指责
// 这里的addCount函数只进行了一次，所以count变量只会被初始化一次，后边被内部的方法修改
// 生成新的方法也不会影响原先的方法的闭包
func addCount() func(int) int {
	var count int = 0
	return func(num int) int {
		count += num
		return count
	}
}

// defer机制
// 在函数中前面标有defer关键字的语句会被压入defer栈
// 在函数执行完之后，会从defer栈中一条一条的取出语句来执行
// 因为是栈，所以要注意defer语句执行的顺序
// 这里使用栈的原因，个人觉得是考虑到资源的释放需要从下向上进行释放
// 需要注意的是，defer语句中的变量，是一个快照，不会受后面语句的影响而输出不同的结果
func deferDemo() int {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("do 1")
	return 1
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
	// 个人认为匿名函数一般在函数传参处会比较灵活，可以类比Java中的静态内部类传入方法
	// 求两个数的和的匿名函数
	num := func(num1 int, num2 int) int {
		return num1 + num2
	}(1, 2)
	fmt.Println(num)

	// 在函数内通过匿名函数定义一个函数，这样就将作用域限制为外层的函数
	get1 := func() int {
		return 1
	}
	fmt.Println(get1())
	fun1 := addCount()
	fmt.Println(fun1(1))
	fmt.Println(fun1(1))
	fun2 := addCount()
	fun2(1)
	fmt.Println(fun1(1))
	fmt.Println(fun1(1))
	fmt.Println(fun1(1))

	// defer
	fmt.Println(deferDemo())
}
