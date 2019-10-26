package main

import (
	"fmt"
	"sync"
	"time"
)

var noCacheWaitGroup sync.WaitGroup

func main()  {

	baton := make(chan int)

	noCacheWaitGroup.Add(1)

	go Runner(baton)

	baton <- 1

	noCacheWaitGroup.Wait()
}

func Runner(baton chan int)  {
	var newRunner int
	//等待接力棒
	runner := <-baton

	fmt.Printf("Runner %d Running With Baton\n", runner)

	//创建下一位跑步者
	if runner != 4{
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	//围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	if runner == 4{
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		noCacheWaitGroup.Done()
		return
	}

	//接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchanged With Runner %d\n", runner, newRunner)

	baton <- newRunner
}