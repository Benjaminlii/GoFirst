package main

import "fmt"

// 继承

type animal struct {
	name string
	age  int
}

func (animal animal) message() {
	fmt.Println("this is a animal!")
}

type wolf struct {
	// 这里就是那个匿名结构体，不需要名字，连_都不要有
	// 一旦进行了嵌套，那么可以使用匿名结构体的所有字段和方法
	// 不进行大小写的区分
	// 当字段名重复时，就近原则
	// 如果嵌套了两个或更多个结构体，那么在进行字段访问的时候
	// 需要知名匿名结构体的名字（使用.来一层一层访问）
	// 这说明Golang时支持多继承的
	animal
	hobby string
}

// 支持方法的重写
func (wolf wolf) message() {
	fmt.Println("this is a wolf!")
}

func main() {
	ani := animal{"A", 1}
	fmt.Println(ani)
	wolf := wolf{animal{"狼", 2}, "吃羊"}
	fmt.Println(wolf)
	ani.message()
	wolf.message()
}
