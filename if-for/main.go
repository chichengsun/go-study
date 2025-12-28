package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
	}
	if age > 0 && age <= 18 {
		fmt.Println("未成年")
	}
	if age > 18 && age <= 35 {
		fmt.Println("青年")
	}
	if age > 35 {
		fmt.Println("中年")
	}

	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
		// fallthrough
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}

	// 传统for循环
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while模式
	i := 0
	sum = 0
	for i <= 100 {
		sum += i
		i++ // go中只有后++
	}
	fmt.Println(sum)

	s := []string{"枫枫", "知道"}
	for index, s2 := range s {
		fmt.Println(index, s2)
	}

	smap := map[string]int{
		"age":   24,
		"price": 1000,
	}
	for key, val := range smap {
		fmt.Println(key, val)
	}
}
