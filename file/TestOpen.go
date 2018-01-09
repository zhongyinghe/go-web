package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "readme.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString("Just a test!\r\n")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
