// package main
package main

import "fmt"

// 入门Golang面向对象

// Cat 类型为猫对象所属的类型
// 结构体是值类型，是值拷贝，但是其中的引用类型成员变量通过值拷贝还是只有一个变量
type Cat struct {
	// 名字
	Name string
	// 年龄
	Age int
	// 颜色
	Color string

	Slice []int
}

func main() {
	// 声明一个对象就会进行内存的分配，不需要make
	var aCat Cat
	aCat.Name = "小白"
	aCat.Age = 2
	aCat.Color = "white"
	aCat.Slice = append(aCat.Slice, 1)
	fmt.Println("aCat :", aCat)
	bCat := aCat
	fmt.Println("aCat.map = ", &(aCat.Slice[0]))
	fmt.Println("bCat.map = ", &(bCat.Slice[0]))

	// 创建结构体变量的四种方式
	// 方式一：直接声明
	var cCat Cat
	fmt.Println(cCat)
	// 方式二：带值初始化
	aCat = Cat{Name: "小白", Age: 1, Color: "white", Slice: []int{}}
	fmt.Println(aCat)

}
