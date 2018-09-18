package main

import "fmt"

//递归


func Factorial(n uint64) uint64 {
	if n > 0 {
		result := n * Factorial(n - 1)
		return result
	}
	return 1
}

func main()  {
	var i int = 15
	fmt.Printf("%d的阶乘是 %d\n", i, Factorial(uint64(i)))
}
