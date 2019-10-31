package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	cnn, err := redis.DialURL("redis://root:123456@112.74.36.53:6379/0")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("连接成功")
	}
	defer cnn.Close()

	_, err = cnn.Do("SET", "username", "nick")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}
