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

	key := "go_list"
	for i := 1; i < 10; i++ {
		_, err = c.Do("rpush", key, i)
		if err != nil {
			fmt.Println("redis set failed:", err)
			return
		}
	}

	_, err = c.Do("lrem", key, 1, 5)
	if err != nil {
		fmt.Println("redis delete failed:", err)
		return
	}

	_, err = c.Do("lpush", key, 5)
}
