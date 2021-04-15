package main

import "fmt"

//如果某个interface含有多个方法，如果想通过interface来调用方法名称来调用对象，需要所有方法都实现
//如果只通过具体的类型来调用对象方法，可以不全部实现。这种情况下对象并没有实现这个接口

type Phone interface {
	call()
	hangup()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	phone := new(NokiaPhone)
	phone.call()
	// var phone Phone

	// phone = new(NokiaPhone)
	// phone.call()

	// phone = new(IPhone)
	// phone.call()
}
