package main

import (
	"fmt"
)

func main() {
	/* 基本数据类型
	 * 整数：
	 * int8
	 * int16
	 * int32
	 * int64
	 * int
	 * uint
	 * uint8 byte别名（单字符类型）
	 * uint16
	 * uint32 rune别名（多字符类型）
	 * uint64
	 *
	 * 浮点数：
	 * float32
	 * float64
	 *
	 * 字符串：
	 * string
	 *
	 * bool：
	 * bool true false
	 */

	// 基本数据类型变量仅声明不赋值，那么变量的值就是类型的零值，比如int是0，bool是false，string是""

	fmt.Println("基本数据类型")

	// --- 代码示例 ---

	// 1. 整数类型
	var i int = 100
	var i8 int8 = 127
	// int8 范围是 -128 到 127
	fmt.Printf("整数: int=%d, int8=%d\n", i, i8)

	// 2. 浮点类型
	var f32 float32 = 3.14
	var f64 float64 = 3.1415926535
	fmt.Printf("浮点数: float32=%f, float64=%.10f\n", f32, f64)

	// 3. 布尔类型
	var isActive bool = true
	var isEnabled bool = false
	fmt.Printf("布尔值: %t, %t\n", isActive, isEnabled)

	// 4. 字符串
	var str string = "Hello Golang"
	fmt.Printf("字符串: %s\n", str)

	// 5. byte 和 rune
	// byte 是 uint8 的别名，用于表示 ASCII 字符
	var a byte = 'A'
	// rune 是 int32 的别名，用于表示 Unicode 字符（如中文）
	var zhong rune = '中'
	fmt.Printf("byte: %c (值:%v), rune: %c (值:%v)\n", a, a, zhong, zhong)

	// 6. 零值演示 (声明但未赋值)
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	fmt.Printf("零值演示: int=%d, float=%v, bool=%v, string=%q\n",
		zeroInt, zeroFloat, zeroBool, zeroString)
}
