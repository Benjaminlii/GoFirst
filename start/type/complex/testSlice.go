package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Golang中的切片

func main() {
	// 切片类型的解释
	// 切片类型实际上是一个数据结构
	// 其中存储了切片引用的第一个元素的地址，当前切片的长度（len），当前切片的容量（cap）

	// 使用切片的三种方式

	// 方式一：引用一段数组
	// 首先定义一个数组
	array := [...]int{1, 2, 3, 4, 5}
	// 切片slice的声明，这里这个切片相当于数组中一段元素的引用
	// 引用的下标为左闭右开
	slice := array[1:3]
	fmt.Println(slice)

	// 方式二：通过make来创建空的切片，相当于指向一片空的空间
	var slice1 []int = make([]int, 4, 10)
	slice1[0] = 0
	slice1[1] = 1
	slice1[2] = 2
	//slice1[4] = 4
	fmt.Println(slice1)

	// 方式三：直接指定具体数组
	// 底层类似make的形式，数组还是没有引用，只能通过slice去访问
	var slice2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	// 切片的扩容
	// append(slice []Type, elem1, elem2....) []Type
	// append函数可以在切片的后面添加元素
	// 如果是在容量范围内，那就直接添加在原基本数组后面，并更新其长度
	// 如果超出了容量范围，那么重新分配基本数组，并复制值
	// 但无论进不进行数组的重新申请，都会返回新的slice结构体
	// 需要通过返回值来接收追加后的元素
	// 扩容后的大小为min(2*oldCap, newCap+1)
	// 如果是引用数组的创建方式，那么会覆盖数组后面的元素
	var slice3 []int = []int{1, 2, 3, 4, 5}
	fmt.Println("slice3 = ", slice3, "len = ", len(slice3), "cap = ", cap(slice3))
	slice3 = append(slice3, 6)
	slice4 := append(slice3, 7)
	fmt.Println("slice3 = ", slice3, "len = ", len(slice3), "cap = ", cap(slice3), "&slice3=", &slice3[0])
	fmt.Println("slice4 = ", slice4, "len = ", len(slice4), "cap = ", cap(slice4), "&slice4=", &slice4[0])

	// 也可以直接追加一个切片或者数组
	// 通过...这个语法糖将slice打散进行值的传递
	slice3 = append(slice3, slice4...)
	fmt.Println("slice3 = ", slice3, "len = ", len(slice3), "cap = ", cap(slice3), "&slice3=", &slice3[0])

	// 字符串也可以进行切片的处理，因为其底层实际上的byte数组
	// 但是这个切片不可以更改，因为得到的还是一个字符串，并且他没有cap
	str := "my name is Benjamin"
	strSlice := str[11:]
	fmt.Printf("strSlice = %s, type = %T", strSlice, strSlice)

	// 如果仅使用append的方式去添加数据，那么不需要进行make
	var slice5 []int
	slice5 = append(slice5, 1)
	slice5 = append(slice5, 2)
	fmt.Println("slice5 = ", slice5)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	// json转Slice
	slice100 := make([]string, 10)
	jsonString := "[\"123\",\"456\",\"789\"]"
	test(jsonString, &slice100)
	fmt.Println(slice100)
	fmt.Println(slice100[0])
	fmt.Printf("%T\n", slice100[0])
}

func test(jsonString string, slice interface{}) error {

	// 转换为JSON格式
	err := json.Unmarshal([]byte(jsonString), &slice)
	if err != nil {
		return errors.New("")
	}
	return nil
}
