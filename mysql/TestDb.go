package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var db *sql.DB

var wg sync.WaitGroup

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	//checkErr(err)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.Ping()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()
	wg.Add(2)
	go showData(1)
	go showData(2)
	wg.Wait()
}

func showData(i int) {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created int
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(i, uid)
		fmt.Println(i, username)
		fmt.Println(i, department)
		fmt.Println(i, created)
	}
	wg.Done()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
