package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Replace these values with your Redis server configuration
	redisHost := "localhost"
	redisPort := "6379"

	// Create a Redis client and connecting to reddis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // No password
		DB:       0,  // Default DB
	})

	//setting the value to myname as Gokul Sujan with 0 expiry time
	err := client.Set(context.Background(), "myname", "Gokul Sujan", 0).Err()
	if err != nil {
		fmt.Println("Error storing data in Redis:", err)
	}

	//getting the value from mykey
	value, err := client.Get(context.Background(), "mykey").Result()
	if err != nil {
		fmt.Println("Error retrieving data from Redis:", err)
	} else {
		fmt.Println("Value from Redis:", value)
	}

}
