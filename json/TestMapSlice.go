package main

import (
	"fmt"
)

func main() {
	ms := make([]map[string]interface{}, 0)

	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["uid"] = i
		m["username"] = "ss" + fmt.Sprintf("%d", i)
		ms = append(ms, m)
	}

	fmt.Println(ms)

	for i, v := range ms {
		fmt.Println(i, v["username"])
	}
}
