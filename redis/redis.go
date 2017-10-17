package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
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
	_, err = c.Do("SET", "mykey", "wangwu")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mykey"))

	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v\n", username)
	}

	//设置过期键
	_, err = c.Do("SET", "exkey", "lisi", "EX", "5")
	if err != nil {
		fmt.Println("set exkey failed:", err)
		return
	}

	lisi, err := redis.String(c.Do("GET", "exkey"))
	if err != nil {
		fmt.Println("get exkey failed:", err)
		return
	} else {
		fmt.Printf("exkey is : %v\n", lisi)
	}

	time.Sleep(time.Second * 8)

	lisi, err = redis.String(c.Do("GET", "exkey"))
	if err != nil {
		fmt.Println("reget exkey failed:", err)
		return
	} else {
		fmt.Printf("exkey2 is : %v\n", lisi)
	}
}
