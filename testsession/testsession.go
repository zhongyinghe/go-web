package main

import (
	"html/template"
	"log"
	"net/http"
	"session"
	_ "session/memory"
)

const WEBROOT = "/home/GoWorkSpace/src/testsession/"

var globalSessions *session.SessionManager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles(WEBROOT + "login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.FormValue("username"))
		http.Redirect(w, r, "/login", 302)
	}
}

func main() {
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
