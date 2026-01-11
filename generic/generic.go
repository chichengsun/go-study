// Package generic 演示Go语言中的泛型知识
package main

import (
	"fmt"
)

// 自定义约束接口，替代标准库中的constraints包

// Ordered 可排序类型约束
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 | ~string
}

// Integer 整数类型约束
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float 浮点数类型约束
type Float interface {
	// 波浪线表示底层类型，基本类型的底层类型就是他自己，而自定义类型，底层类型就是它基于的基本类型
	// 接受float32类型
	// 接受float64类型
	// 接受任何底层类型是float32的自定义类型
	// 接受任何底层类型是float64的自定义类型
	~float32 | ~float64
}

// 1. 泛型基础
// Go 1.18引入了泛型，允许编写类型参数化的函数和类型
// 泛型使用方括号[]定义类型参数

// 泛型函数示例：打印任意类型的切片
func PrintSlice[T any](s []T) {
	fmt.Printf("切片类型: %T, 内容: %v\n", s, s)
}

// 泛型函数示例：返回两个值中的较大者
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 2. 类型约束
// 类型约束限制了类型参数可以使用的具体类型
// Go提供了一些预定义的约束，如constraints.Ordered, constraints.Integer等

// 使用内置约束的泛型函数
func SumNumbers[T Integer | Float](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 3. 自定义类型约束
// 可以使用接口定义自定义类型约束

// Numberish 约束：包含所有数值类型
type Numberish interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~complex64 | ~complex128
}

// Add 泛型函数：使用自定义约束
func Add[T Numberish](a, b T) T {
	return a + b
}

// 4. 泛型结构体
// 结构体也可以使用类型参数

// Stack 泛型栈实现
type Stack[T any] struct {
	elements []T
}

// Push 向栈中添加元素
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop 从栈中弹出元素
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

// IsEmpty 检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size 返回栈的大小
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

// 5. 泛型接口
// 接口也可以使用类型参数

// Comparable 可比较接口
type Comparable[T any] interface {
	Compare(other T) int
}

// 6. 泛型与方法的组合
// 泛型类型可以有方法，这些方法可以使用类型的参数

// LinkedList 泛型链表实现
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Node 链表节点
type Node[T any] struct {
	value T
	next  *Node[T]
}

// NewLinkedList 创建新的链表
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add 在链表末尾添加元素
func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

// Get 获取指定索引的元素
func (l *LinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, false
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, true
}

// Size 返回链表大小
func (l *LinkedList[T]) Size() int {
	return l.size
}

// 7. 多个类型参数
// 泛型函数和类型可以有多个类型参数

// Pair 泛型对类型
type Pair[K, V any] struct {
	Key   K
	Value V
}

// Map 泛型映射函数
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter 泛型过滤函数
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 泛型归约函数
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// 8. 类型推断
// Go可以自动推断泛型函数的类型参数

func demonstrateTypeInference() {
	fmt.Println("=== 类型推断演示 ===")

	// 不需要显式指定类型参数，Go会自动推断
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("整数切片: %v\n", intSlice)

	// 自动推断T为int
	sum := SumNumbers(intSlice)
	fmt.Printf("整数和: %d\n", sum)

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Printf("浮点数切片: %v\n", floatSlice)

	// 自动推断T为float64
	sumFloat := SumNumbers(floatSlice)
	fmt.Printf("浮点数和: %.2f\n", sumFloat)

	// 自动推断T为string
	strSlice := []string{"hello", "world", "go"}
	PrintSlice(strSlice)
}

// 9. 泛型的实际应用场景

// 9.1 通用的数据结构
func demonstrateGenericDataStructures() {
	fmt.Println("\n=== 泛型数据结构演示 ===")

	// 使用泛型栈
	fmt.Println("泛型栈:")
	intStack := &Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("弹出: %d\n", val)
		}
	}

	// 使用泛型链表
	fmt.Println("\n泛型链表:")
	stringList := NewLinkedList[string]()
	stringList.Add("Go")
	stringList.Add("泛型")
	stringList.Add("编程")

	for i := 0; i < stringList.Size(); i++ {
		if val, ok := stringList.Get(i); ok {
			fmt.Printf("索引 %d: %s\n", i, val)
		}
	}
}

// 9.2 通用的算法
func demonstrateGenericAlgorithms() {
	fmt.Println("\n=== 泛型算法演示 ===")

	// 使用Map函数
	numbers := []int{1, 2, 3, 4, 5}
	squared := Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("平方后: %v\n", squared)

	// 使用Filter函数
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("偶数: %v\n", evenNumbers)

	// 使用Reduce函数
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("总和: %d\n", sum)
}

// 9.3 通用的工具函数
func demonstrateGenericUtilities() {
	fmt.Println("\n=== 泛型工具函数演示 ===")

	// 查找切片中的最大值
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"hello\", \"world\"): %s\n", Max("hello", "world"))

	// 数值相加
	fmt.Printf("Add(10, 20): %d\n", Add(10, 20))
	fmt.Printf("Add(3.14, 2.71): %.2f\n", Add(3.14, 2.71))
}

// 10. 泛型的限制和注意事项
func demonstrateGenericLimitations() {
	fmt.Println("\n=== 泛型的限制和注意事项 ===")

	// 1. 不能将泛型类型与具体类型进行断言
	// var x interface{} = Stack[int]{}
	// if s, ok := x.(Stack[int]); ok { // 这是可以的
	//     fmt.Println("这是一个整数栈")
	// }

	// 2. 不能使用泛型类型作为map的键
	// m := make(map[Stack[int]]string) // 编译错误：Stack[int]不可比较

	// 3. 泛型方法不能有类型参数
	// func (s *Stack[T]) PushMany[U any](values ...U) { // 编译错误
	//     // ...
	// }

	// 4. 泛型类型不能有方法，这些方法使用不同的类型参数
	// func (s *Stack[T]) ConvertTo[U any]() []U { // 编译错误
	//     // ...
	// }

	fmt.Println("泛型的限制和注意事项已注释在代码中")
}

// 11. 泛型与反射的结合
func demonstrateGenericsWithReflection() {
	fmt.Println("\n=== 泛型与反射的结合 ===")

	// 泛型可以与反射结合使用，但需要注意类型安全
	// 这里只是演示概念，实际使用中要谨慎

	fmt.Println("泛型与反射的结合使用需要谨慎，确保类型安全")
}

// 12. 性能考虑
func demonstratePerformanceConsiderations() {
	fmt.Println("\n=== 泛型性能考虑 ===")

	// 1. 编译时特化：Go编译器会为每个使用的具体类型生成特化版本
	//    这意味着泛型代码在运行时没有额外的性能开销

	// 2. 内存使用：泛型不会引入额外的内存开销

	// 3. 编译时间：使用泛型可能会增加编译时间，因为需要为每个具体类型生成代码

	// 4. 二进制大小：可能会增加二进制文件大小，因为包含了多个特化版本

	fmt.Println("泛型在运行时没有额外性能开销，但可能增加编译时间和二进制大小")
}

func main() {
	// 基础泛型演示
	fmt.Println("=== 基础泛型演示 ===")

	// 使用PrintSlice泛型函数
	intSlice := []int{1, 2, 3, 4, 5}
	PrintSlice(intSlice)

	stringSlice := []string{"Go", "泛型", "编程"}
	PrintSlice(stringSlice)

	// 使用Max泛型函数
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))

	// 使用SumNumbers泛型函数
	fmt.Printf("SumNumbers([]int{1, 2, 3, 4, 5}): %d\n", SumNumbers([]int{1, 2, 3, 4, 5}))
	fmt.Printf("SumNumbers([]float64{1.1, 2.2, 3.3}): %.2f\n", SumNumbers([]float64{1.1, 2.2, 3.3}))

	// 使用Add泛型函数
	fmt.Printf("Add(10, 20): %d\n", Add(10, 20))
	fmt.Printf("Add(3.14, 2.71): %.2f\n", Add(3.14, 2.71))

	// 演示其他泛型概念
	demonstrateTypeInference()
	demonstrateGenericDataStructures()
	demonstrateGenericAlgorithms()
	demonstrateGenericUtilities()
	demonstrateGenericLimitations()
	demonstrateGenericsWithReflection()
	demonstratePerformanceConsiderations()
}
