package main

import "fmt"

// Golang中的map

func main() {
	// golang中的map
	// 创建方式一：通过make去创建
	// 使用[keyType]valueType的形式表示类型
	// 存储的方式类似python
	// make中可以直接声明容量，后续可以动态变化
	var myMap map[string]int64 = make(map[string]int64, 1)
	myMap["Benjamin"] = 4173130
	myMap["a"] = 12341234
	myMap["b"] = 12341234
	myMap["c"] = 12341234
	fmt.Println(myMap)
	fmt.Println(myMap["Benjamin"])
	fmt.Println(4173130)

	// 创建方式二：直接赋值
	// 这里面的每一个key-value都需要一个逗号
	var myMap1 = map[string]int{
		"Benjamin": 1,
		"a":        1,
	}
	fmt.Println(myMap1)

	// map中信息的删除使用delete内置函数
	delete(myMap, "c")
	fmt.Println(myMap)

	// map的情况可以遍历key一个一个进行删除
	// 也可以直接make一个新的map，让gc去清除这个map

	// 对于map的查找，还会返回一个bool，用来表示是否存在这个key
	num, bool := myMap["Ben"]
	if bool {
		fmt.Println(num)
	}

	// 关于map的遍历
	// map只能使用for-range循环遍历，其中下标项代表的是key
	fmt.Println("myMap:")
	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}

	// map和切片
	// 可以理解为动态增长的map数组
	// 这里的mapArray对应Java中的ArrayList<HashMap<String, String>>
	mapArray := make([]map[string]string, 1)
	mapArray[0] = map[string]string{}
	mapArray = append(mapArray, map[string]string{})
	mapArray[0]["name"] = "Benjamin"
	mapArray[0]["age"] = "21"
	mapArray[1]["name"] = "Ben"
	mapArray[1]["age"] = "20"
	// 直接append进去会更合理
	myMap2 := map[string]string{
		"name": "lt",
		"age":  "21",
	}
	mapArray = append(mapArray, myMap2)
	fmt.Println(mapArray)
}
