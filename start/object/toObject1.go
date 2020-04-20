// package main
package main

import (
	"fmt"
)

// 入门Golang面向对象

// Cat 类型为猫对象所属的类型
// 结构体是值类型，是值拷贝
// 但其中的引用类型成员变量通过值拷贝的结果是连个不同的指针指向同一个空间
// 字段后面加上反引号引上的tag，可以在序列化时更改字段名，方便其他语言的开发
type Cat struct {
	// 名字
	Name string `json:"name"`
	// 年龄
	Age int `json:"age"`
	// 颜色
	Color string `json:"color"`

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
	fmt.Printf("aCat{&(Name):%p,&(age):%p,&(Color):%p}\n", &aCat.Name, &aCat.Age, &aCat.Color)
	fmt.Printf("bCat{&(Name):%p,&(age):%p,&(Color):%p}\n", &bCat.Name, &bCat.Age, &bCat.Color)
	fmt.Printf("aCat.Slice = %p\n", &(aCat.Slice))
	fmt.Printf("bCat.Slice = %p\n", &(bCat.Slice))

	// 创建结构体变量的四种方式
	// 方式一：直接声明
	var cCat Cat
	fmt.Println(cCat)
	// 方式二：带值初始化
	aCat = Cat{Name: "小白", Age: 1, Color: "white"}
	fmt.Println(aCat)
	// 方式三：使用new创建
	// 这样会得到一个指针
	// Golang做了优化，支持结构体指针.字段名的访问方式，不需要再去加间接访问符
	pCat := new(Cat)
	fmt.Printf("%p\n", pCat)

	str1 := "abc"
	str2 := str1
	fmt.Printf("&(str1)=%p, &(str2)=%p\n", &(str1), &(str2))
	fmt.Println(str1[0])

	fmt.Printf("aCat{&(Name):%p,&(age):%p,&(Color):%p}\n", &aCat.Name, &aCat.Age, &aCat.Color)
	fmt.Println(aCat)
	aCat.Name = "xiaobai"
	fmt.Printf("aCat{&(Name):%p,&(age):%p,&(Color):%p}\n", &aCat.Name, &aCat.Age, &aCat.Color)
	fmt.Println(aCat)

	// 结构体之间转换时必须要有完全相同的字段才可以

}
