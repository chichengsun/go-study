package main

import (
	"fmt"
)

func main() {
	// 只声明，不赋值，默认是0
	var v0 int

	// 先声明
	var v1 int
	v1 = 1

	// 声明的同时赋值
	var v2 int = 2

	// 自动推断
	var v3 = 3

	// 声明赋值
	v4 := 4

	fmt.Printf("v0=%v, v1=%v, v2=%v, v3=%v, v4=%v\n", v0, v1, v2, v3, v4)

	// 批量声明
	var (
		v5 int
		v6 int = 2
		v7     = 3
	)

	fmt.Printf("v1=%v, v2=%v, v3=%v", v5, v6, v7)

	//常量
	const cona int = 1
	fmt.Println(cona)

	const (
		a int = iota
		b     = 2
		c     = iota
		d     = iota
		e     = 1
		f     = iota
	)
	fmt.Println(a, b, c, d, e, f)
	// 0 2 2 3 1 5
	const (
		g = iota
		h = iota
	)
	fmt.Println(g, h)
	// 0 1

}
