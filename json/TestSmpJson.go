package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	js, err := simplejson.NewJson([]byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	}
}`))

	if err != nil {
		fmt.Println(err)
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(ms)
}
