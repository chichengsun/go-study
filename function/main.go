package main

import (
	"fmt"
	"time"
)

func add(n1 int, n2 int) {
	fmt.Println(n1, n2)
}

// 参数类型一样，可以合并在一起
func add1(n1, n2 int) {
	fmt.Println(n1, n2)
}

// 多个参数
func add2(numList ...int) {
	fmt.Println(numList)
}

func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	return func(numList ...int) int {
		var sum int
		for _, i2 := range numList {
			sum += i2
		}
		return sum
	}
}

func pointer(num *int) {
	fmt.Println(num) // 内存值是一样的
	*num = 2         // 这里的修改会影响外面的num
}

func main() {
	add(1, 2)
	add1(1, 2)
	add2(1, 2)
	add2(1, 2, 3, 4)

	// 匿名函数
	var add = func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))

	fmt.Println(awaitAdd(2)(1, 2, 3))

	num := 10
	fmt.Println("before pointer:", num)
	pointer(&num)
	fmt.Println("after pointer:", num)
}
