package LearnGolang

import "fmt"

func TestValues() {
	var n [10]int //长度为10的数组

	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Println("value of slice ", slice[i])
	}
}
