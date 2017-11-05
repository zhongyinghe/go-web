package main

import (
	"fmt"
	"regmatch"
)

func main() {
	b := regmatch.IsIP("192.168.1.1")
	fmt.Println(b)

	chinese := "我是中国人"
	ch := regmatch.IsChinese(chinese)
	fmt.Println(ch)

	email := "xxxxx@qq.com"
	isE := regmatch.IsEmail(email)
	fmt.Println(isE)

	phone := "xxxxxxxxxx"
	isPhone := regmatch.IsPhone(phone)
	fmt.Println(isPhone)

	usercard := "123456789369257"
	isCard := regmatch.IsUserCard(usercard)
	fmt.Println(isCard)

	isrange := regmatch.InRange("orange", []string{"bannaner", "orange", "apple"})
	fmt.Println(isrange)

	isRangeInRange := regmatch.RangeInRange([]string{"abc", "abf"}, []string{"abc", "abe", "abd"})
	fmt.Println(isRangeInRange)

	num := 12369
	isNumber := regmatch.IsNumber(num)
	fmt.Println(isNumber)
}
