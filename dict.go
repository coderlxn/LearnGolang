package main

import "fmt"

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

}
