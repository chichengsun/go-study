package main

import (
	"fmt"
)

func main() {
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)
	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)
	var array2 = [...]int{1, 2, 3}
	fmt.Println(array2)

	// 说明：[...]int 是“推断长度的数组类型”，最终类型仍是 [3]int，不是切片。
	fmt.Printf("类型: array=%T, array1=%T, array2=%T\n", array, array1, array2)

	array[0] = 4
	fmt.Println(array)

	// 语法与类型：数组类型包含长度，[3]int；切片不包含长度，[]int。数组长度固定且是类型的一部分，切片长度可变。
	// 值语义 vs 引用语义：数组是值类型，赋值/传参会拷贝整份数据；切片是对底层数组的引用视图（指针+len+cap），赋值/传参共享底层数据。
	// 修改影响：修改切片元素会影响底层数组；数组拷贝彼此独立，修改副本不影响原数组。
	// 零值与创建：数组零值为各元素的零值；切片零值为 nil（len=0、cap=0），常用 make 创建并可 append。
	// 追加与扩容：数组不可追加；切片可 append，必要时扩容并可能更换底层数组。
	// 可比较性：数组可用 == 比较（元素可比较时）；切片只能与 nil 比较，不能用 == 比较内容。
	// 使用场景：固定大小数据用数组；通用可变长度集合、函数参数传递等用切片。

	// 数组是值类型（赋值会拷贝整份数据，长度固定），切片是对底层数组的引用视图（长度可变，cap 反映底层数组容量）。
	// --- slice 基础 ---
	// 1) 通过字面量创建
	s1 := []int{10, 20, 30}
	fmt.Println("slice literal:", s1, "len:", len(s1), "cap:", cap(s1))

	// 2) 由数组切片
	s2 := array[:2] // 使用上面的 array
	fmt.Println("slice from array:", s2, "len:", len(s2), "cap:", cap(s2))
	// 切片修改会影响底层数组
	s2[0] = 99
	fmt.Println("slice change affects array:", s2, "array:", array)

	// 3) 通过 make 指定 len 和 cap
	s3 := make([]int, 0, 5)
	s3 = append(s3, 1, 2, 3)
	fmt.Println("slice via make+append:", s3, "len:", len(s3), "cap:", cap(s3))

	// 4) 追加后容量变化
	s3 = append(s3, 4, 5, 6)
	fmt.Println("slice after grow:", s3, "len:", len(s3), "cap:", cap(s3))

	// 数组赋值是拷贝，修改副本不影响原数组
	arrayCopy := array
	arrayCopy[0] = 888
	fmt.Println("array copy:", arrayCopy, "original array:", array)

	// --- map 基础 ---
	// 1) 字面量创建
	m1 := map[string]int{"apple": 3, "banana": 5}
	fmt.Println("map literal:", m1)

	// 2) make 创建并赋值
	m2 := make(map[string]string)
	m2["name"] = "gopher"
	m2["lang"] = "go"
	fmt.Println("map via make:", m2)

	// 3) 访问与存在性判断
	if val, ok := m2["name"]; ok {
		fmt.Println("name exists:", val)
	}

	// 4) 删除键
	delete(m2, "lang")
	fmt.Println("map after delete:", m2)
}
