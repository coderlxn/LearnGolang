package main

import (
	"fmt"
	"math"
)

//闭包
func getSquence() func()  int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//带参数的闭包
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
	i := 0
	return func(x3 int,x4 int) (int,int,int){
		i++
		return i,x1+x2,x3+x4
	}
}

func main()  {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	//函数作为值
	fmt.Println(getSquareRoot(9))

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSquence()

	// 调用 nextNumber 函数，i 变量自增 1 并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSquence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())


	//带参数的闭包
	addFunc := add(1,2)
	fmt.Println(addFunc(1,1))
	fmt.Println(addFunc(0,0))
	fmt.Println(addFunc(2,2))
}
