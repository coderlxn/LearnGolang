package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			fmt.Println("return a value ", i)
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}

	// channel数组
	var chanList [2]chan int
	chanList[0] = make(chan int)
	chanList[1] = make(chan int)

	go func() {
		for i := 0; i < 10; i = i + 1 {
			chanList[0] <- i
		}
		close(chanList[0])
	}()
	for i := range chanList[0] {
		fmt.Println(i)
	}

	// channel用作数据同步，不关系管道内传输的数据，一遍使用无类型的匿名结构体
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<-c2

	fmt.Println("Finished")
}
