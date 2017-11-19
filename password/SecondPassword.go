package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("%x\n", h.Sum(nil))
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(pwmd5)

	salt1 := "@#$%"
	salt2 := "^&*()"

	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, pwmd5)
	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}
