package main

import (
	"database/sql"
	//"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo WHERE uid=?", 1)
	defer rows.Close()
	if err != nil {
		checkErr(err)
		return
	}

	res := FetchOneFromRows(rows)
	fmt.Println(res)

}

func FetchOneFromRows(rows *sql.Rows) map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	if !rows.Next() {
		checkErr(errors.New("获取数据失败!"))
		return nil
	}

	r := make(map[string]interface{})
	rows.Scan(scanArgs...)
	for i, col := range columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		r[col] = v
	}
	return r
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
