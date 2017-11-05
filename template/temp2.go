package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
			{{end}}
		`)

	f1 := Friend{Fname: "jack"}
	f2 := Friend{Fname: "Mary"}
	p := Person{
		UserName: "simon",
		Emails:   []string{"123456@163.com", "963258@qq.com"},
		Friends:  []*Friend{&f1, &f2},
	}

	t.Execute(os.Stdout, p)
}
