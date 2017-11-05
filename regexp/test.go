package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println(os.Args[1])

	var a interface{}

	a = 123
	s := fmt.Sprintf("%v", a)
	fmt.Println(s)

	fmt.Println(reflect.TypeOf(s))
}
