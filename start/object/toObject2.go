package main

import "fmt"

// 方法

// Dog 狗
type Dog struct {
	// 名字
	Name string `json:"name"`
	// 年龄
	Age int `json:"age"`
	// 颜色
	Color string `json:"color"`

	Slice []int
}

// ADogWoof 狗叫
func (dog Dog) ADogWoof() {
	fmt.Printf("%s:Woof!\n", dog.Name)
}

// OneDogWoof 狗叫
func (Dog) OneDogWoof() {
	fmt.Println("Woof!")
}

// 如果需要在方法内对对象进行修改，那么需要传入指针
// 对于内部指针的使用，不需要加间接访问符了，编译器做了优化

// HaveABirthday 年龄增长一岁
func (dog *Dog) HaveABirthday() {
	dog.Age++
}

// 如果定义了一个类的String方法
// 那么在fmt中输出这个对象时就会自动调用这个对象的这个方法替换变量进行输出
func (dog *Dog) String() string {
	return fmt.Sprintf("dog{Name:%v, Age:%v, Color:%v}\n", dog.Name, dog.Age, dog.Color)
}

func main() {
	dog := Dog{Name: "旺财", Age: 3, Color: "Yellow"}
	dog.ADogWoof()
	fmt.Println(dog)
	dog.HaveABirthday()
	fmt.Println(dog)

}
