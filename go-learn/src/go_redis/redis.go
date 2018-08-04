package go_redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"encoding/json"
)

func RedisClient()  {
	fmt.Println("==== redis client ====")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	println(pong)

	value, err := json.Marshal([]string{"a","b"})
	err = client.Set("products", string(value),0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("products").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("product", val)
}