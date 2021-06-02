package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string = "this is a program"
	var i int = 1
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
}
