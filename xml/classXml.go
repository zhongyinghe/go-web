package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type OneClass struct {
	XMLName  xml.Name  `xml:"Class"`
	Version  string    `xml:"Version,attr"`
	Students []Student `xml:"Student"`
}

type Student struct {
	StuNo string `xml:"StuNo"`
	Name  string `xml:"Name"`
	Sex   int    `xml:"Sex"`
	Age   int    `xml:"Age"`
}

func main() {
	c := &OneClass{Version: "1"}
	c.Students = append(c.Students, Student{"200712300160", "zhongyinghe", 1, 32})
	c.Students = append(c.Students, Student{"200712300161", "lzx", 1, 30})
	c.Students = append(c.Students, Student{"200712300162", "zw", 2, 29})
	c.Students = append(c.Students, Student{"200712300163", "xw", 1, 30})
	c.Students = append(c.Students, Student{"200712300164", "zzm", 1, 31})
	c.Students = append(c.Students, Student{"200712300165", "xk", 1, 31})

	output, err := xml.MarshalIndent(c, "  ", "    ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	headerBytes := []byte(xml.Header)
	output = append(headerBytes, output...)
	fileName := "/home/GoWorkSpace/src/xml/oneclass.xml"
	ioutil.WriteFile(fileName, output, os.ModeAppend)
	os.Chmod(fileName, 0775)
	fmt.Println("ok")
	//os.Stdout.Write([]byte(xml.Header))

	//os.Stdout.Write(output)
	//fmt.Println(string(output))
}
