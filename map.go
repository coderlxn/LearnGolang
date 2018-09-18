package main

import "fmt"

//不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

func main()  {
	var countryCaptialMap map[string]string
	countryCaptialMap = make(map[string]string)

	countryCaptialMap["France"] = "Pairs"
	countryCaptialMap["Italy"] = "罗马"
	countryCaptialMap["Japan"] = "东京"
	countryCaptialMap["India"] = "新德里"

	//键值输出
	for country := range countryCaptialMap {
		fmt.Println(country, "首都是", countryCaptialMap[country])
	}

	//查看元素是否存在
	captial, ok := countryCaptialMap["美国"]
	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}

	//删除元素
	delete(countryCaptialMap, "Japan")
	for country := range countryCaptialMap {
		fmt.Println(country, "首都是", countryCaptialMap[country])
	}

	//删除不存在的元素
	delete(countryCaptialMap, "China")
	for country := range countryCaptialMap {
		fmt.Println(country, "首都是", countryCaptialMap[country])
	}
}