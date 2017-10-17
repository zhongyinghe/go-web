package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const WEBROOT = "/home/GoWorkSpace/src/webform/"

func sign(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, terr := template.ParseFiles(WEBROOT + "sign.gtpl")

		if terr != nil {
			fmt.Println(terr)
			return
		}

		t.Execute(w, nil)
	} else {
		r.ParseForm()
		username := r.Form.Get("username")
		username = template.HTMLEscapeString(username)
		fmt.Println("hello")
		fmt.Fprintf(w, username)
	}
}

func suntpl(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))

	if err != nil {
		fmt.Println(err)
	}
}

func postOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, terr := template.ParseFiles(WEBROOT + "postOne.gtpl")
		if terr != nil {
			fmt.Println(terr)
			return
		}

		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")

		if token == "" {
			fmt.Fprintf(w, "token不存在")
			return
		}

		io.WriteString(w, "token为:"+token)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, terr := template.ParseFiles(WEBROOT + "upload.gtpl")
		if terr != nil {
			fmt.Println(terr)
			return
		}

		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile(WEBROOT+"upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/suntpl", suntpl)
	http.HandleFunc("/postOne", postOne)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
