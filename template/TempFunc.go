package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}

	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return (substrs[0] + "at" + substrs[1])
}

func main() {
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.|emailDeal}}
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
