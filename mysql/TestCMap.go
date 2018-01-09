package main

import (
	"database/sql"
	//"encoding/json"
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
	rows, err := db.Query("SELECT * FROM userinfo")
	defer rows.Close()
	checkErr(err)
	doRows(rows)

}

func doRows(rows *sql.Rows) {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	records := make([]map[string]string, 0)

	for rows.Next() {

		err := rows.Scan(scanArgs...)
		checkErr(err)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
		//fmt.Println(record)
	}

	fmt.Println(records)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
