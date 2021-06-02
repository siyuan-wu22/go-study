package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "1"
	fmt.Println(reflect.TypeOf(s))
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(reflect.TypeOf(i))
}
