package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "append.txt"
	_, errf := os.Stat(fileName)
	if os.IsNotExist(errf) {
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Close()
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		file.Write([]byte("Let's Go \r\n"))
		file.Sync()
	}
}
