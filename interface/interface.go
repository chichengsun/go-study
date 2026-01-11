// Package interface 演示Go语言中接口的使用
package main

import (
	"fmt"
	"math"
)

// 1. 接口的基本定义
// 接口是一种抽象类型，它定义了一组方法签名
// 任何类型只要实现了接口中定义的所有方法，就被认为实现了该接口

// Shaper 定义了一个形状接口，包含Area方法
type Shaper interface {
	Area() float64
}

// 2. 接口的实现
// Rectangle 结构体实现了Shaper接口
type Rectangle struct {
	Width, Height float64
}

// Area 方法实现了Shaper接口的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 结构体也实现了Shaper接口
type Circle struct {
	Radius float64
}

// Area 方法实现了Shaper接口的Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 3. 空接口
// 空接口interface{}没有任何方法，所以所有类型都实现了空接口
// 空接口可以存储任何类型的值

func printValue(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// 4. 类型断言
// 类型断言用于检查接口值是否包含特定类型的值

func typeAssertionDemo() {
	var shaper Shaper = Rectangle{Width: 3, Height: 4}

	// 安全的类型断言
	if rect, ok := shaper.(Rectangle); ok {
		fmt.Printf("Rectangle width: %.2f, height: %.2f\n", rect.Width, rect.Height)
	}

	// 不安全的类型断言（如果失败会panic）
	// circle := shaper.(Circle) // 这会panic，因为shaper实际上是Rectangle类型
}

// 5. 类型选择
// 类型选择是一种按顺序从几个类型断言中选择分支的结构

func typeSwitchDemo() {
	var shaper Shaper

	shaper = Rectangle{Width: 5, Height: 6}
	switch v := shaper.(type) {
	case Rectangle:
		fmt.Printf("Rectangle with width %.2f and height %.2f\n", v.Width, v.Height)
	case Circle:
		fmt.Printf("Circle with radius %.2f\n", v.Radius)
	default:
		fmt.Printf("Unknown shape: %T\n", v)
	}

	shaper = Circle{Radius: 7}
	switch v := shaper.(type) {
	case Rectangle:
		fmt.Printf("Rectangle with width %.2f and height %.2f\n", v.Width, v.Height)
	case Circle:
		fmt.Printf("Circle with radius %.2f\n", v.Radius)
	default:
		fmt.Printf("Unknown shape: %T\n", v)
	}
}

// 6. 接口组合
// 接口可以通过组合其他接口来创建新的接口

// Writer 接口定义了写入方法
type Writer interface {
	Write([]byte) (int, error)
}

// Closer 接口定义了关闭方法
type Closer interface {
	Close() error
}

// WriteCloser 接口组合了Writer和Closer接口
type WriteCloser interface {
	Writer
	Closer
}

// 7. 接口的值
// 接口的值由两部分组成：具体的类型和该类型的值
// 这被称为接口的动态类型和动态值

func interfaceValueDemo() {
	var shaper Shaper

	fmt.Printf("shaper is nil: %v\n", shaper == nil)

	shaper = Rectangle{Width: 10, Height: 20}
	fmt.Printf("shaper value: %v, type: %T\n", shaper, shaper)

	// 将接口值设为nil
	shaper = nil
	fmt.Printf("shaper is nil again: %v\n", shaper == nil)
}

// 8. 接口的比较
// 接口值可以使用==和!=进行比较
// 只有当两个接口值的动态类型相同且动态值相等（或都为nil）时，它们才相等

func interfaceComparisonDemo() {
	var shaper1, shaper2, shaper3, shaper4 Shaper

	shaper1 = Rectangle{Width: 3, Height: 4}
	shaper2 = Rectangle{Width: 3, Height: 4}
	shaper3 = (*Circle)(nil)
	shaper4 = (*Circle)(nil)

	fmt.Printf("shaper1 == shaper2: %v\n", shaper1 == shaper2)

	shaper2 = Rectangle{Width: 4, Height: 3}
	fmt.Printf("shaper1 == shaper2: %v\n", shaper1 == shaper2)

	fmt.Printf("shaper3 == shaper4: %v\n", shaper3 == shaper4)
}

// 9. 接口与指针接收者
// 如果接口方法使用指针接收者实现，则只有指针类型实现了该接口

type PointerReceiver struct {
	Value int
}

func (p *PointerReceiver) GetValue() int {
	return p.Value
}

// ValueGetter 接口
type ValueGetter interface {
	GetValue() int
}

func pointerReceiverDemo() {
	// 只有*PointerReceiver实现了ValueGetter接口
	var getter ValueGetter = &PointerReceiver{Value: 42}
	fmt.Printf("Value: %d\n", getter.GetValue())

	// 下面这行会编译错误，因为PointerReceiver没有实现ValueGetter接口
	// var getter2 ValueGetter = PointerReceiver{Value: 42}
}

// 10. 接口的嵌套
// 接口可以嵌套在其他接口中

type Reader interface {
	Read([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	fmt.Println("=== Go Interface Demo ===")

	// 1. 基本接口使用
	fmt.Println("\n1. Basic Interface Usage:")
	shapes := []Shaper{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Rectangle{Width: 2, Height: 6},
	}

	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f\n", shape, shape.Area())
	}

	// 2. 空接口
	fmt.Println("\n2. Empty Interface:")
	printValue(42)
	printValue("hello")
	printValue(Rectangle{Width: 1, Height: 2})

	// 3. 类型断言
	fmt.Println("\n3. Type Assertion:")
	typeAssertionDemo()

	// 4. 类型选择
	fmt.Println("\n4. Type Switch:")
	typeSwitchDemo()

	// 5. 接口的值
	fmt.Println("\n5. Interface Values:")
	interfaceValueDemo()

	// 6. 接口的比较
	fmt.Println("\n6. Interface Comparison:")
	interfaceComparisonDemo()

	// 7. 指针接收者
	fmt.Println("\n7. Pointer Receiver:")
	pointerReceiverDemo()
}
