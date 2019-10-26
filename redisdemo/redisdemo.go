package main

import "fmt"
import "github.com/go-redis/redis"


func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379",
				Password:"",
				DB: 0})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379",
		Password:"",
		DB: 0})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err.Error())
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("key2", val2)
	}
}

func readUserDomain()  {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.1.60:6379",
		Password:"",
		DB:0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("60 redis", pong)
	val, err := client.HGet("user:0d70e88d7ff0cf72233733bd03572dae", "company_domain").Result()
	if err != nil {
		panic(err.Error())
	}
	println("user domain", val)
}

func multiRedisTest() {

}

func main() {
	ExampleNewClient()
	ExampleClient()
	readUserDomain()
}
