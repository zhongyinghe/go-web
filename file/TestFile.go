package main

import (
	"fmt"
	"os"
)

func main() {
	readme := "readme.txt"
	file, err := os.Create(readme)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < 10; i++ {
		file.WriteString("Just a test!\r\n")
		file.Write([]byte("Just a test!\r\n"))
	}
}
