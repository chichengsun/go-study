package main

import (
	"fmt"
	"os"
)

//init 函数特点：
//
//无需手动调用，Go 运行时自动执行
//同一包中多个 init 函数按声明顺序执行
//不同包的 init 函数按导入顺序执行
//每个包的 init 函数在 main 函数之前执行
//
//defer 函数特点：
//
//延迟执行，在包含它的函数返回时执行
//多个 defer 按后进先出 (LIFO) 顺序执行
//defer 函数的参数在声明时就会求值
//常用于资源清理（如关闭文件、释放锁等）

// 第一个 init 函数
func init() {
	fmt.Println("main package: first init() function")
}

// 第二个 init 函数
func init() {
	fmt.Println("main package: second init() function")
}

func main() {
	fmt.Println("main function starts")

	// defer 基本用法
	defer fmt.Println("defer 1: executed when main() exits")
	defer fmt.Println("defer 2: executed when main() exits")

	// defer 与函数参数
	x := 10
	defer fmt.Printf("defer with parameter: x = %d\n", x) // 这里会立即求值x=10
	x = 20

	// defer 与匿名函数
	defer func() {
		fmt.Printf("defer with anonymous function: x = %d\n", x) // 这里会在执行时求值x=20
	}()

	// 打开文件示例，defer 用于关闭资源
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // 确保文件会被关闭
	fmt.Println("File created successfully")

	// 写入文件
	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file")

	fmt.Println("main function ends")
}
