package main

import (
	"fmt"
)

func main() {
	// https://docs.fengfengzhidao.com/#/docs/%E6%96%B0golang%E5%9F%BA%E7%A1%80/1.%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA
	fmt.Println("hello world")
	fmt.Println("Please input your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Your name is: " + name)
}
