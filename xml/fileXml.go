package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Result struct {
	Person []Person
}

type Person struct {
	Name      string `xml:",attr"`
	Age       int    `xml:",attr"`
	Career    string
	Interests Interests
}

type Interests struct {
	Interest []string
}

func main() {
	content, err := ioutil.ReadFile("studygolang.xml")
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)

	for _, p := range result.Person {
		log.Println(p.Name, p.Age, p.Career, p.Interests.Interest)
		/*for _, i := range p.Interests.Interest {
			log.Println(i)
		}*/
	}
}
