package main

import (
	"math/rand"
	"time"
	"log"
)

// PasswordLength 密码长度
const PasswordLength = 256

// Password 密码内容
type Password [PasswordLength]byte

func init() {
	//
	log.Println("init func called")
	rand.Seed(time.Now().Unix())
}

// RandPassword 生成密码
func RandPassword() *Password {
	intArr := rand.Perm(PasswordLength)
	password := &Password{}
	for i, v := range intArr {
		password[i] = byte(v)
		for i == v {
			return RandPassword()
		}
	}
	return password
}

func main()  {
	password := RandPassword()
	for i, v := range password {
		log.Println("pass index ", i, " pass value ", v)
	}
}