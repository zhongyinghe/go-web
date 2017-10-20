package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()
	c.Do("AUTH", "123456") //授权

	key := "htable_user"

	n, err := c.Do("HSET", key, "name", "zhongyinghe")
	if err != nil {
		fmt.Println(err)
		return
	}

	if n == int64(1) {
		fmt.Println("Success")
	}

	name, err := redis.String(c.Do("HGet", key, "name"))
	if err != nil {
		fmt.Println("hget error: ", err)
		return
	}
	fmt.Println(name)

	n, err = c.Do("HMSET", key, "age", "32", "sex", "man")
	if err != nil {
		fmt.Println(err)
		return
	}

	r2, err := redis.Strings(c.Do("HMGET", key, "name", "age", "sex"))
	if err != nil {
		fmt.Println("MGet error: ", err)
		return
	}
	fmt.Println(r2)
	fmt.Println(reflect.TypeOf(r2))
}
