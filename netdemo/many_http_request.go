package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main()  {
	var waiter sync.WaitGroup
	const requestCount = 5
	waiter.Add(requestCount)
	for i := 0; i < requestCount; i++ {
		go func(index int) {
			userID := rand.Int()
			begin := time.Now().Unix()
			fmt.Println("start request user info : ", index)
			url := fmt.Sprintf("http://localhost:8880/user/%d", userID)
			r, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(userID, string(body))
			}
			fmt.Println("finish on request : ", index, " take time : ", time.Now().Unix() - begin)
			waiter.Done()
		}(i)
	}
	waiter.Wait()
}
