package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"redispool"
)

var exit chan struct{} = make(chan struct{})

func test() {
	conn := redispool.RedisPool.Get()
	defer conn.Close()

	v, err := conn.Do("SET", "pool-key-test", "goroutine-test-pool")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("goroutine is ", v)

	v, err = redis.String(conn.Do("GET", "pool-key-test"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("goroutine is ", v)
	close(exit)
}

func main() {
	conn := redispool.RedisPool.Get()
	defer conn.Close()
	go test()

	v, err := conn.Do("SET", "pool-key-client", "testRedisPoolClient")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

	v, err = redis.String(conn.Do("GET", "pool-key-client"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

	<-exit
}
