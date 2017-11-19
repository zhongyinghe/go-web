package main

import (
	"ende"
	"fmt"
	"io/ioutil"
)

func main() {
	publicKey, err := ioutil.ReadFile("./public.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(publicKey))

	privateKey, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(privateKey))

	str := "hello, RSA! Are you ok?"
	result := ende.RsaEncode(str, publicKey)
	fmt.Println(result)

	oriStr := ende.RsaDecode(result, privateKey)
	fmt.Println(oriStr)
}
