package main

import (
	"database/sql"
	"encoding/json"
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
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	records := make([]map[string]string, 0, 40)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		err = rows.Scan(scanArgs...)
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

	//fmt.Println(records)
	/*b, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err.Error())
	}
	*/
	b := ResponseJSON(200, "转化为json格式数据", records)
	fmt.Println(string(b))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnJSON(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return b
}

func ResponseJSON(code int, message string, i interface{}) []byte {
	rd := ReturnData{
		Code:    code,
		Message: message,
		Data:    i,
	}
	return ReturnJSON(rd)
}

//返回数据的格式结构
type ReturnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
