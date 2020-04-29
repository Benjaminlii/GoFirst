// package main
package main

import "fmt"

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
}
