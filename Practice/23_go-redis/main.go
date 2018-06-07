package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	rclnt := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rclnt.Set("KEY", "VALUE1", 0).Err()
	if err != nil {
		fmt.Println("Set:", err)
	}
	val, err := rclnt.Get("KEY").Result()
	if err != nil {
		fmt.Println("Get:", err)
	}
	fmt.Println("KEY:", val)

}
