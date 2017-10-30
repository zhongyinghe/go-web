package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var t xml.Token
	var err error

	input := `<Person><FirstName>Zhong</FirstName><LastName>Yinghe</LastName></Person>`
	inputReader := strings.NewReader(input)
	fmt.Println(reflect.TypeOf(inputReader))
	decoder := xml.NewDecoder(inputReader)

	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token Name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
		}
	}
}
