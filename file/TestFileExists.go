package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("append.txt")
	fmt.Println(err)
	fmt.Println(os.IsNotExist(err))
}
