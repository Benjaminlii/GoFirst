// package main
package main

import "fmt"

// 先了解一下Golang中的数组

func main() {
	// 一维数组
	var array [10]int
	fmt.Println(array)
	array1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array1)
	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array2)

	// 二维数组
	arrays := [2][3]int{}
	fmt.Println(arrays)
	arrays1 := [][5]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}}
	fmt.Println(arrays1)
}
