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
	rows, err := db.Query("SELECT * FROM userinfo WHERE uid=?", 1)
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

	records := make([]map[string]interface{}, 0)

	for rows.Next() {

		err := rows.Scan(scanArgs...)
		checkErr(err)
		record := make(map[string]interface{})
		fmt.Println(values)
		/*for i, col := range values {
			if col != nil {
				//record[columns[i]] = string(col.([]byte))
				fmt.Println(i, string(col.([]byte)))
			}
		}*/
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
				fmt.Println(col)
			} else {
				v = val
			}
			record[col] = v
		}

		records = append(records, record)
		fmt.Println(record)
	}

	fmt.Println(records)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
