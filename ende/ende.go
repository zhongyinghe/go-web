package ende

import (
	"encoding/base64"
	"ende/paes"
	"ende/pdes"
	"ende/prsa"
	"fmt"
)

//des编码
func DesEncode(str, key string) string {
	if len(key) != 8 {
		fmt.Println("key size must be 8")
		return ""
	}

	bResult, err := pdes.DesEncrypt([]byte(str), []byte(key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(bResult)
}

//des解码
func DesDecode(str, key string) string {
	if len(key) != 8 {
		fmt.Println("key size must be 8")
		return ""
	}

	//base64解码
	buf, _ := base64.StdEncoding.DecodeString(str)

	bResult, err := pdes.DesDecrypt(buf, []byte(key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bResult)
}

//3des加密
func Des3Encode(str, key string) string {
	if len(key) != 24 {
		fmt.Println("key size must be 24")
		return ""
	}

	bResult, err := pdes.TripleDesEncrypt([]byte(str), []byte(key))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(bResult)
}

//3des解密
func Des3Decode(str, key string) string {
	if len(key) != 24 {
		fmt.Println("key size must be 24")
		return ""
	}

	buf, _ := base64.StdEncoding.DecodeString(str)
	bResult, err := pdes.TripleDesDecrypt(buf, []byte(key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bResult)
}

//aes加盟
func AesEncode(str, key string) string {
	if len(key) != 24 {
		fmt.Println("key size must be 24")
		return ""
	}

	bResult, err := paes.AesEncrypt([]byte(str), []byte(key))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bResult)
}

//aes解密
func AesDecode(str, key string) string {
	if len(key) != 24 {
		fmt.Println("key size must be 24")
		return ""
	}

	buf, _ := base64.StdEncoding.DecodeString(str)
	bResult, err := paes.AesDecrypt(buf, []byte(key))

	if err != nil {
		return ""
	}

	return string(bResult)
}

//rsa加密
func RsaEncode(str string, publicKey []byte) string {
	bResult, err := prsa.RsaEncrypt([]byte(str), publicKey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bResult)
}

//rsa解密
func RsaDecode(str string, privateKey []byte) string {
	buf, _ := base64.StdEncoding.DecodeString(str)
	bResult, err := prsa.RsaDecrypt(buf, privateKey)
	if err != nil {
		return ""
	}
	return string(bResult)
}
