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
	_, err = c.Do("SET", "zhangkey", "zhangsan")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	is_key_exit, err := redis.Bool(c.Do("EXISTS", "zhangkey"))

	if err != nil {
		fmt.Println("error is :", err)
	} else {
		fmt.Printf("exists or not: %v\n", is_key_exit)
	}
}
