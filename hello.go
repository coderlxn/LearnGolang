package main

var a = "菜鸟教程"
var b string = "bunnno"
var c bool

var g int

func main()  {
	const LENGTH = 10
	const WIDTH = 5
	var area int

	area = LENGTH * WIDTH

	const (
		iota1 = iota
		iota2
		iota3
	)

	const (
		samevalue = 111
		samevalue1
		samevalue2
		samevalue3
	)


	println("Hello, go!", a, b, c, area)
	println(iota1, iota2, iota3)
	println(samevalue, samevalue1, samevalue2, samevalue3)

	var strx, stry string
	strx = "mahesh"
	stry = "Kumar"
	strx, stry = stry, strx
	println(strx, stry)
}