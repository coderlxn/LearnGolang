package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	mutexCounter int
	mutexWg sync.WaitGroup
	mutex sync.Mutex
)

func main()  {
	mutexWg.Add(2)

	go mutexIncCounter(1)
	go mutexIncCounter(2)

	mutexWg.Wait()
	fmt.Println("Final Counter:", mutexCounter)
}

func mutexIncCounter(id int)  {
	defer mutexWg.Done()

	for count := 0; count < 2; count++ {

		mutex.Lock()
		{
			value := mutexCounter
			runtime.Gosched()

			value++
			mutexCounter = value
		}
		mutex.Unlock()
	}
}
