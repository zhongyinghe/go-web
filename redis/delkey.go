package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()
	c.Do("AUTH", "123456") //授权

	//设置名称
	_, err = c.Do("SET", "mydelkey", "wangwu")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mydelkey"))

	if err != nil {
		fmt.Println("redis get mydelkey failed:", err)
	} else {
		fmt.Printf("Get mydelkey: %v\n", username)
	}

	_, err = c.Do("DEL", "mydelkey")
	if err != nil {
		fmt.Println("delete mydelkey failed:", err)
		return
	}

	username, err = redis.String(c.Do("GET", "mydelkey"))

	if err != nil {
		fmt.Println("redis get mydelkey failed:", err)
	} else {
		fmt.Printf("Get mydelkey: %v\n", username)
	}
}
