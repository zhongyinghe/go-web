package paes

import (
	"crypto/aes"
	"crypto/cipher"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	bResult := make([]byte, len(origData))
	cfb.XORKeyStream(bResult, origData)
	return bResult, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	bResult := make([]byte, len(crypted))
	cfbdec.XORKeyStream(bResult, crypted)
	return bResult, nil
}
