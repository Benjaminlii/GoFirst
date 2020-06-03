// package main
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type a struct {
	name string
	age  int
}

func (a *a) print() {
	fmt.Printf("a.name = %v, a.age = %v\n", a.name, a.age)
}

func main() {
	var aObj a = a{name: "lt", age: 20}
	bObj := new(a)
	bObj.name = "lt1"
	bObj.age = 20
	aObj.print()
	bObj.print()

	fmt.Println("---------------")
	a, err := json.Marshal(nil)
	if err != nil {
		fmt.Println("a")
		return
	}
	fmt.Printf("[%v]\n", len(string(a)))

	timestampNow := time.Now().Unix()
	fmt.Println(timestampNow)
	fmt.Println(1591718399000)
	if timestampNow < 1591718399000 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}

	tm := time.Unix(1591718399, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))

}
