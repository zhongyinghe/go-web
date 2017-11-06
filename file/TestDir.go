package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)

	//创建目录
	os.Mkdir("zyh/abc", 0777)

	//删除目录
	os.MkdirAll("bbs/bbq", 0777)
	err := os.Remove("zyh")
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("zyh")
	os.RemoveAll("bbs")
}
