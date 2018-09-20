package main

import "fmt"

func main() {
	s := make([]int, 2)

	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", &s[0])
	fmt.Printf("%p\n", &s[1])
}