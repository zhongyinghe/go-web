package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get err.")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//将HTML标签全转换成小写
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = reg.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	reg, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = reg.ReplaceAllString(src, "")

	//去除SCRIPT
	reg, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = reg.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	reg, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = reg.ReplaceAllString(src, "\n")

	//去除连续的换行符
	reg, _ = regexp.Compile("\\s{2,}")
	src = reg.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
