package main

import "fmt"

func main()  {
	var a int = 20
	var ptr *int
	var pptr **int

	ptr = &a

	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip 变量存储的指针地址: %x\n", ptr)
	fmt.Printf("*ip变量的值: %d\n", *ptr)

	pptr = &ptr
	if ptr == nil {
		fmt.Println("empty pointer")
	} else {
		fmt.Println("value of pointer of pointer ", **pptr)
	}
}
