package redispool

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	RedisPool *redis.Pool
	server    string = "127.0.0.1:6379"
	password  string = "123456"
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

func init() {
	RedisPool = NewPool(server, password)
}
