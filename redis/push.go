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

	_, err = c.Do("lpush", "runkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	_, err = c.Do("lpush", "runkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	_, err = c.Do("lpush", "runkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	values, _ := redis.Values(c.Do("lrange", "runkey", "0", "-1"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

}
