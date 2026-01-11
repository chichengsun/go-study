// Package interface 演示Go语言中接口的动态类型、动态值以及==运算符的比较规则
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 1. 接口的动态类型和动态值详解
// 接口在Go中的内部表示通常包含两个字：
// - 第一个字：指向类型信息的指针（动态类型）
// - 第二个字：指向数据的指针（动态值）

// 接口的动态类型：存储在接口中的具体值的类型
// 接口的动态值：存储在接口中的具体值

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func demonstrateInterfaceInternals() {
	fmt.Println("=== 接口的动态类型和动态值 ===")

	var animal Animal

	// 1. nil接口：动态类型和动态值都为nil
	fmt.Printf("1. nil接口: animal == nil? %v\n", animal == nil)

	// 2. 非nil接口，但值为nil
	animal = (*Dog)(nil) // 动态类型是*Dog，但动态值是nil
	fmt.Printf("2. 非nil接口，但值为nil: animal == nil? %v\n", animal == nil)

	// 3. 完全非nil的接口
	animal = Dog{Name: "Buddy"}
	fmt.Printf("3. 完全非nil的接口: animal == nil? %v\n", animal == nil)

	// 使用反射查看接口的动态类型和动态值
	fmt.Println("\n使用反射查看接口内部:")

	animal = Dog{Name: "Buddy"}
	val := reflect.ValueOf(animal)
	fmt.Printf("动态类型: %v\n", val.Type())
	fmt.Printf("动态值: %v\n", val.Interface())

	animal = Cat{Name: "Whiskers"}
	val = reflect.ValueOf(animal)
	fmt.Printf("动态类型: %v\n", val.Type())
	fmt.Printf("动态值: %v\n", val.Interface())
}

// 2. ==运算符的比较规则详解

func demonstrateComparisonRules() {
	fmt.Println("\n=== ==运算符的比较规则 ===")

	// 2.1 基本类型的比较
	fmt.Println("\n2.1 基本类型的比较:")

	// 可比较的基本类型
	fmt.Printf("int: 5 == 5? %v\n", 5 == 5)
	fmt.Printf("float64: 3.14 == 3.14? %v\n", 3.14 == 3.14)
	fmt.Printf("string: \"hello\" == \"hello\"? %v\n", "hello" == "hello")
	fmt.Printf("bool: true == true? %v\n", true == true)

	// 不可比较的基本类型
	// slice、map、function、channel（除了nil比较）

	// 2.2 自定义类型的比较
	fmt.Println("\n2.2 自定义类型的比较:")

	// 可比较的自定义类型（所有字段都是可比较的）
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := Person{Name: "Bob", Age: 25}

	fmt.Printf("Person: p1 == p2? %v\n", p1 == p2) // true
	fmt.Printf("Person: p1 == p3? %v\n", p1 == p3) // false

	// 不可比较的自定义类型（包含不可比较的字段）
	type PersonWithSlice struct {
		Name    string
		Hobbies []string
	}

	// 下面这行会编译错误，因为包含slice字段，不可比较
	// fmt.Printf("PersonWithSlice: p4 == p5? %v\n", p4 == p5)

	// 2.3 接口的比较
	fmt.Println("\n2.3 接口的比较:")

	// 接口比较规则：
	// 1. 两个接口值相等，当且仅当它们的动态类型相同且动态值相等
	// 2. 或者两个接口值都为nil

	var animal1, animal2 Animal

	// 两个nil接口相等
	fmt.Printf("两个nil接口相等: %v\n", animal1 == animal2) // true

	// 动态类型和动态值都相同
	animal1 = Dog{Name: "Buddy"}
	animal2 = Dog{Name: "Buddy"}
	fmt.Printf("相同动态类型和值: %v\n", animal1 == animal2) // true

	// 动态类型相同，但动态值不同
	animal2 = Dog{Name: "Max"}
	fmt.Printf("相同动态类型，不同值: %v\n", animal1 == animal2) // false

	// 动态类型不同
	animal2 = Cat{Name: "Buddy"}
	fmt.Printf("不同动态类型: %v\n", animal1 == animal2) // false

	// 一个nil，一个非nil
	animal1 = nil
	fmt.Printf("一个nil，一个非nil: %v\n", animal1 == animal2) // false

	// 特殊情况：动态类型相同，但动态值都是nil
	animal1 = (*Dog)(nil)
	animal2 = (*Dog)(nil)
	fmt.Printf("动态类型相同，动态值都是nil: %v\n", animal1 == animal2) // true
}

// 3. 接口比较的陷阱
func demonstrateInterfaceComparisonPitfalls() {
	fmt.Println("\n=== 接口比较的陷阱 ===")

	// 陷阱1：包含不可比较类型的接口
	type UncomparableStruct struct {
		Data []int
	}

	var iface1, iface2 interface{}

	iface1 = UncomparableStruct{Data: []int{1, 2, 3}}
	iface2 = UncomparableStruct{Data: []int{1, 2, 3}}

	// 下面这行会panic，因为UncomparableStruct包含slice，不可比较
	// fmt.Printf("包含不可比较类型的接口: %v\n", iface1 == iface2)

	// 安全的比较方式
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("比较包含不可比较类型的接口时panic: %v\n", r)
		}
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("比较包含不可比较类型的接口时panic: %v\n", r)
			}
		}()
		_ = iface1 == iface2
	}()

	// 陷阱2：nil接口与非nil接口但值为nil的区别
	var animal Animal
	var dogPtr *Dog

	fmt.Printf("animal == nil? %v\n", animal == nil) // true
	fmt.Printf("dogPtr == nil? %v\n", dogPtr == nil) // true

	animal = dogPtr
	fmt.Printf("animal = dogPtr后, animal == nil? %v\n", animal == nil) // false!
	fmt.Printf("但是animal的动态值是nil吗? %v\n", animal == (*Dog)(nil))       // true
}

// 4. 使用reflect.DeepEqual进行深度比较
func demonstrateDeepEqual() {
	fmt.Println("\n=== 使用reflect.DeepEqual进行深度比较 ===")

	// 对于不可比较的类型，可以使用reflect.DeepEqual
	type PersonWithSlice struct {
		Name    string
		Hobbies []string
	}

	p1 := PersonWithSlice{Name: "Alice", Hobbies: []string{"reading", "coding"}}
	p2 := PersonWithSlice{Name: "Alice", Hobbies: []string{"reading", "coding"}}
	p3 := PersonWithSlice{Name: "Alice", Hobbies: []string{"reading"}}

	fmt.Printf("p1 == p2 (不可比较): 编译错误\n")
	fmt.Printf("reflect.DeepEqual(p1, p2): %v\n", reflect.DeepEqual(p1, p2)) // true
	fmt.Printf("reflect.DeepEqual(p1, p3): %v\n", reflect.DeepEqual(p1, p3)) // false

	// 对于接口，reflect.DeepEqual也会比较动态类型和动态值
	var animal1, animal2 Animal
	animal1 = Dog{Name: "Buddy"}
	animal2 = Dog{Name: "Buddy"}

	fmt.Printf("animal1 == animal2: %v\n", animal1 == animal2)                                   // true
	fmt.Printf("reflect.DeepEqual(animal1, animal2): %v\n", reflect.DeepEqual(animal1, animal2)) // true

	animal2 = Cat{Name: "Buddy"}
	fmt.Printf("animal1 == animal2: %v\n", animal1 == animal2)                                   // false
	fmt.Printf("reflect.DeepEqual(animal1, animal2): %v\n", reflect.DeepEqual(animal1, animal2)) // false
}

// 5. 接口的内部结构演示（使用unsafe包，仅用于演示）
func demonstrateInterfaceInternalStructure() {
	fmt.Println("\n=== 接口的内部结构演示 ===")

	// 注意：这里使用unsafe包仅用于演示，实际编程中不应该依赖这些实现细节
	var animal Animal = Dog{Name: "Buddy"}

	// 获取接口的内部表示
	iface := (*[2]uintptr)(unsafe.Pointer(&animal))

	fmt.Printf("接口内部表示:\n")
	fmt.Printf("  类型信息指针: 0x%x\n", iface[0])
	fmt.Printf("  数据指针: 0x%x\n", iface[1])

	// nil接口
	var nilAnimal Animal
	nilIface := (*[2]uintptr)(unsafe.Pointer(&nilAnimal))

	fmt.Printf("nil接口内部表示:\n")
	fmt.Printf("  类型信息指针: 0x%x\n", nilIface[0])
	fmt.Printf("  数据指针: 0x%x\n", nilIface[1])

	// 非nil接口，但值为nil
	var dogPtr *Dog
	animal = dogPtr
	iface = (*[2]uintptr)(unsafe.Pointer(&animal))

	fmt.Printf("非nil接口但值为nil的内部表示:\n")
	fmt.Printf("  类型信息指针: 0x%x\n", iface[0])
	fmt.Printf("  数据指针: 0x%x\n", iface[1])
}

func main() {
	demonstrateInterfaceInternals()
	demonstrateComparisonRules()
	demonstrateInterfaceComparisonPitfalls()
	demonstrateDeepEqual()
	demonstrateInterfaceInternalStructure()
}
