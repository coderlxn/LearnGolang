package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	//使用slice作为key    错误
	//dict := map[[]string]int {} //

	//使用slice作为value
	dict := map[int][]string{}
	dict[123] = []string{"abc", "efg"}
	value, exists := dict[123]
	if exists {
		fmt.Println(value)
	}

	//使用delete删除键值对
	delete(dict, 123)

	value, exists = dict[123]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Value not exists")
	}

	//随机从map中选取一个值
	fruits := map[string]string{"apple": "app", "orange": "org", "banana": "ban", "cake": "cake", "mice": "mice"}
	count := map[string]int{"apple": 0, "orange": 0, "banana": 0, "cake": 0, "mice": 0}
	for i := 0; i < 999999; i++ {
		var key string
		for key = range fruits {
			//fmt.Println(key)
			count[key]++
			break
		}
	}
	fmt.Println(count)
	count = map[string]int{"apple": 0, "orange": 0, "banana": 0, "cake": 0, "mice": 0}
	keys := []string{"apple", "orange", "banana", "cake", "mice"}
	for i := 0; i < 999999; i++ {
		key := keys[rand.Intn(len(keys))]
		count[key]++
	}
	fmt.Println(count)

}
