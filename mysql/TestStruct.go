package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserInfo struct {
	Uid        int    `table:"userinfo" column:"uid" json:"-"`
	UserName   string `column:"username"`
	Departname string `column:"departname"`
	Created    int64  `column:"created" json:"-"`
}

func main() {
	u := UserInfo{
		Uid:        11,
		UserName:   "siscall",
		Departname: "AI",
		Created:    time.Now().Unix(),
	}

	/*b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/

	b := ResponseJSON(200, "Struct转化为json数据", u)

	fmt.Println(string(b))
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
