package main

import (
	"fmt"

	"github.com/kataras/iris/sessions/sessiondb/redis"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human
	comany string
}

// SayHi 在human上定义一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s", h.name, h.phone)
}

func main() {
	fmt.Println("Hello Go")
	redisConn = redis.cl(&redis.Config{
		Addr:     "localhost",
		Password: "",
	})
}
