package main

import (
	"ende"
	"fmt"
)

func main() {
	str := "i am test yy gg ddd"
	key := "##@()&*^"
	strCode := ende.DesEncode(str, key)
	fmt.Println(strCode)

	result := ende.DesDecode(strCode, key)
	fmt.Println(result)

	str2 := "this is a tester"
	key2 := "sfe023f_sefiel#fi32lf3e!"
	strCode2 := ende.Des3Encode(str2, key2)
	fmt.Println(strCode2)

	result2 := ende.Des3Decode(strCode2, key2)
	fmt.Println(result2)

	fmt.Println("----------------------")

	key3 := "sfe023f_sefiel#fi32lf3e!"
	fmt.Println(len(key3))
	str3 := "this is a tester"
	strCode3 := ende.AesEncode(str3, key3)
	fmt.Println(strCode3)

	result3 := ende.AesDecode(strCode3, key3)
	fmt.Println(result3)
}
