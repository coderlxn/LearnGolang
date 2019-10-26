package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func main()  {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int)  {
	defer wg.Done()

	for count := 0; count < 2; count++ {

		//下面代码会导致竟态问题
		//value := counter
		////当前goroutine从线程退出，并放回到队列
		//runtime.Gosched()
		//value++
		//counter = value

		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}