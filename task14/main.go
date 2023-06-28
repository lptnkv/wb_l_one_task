package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	ch := make(chan int)

	x = 5
	getTypeWithSwitch(x)
	getTypeWithFmt(x)
	getTypeWithReflect(x)

	x = false
	getTypeWithSwitch(x)
	getTypeWithFmt(x)
	getTypeWithReflect(x)

	x = "Hello, world!"
	getTypeWithSwitch(x)
	getTypeWithFmt(x)
	getTypeWithReflect(x)

	x = ch
	getTypeWithSwitch(x)
	getTypeWithFmt(x)
	getTypeWithReflect(x)
}

func getTypeWithSwitch(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("This is int")
	case bool:
		fmt.Println("This is bool")
	case chan int:
		fmt.Println("This is chan int")
	case string:
		fmt.Println("This is string")
	default:
		fmt.Println("Don't know this type")
	}
}

func getTypeWithFmt(a interface{}) {
	fmt.Printf("Type of %v is %T\n", a, a)
}

func getTypeWithReflect(a interface{}) {
	fmt.Println(reflect.TypeOf(a))
}
