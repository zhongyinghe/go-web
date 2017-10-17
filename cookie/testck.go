package main

import (
	"fmt"
	"log"
	"net/http"
	//"reflect"
	"time"
)

func setWebCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "zyh", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func getWebCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie.Name)
	//fmt.Println(reflect.TypeOf(username))
	fmt.Fprint(w, cookie.Value)
}

func main() {
	http.HandleFunc("/setWebCookie", setWebCookie)
	http.HandleFunc("/getWebCookie", getWebCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
