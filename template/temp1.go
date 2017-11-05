package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Email    string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.Email}}")
	p := Person{UserName: "Simon", Email: "123456@163.com"}
	t.Execute(os.Stdout, p)
}
