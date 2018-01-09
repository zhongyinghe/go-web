package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisPool *redis.Pool
)

func NewPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
	}
}

func main() {
	redisPool = NewPool("127.0.0.1:6379", "123456")
	conn := redisPool.Get()
	defer conn.Close()

	v, err := conn.Do("SET", "pool-key", "testRedisPool")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

	v, err = redis.String(conn.Do("GET", "pool-key"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
}
