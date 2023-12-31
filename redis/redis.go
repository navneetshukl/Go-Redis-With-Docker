package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func RedisConnection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	} else {
		fmt.Println("Redis Ping Response:", pong)
	}

	return client
}

func InsertIntoRedis(key string, value interface{}) {
	client := RedisConnection()

	expirationTime := 100000

	err := client.Set(key, value, time.Duration(expirationTime)*time.Second).Err()

	if err != nil {
		fmt.Println("Error in inserting ", err)
		return

	}
	fmt.Println("Inserting Successful")
}

func GetFromRedis(key string) string {

	client := RedisConnection()

	ans, err := client.Get(key).Result()

	if err != nil {
		fmt.Println("Error in Fetching value from redis ", err)
	}
	fmt.Println("Value is ", ans)
	return ans
}
