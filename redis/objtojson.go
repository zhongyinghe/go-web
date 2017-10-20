package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (user User) SayName() {
	fmt.Println("My Name is " + user.Name)
}

func (user User) SayAge() {
	fmt.Println("My Age is " + strconv.Itoa(user.Age))
}

func main() {
	user := User{Name: "zhongyinghe", Age: 32}
	//user.SayName()
	//user.SayAge()
	jsonVal, _ := json.Marshal(user)
	fmt.Println(jsonVal)

	/*n_user := User{}
	err := json.Unmarshal(jsonVal, &n_user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n_user)
	n_user.SayName()
	n_user.SayAge()*/

	//redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()
	c.Do("AUTH", "123456") //授权

	key := "redis_user"
	n, err := c.Do("SETNX", key, jsonVal)
	if err != nil {
		fmt.Println(err)
		return
	}

	if n == int64(1) {
		fmt.Println("Success")
	}

	valueGet, err := redis.Bytes(c.Do("GET", key))

	if err != nil {
		fmt.Println(err)
		return
	}

	user2 := User{}
	errShal := json.Unmarshal(valueGet, &user2)
	if errShal != nil {
		fmt.Println(errShal)
		return
	}

	fmt.Println(user2)
	user2.SayName()
	user2.SayAge()
}
