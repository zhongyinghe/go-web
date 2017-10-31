package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	//加入XML头
	headerBytes := []byte(xml.Header)
	//拼接XML头和实际XML内容
	xmlOutPutData := append(headerBytes, output...)
	//写入文件
	fileName := "/home/GoWorkSpace/src/xml/studygolang_test.xml"

	ioutil.WriteFile(fileName, xmlOutPutData, os.ModeAppend)
	os.Chmod(fileName, 0666)
	fmt.Println("ok")
	//os.Stdout.Write([]byte(xml.Header))

	//os.Stdout.Write(output)
	//fmt.Println(string(output))
}
