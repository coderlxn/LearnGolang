package main

import "fmt"

func main()  {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("c: %d\n",c)
	c = a - b
	fmt.Printf("c: %d\n",c)
	c = a * b
	fmt.Printf("c: %d\n",c)
	c = a / b
	fmt.Printf("c: %d\n",c)
	c = a % b
	fmt.Printf("c: %d\n",c)
	a++
	//go 不支持a++作为参数
	fmt.Printf("a: %d\n",a)
	a--
	fmt.Printf("a: %d\n",a)

	//指针的赋值和解引用
	var ptr *int
	ptr = &a
	fmt.Printf("ptr: %d\n", *ptr)
}