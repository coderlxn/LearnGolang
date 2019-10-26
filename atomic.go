package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg1 sync.WaitGroup
)

func main()  {
	wg1.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	wg1.Wait()
}

func doWork(name string)  {
	defer wg1.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
