// package main
package main

import "fmt"

// 接口

// Golang中的接口是一组方法，不允许有变量存在
// Golang中不需要显式的进行接口的实现
// 只需要实现接口中的所有方法，就算是实现了接口
// 没有实现接口的对象在进行传参时会报错
type myInterface interface {
	doSomething()
}

// 接口也是可以继承的
// 继承的方式和结构体相同
// 只有实现了所有继承的方法，才算实现了子接口
type interfaceA interface {
	A()
}
type interfaceB interface {
	B()
}
type interfaceAAndB interface {
	interfaceA
	interfaceB
}

type one struct {
}

func (one one) doSomething() {
	fmt.Println("one do something!")
}

type two struct {
}

func (two two) doSomething() {
	fmt.Println("two do something!")
}

// 传参使用接口
func doSomething(myin myInterface) {
	myin.doSomething()
}

func main() {
	oneObj := one{}
	twoObj := two{}
	doSomething(oneObj)
	doSomething(twoObj)

	// 使用接口引用实现接口的指针
	// 接口变量必须使用var关键字创建，因为不能自动判断其类型
	var myinf myInterface = oneObj
	myinf.doSomething()

	var inter interface{} = oneObj
	twoObj1, isOk := inter.(two)
	if !isOk {
		fmt.Println("err!")
	}
	fmt.Println(twoObj1)

	switch inter.(type) {
	case one:
		fmt.Println("one")
	case two:
		fmt.Println("two")
	}

}
