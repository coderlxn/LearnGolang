package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgNoCache sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wgNoCache.Add(2)

	//两个选手
	go player("Nadal", court)
	go player("Djokovic", court)

	//发球
	court <- 1

	wgNoCache.Wait()
}

func player(name string, court chan int)  {
	//
	defer wgNoCache.Done()

	for {
		//等待球被击打过来
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			//关闭通道，我们输了
			close(court)
			return
		}

		//显示击球数
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		//将球打向对手
		court <- ball
	}
}