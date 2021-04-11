package main

func main()  {
	const (
		a = iota  //开始计数
		b
		c
		d = "ha"	//计数暂停
		e
		f = 100
		g
		h = iota	//恢复计数
		i
	)
	println(a, b, c, d, e, f, g , h, i)

	const (
		j = 1 << iota	// 1 << 0
		k = 3 << iota	// 3 << 1
		l	// 3 << 2
		m	// 3 << 3
	)
	println(j, k, l, m)

}
