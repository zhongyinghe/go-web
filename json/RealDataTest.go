package main

import (
	"encoding/json"
	"fmt"
	//"github.com/bitly/go-simplejson"
	"reflect"
)

var jsonStr = `
{"code":200,"message":"获取用户信息成功!","data":[{"created":"123456789","departname":"研发部门","uid":"1","username":"astaxieupdate"},{"created":"1507730274","departname":"研发部门","uid":"2","username":"zhongyinghe"},{"created":"1513047756","departname":"研发部门","uid":"5","username":"bird"},{"created":"1513241841","departname":"golang","uid":"6","username":"abc"},{"created":"1513243164","departname":"golang_kk","uid":"7","username":"abc_kk"},{"created":"1513243230","departname":"golang_kk","uid":"8","username":"abc_kk"},{"created":"1513243412","departname":"研发","uid":"9","username":"efg"},{"created":"1513243412","departname":"研发","uid":"10","username":"efg"},{"created":"1513313641","departname":"AI","uid":"11","username":"abc"},{"created":"1513313909","departname":"AI","uid":"12","username":"abc123"}]}
`

func main() {
	/*js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err.Error())
	}

	users, _ := js.Get("data").Array()
	fmt.Println(users)

	for i, _ := range users {
		//fmt.Println(i, reflect.TypeOf(u).Kind())
		user := js.Get("data").GetIndex(i)
		username := user.Get("username").MustString()
		fmt.Println(username)
	}*/

	b := []byte(jsonStr)
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}

	//datas := I2Slice(m["data"])
	datas := m["data"].([]interface{})
	for i, v := range datas {
		mv := v.(map[string]interface{})
		fmt.Println(i, mv["username"])
	}
}

func I2Slice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}

	l := v.Len()
	ret := make([]interface{}, l)

	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
