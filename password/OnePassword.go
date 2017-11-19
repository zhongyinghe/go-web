package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	h := sha256.New()
	io.WriteString(h, "12345678")
	fmt.Printf("%x\n", h.Sum(nil))

	h2 := sha1.New()
	io.WriteString(h2, "12345678")
	fmt.Printf("%x\n", h2.Sum(nil))

	h3 := md5.New()
	io.WriteString(h3, "12345678")
	fmt.Printf("%x\n", h3.Sum(nil))
}
