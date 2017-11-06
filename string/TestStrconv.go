package main

import (
	"fmt"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')

	fmt.Println(string(str))
	fmt.Println("-----------append end-------------")

	a := strconv.FormatBool(false)
	//b := strconv.FormatFloat(123.23, "g", 12, 64)//错误写法"g",而是'g'
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
	fmt.Println("-----------format end-------------")

	f, err := strconv.ParseBool("false")
	checkError(err)

	g, err := strconv.ParseFloat("123.23", 64)
	checkError(err)

	h, err := strconv.ParseInt("12345", 10, 64)
	checkError(err)

	i, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(f, g, h, i)
	fmt.Println("-----------parse end-------------")
}
