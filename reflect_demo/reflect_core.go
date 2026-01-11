// Package reflect_demo 提供Go反射的核心知识点总结
package main

import (
	"fmt"
	"reflect"
)

// Calculator 用于演示方法调用
type Calculator struct{}

// Add 方法
func (c Calculator) Add(a, b int) int { return a + b }

// 核心概念演示
func main() {
	fmt.Println("=== Go反射核心知识点 ===\n")

	// 1. Type vs Value
	fmt.Println("1. Type vs Value")
	fmt.Println("   Type: 类型信息，类似Java的Class")
	fmt.Println("   Value: 值信息，包含类型和具体值")

	var x int = 42
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)

	fmt.Printf("   ValueOf(%v) -> %v (Kind: %v)\n", x, v, v.Kind())
	fmt.Printf("   TypeOf(%v) -> %v (Kind: %v)\n", x, t, t.Kind())

	// 2. Kind的分类
	fmt.Println("\n2. Kind的分类")
	fmt.Println("   Kind是比Type更基础的类型分类")

	kinds := []interface{}{
		42,                     // Int
		"hello",                // String
		[]int{1, 2},            // Slice
		map[string]int{"a": 1}, // Map
		struct{}{},             // Struct
		(*int)(nil),            // Ptr
		make(chan int),         // Chan
		func() {},              // Func
	}

	for _, val := range kinds {
		v := reflect.ValueOf(val)
		fmt.Printf("   %T -> Kind: %v\n", val, v.Kind())
	}

	// 3. 修改值的条件
	fmt.Println("\n3. 修改值的条件")
	fmt.Println("   - 必须通过指针传递")
	fmt.Println("   - 使用Elem()获取指针指向的值")
	fmt.Println("   - 使用CanSet()检查是否可修改")

	num := 10
	fmt.Printf("   修改前: %d\n", num)

	vNum := reflect.ValueOf(&num).Elem() // 关键：指针 + Elem
	if vNum.CanSet() {
		vNum.SetInt(100)
		fmt.Printf("   修改后: %d\n", num)
	}

	// 4. 结构体操作
	fmt.Println("\n4. 结构体操作")
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	vp := reflect.ValueOf(p)

	// 获取字段数量
	fmt.Printf("   字段数量: %d\n", vp.NumField())

	// 遍历字段
	for i := 0; i < vp.NumField(); i++ {
		field := vp.Field(i)
		fieldType := vp.Type().Field(i)
		fmt.Printf("   %s = %v (标签: %s)\n",
			fieldType.Name, field.Interface(), fieldType.Tag.Get("json"))
	}

	// 修改结构体字段
	p2 := Person{Name: "Bob", Age: 25}
	vp2 := reflect.ValueOf(&p2).Elem()

	nameField := vp2.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Charlie")
	}

	fmt.Printf("   修改后: %+v\n", p2)

	// 5. 方法调用
	fmt.Println("\n5. 方法调用")
	c := Calculator{}
	vc := reflect.ValueOf(c)

	method := vc.MethodByName("Add")
	if method.IsValid() {
		args := []reflect.Value{
			reflect.ValueOf(5),
			reflect.ValueOf(3),
		}
		result := method.Call(args)
		fmt.Printf("   Add(5, 3) = %v\n", result[0].Int())
	}

	// 6. 动态创建
	fmt.Println("\n6. 动态创建")

	// 创建整数
	intType := reflect.TypeOf(0)
	intVal := reflect.New(intType).Elem()
	intVal.SetInt(42)
	fmt.Printf("   创建int: %v\n", intVal.Interface())

	// 创建切片
	sliceType := reflect.TypeOf([]int{})
	sliceVal := reflect.MakeSlice(sliceType, 0, 3)
	sliceVal = reflect.Append(sliceVal, reflect.ValueOf(1))
	sliceVal = reflect.Append(sliceVal, reflect.ValueOf(2))
	fmt.Printf("   创建slice: %v\n", sliceVal.Interface())

	// 创建Map
	mapType := reflect.TypeOf(map[string]int{})
	mapVal := reflect.MakeMap(mapType)
	mapVal.SetMapIndex(reflect.ValueOf("key"), reflect.ValueOf(42))
	fmt.Printf("   创建map: %v\n", mapVal.Interface())

	// 7. 实际应用：JSON序列化原理
	fmt.Println("\n7. 实际应用：JSON序列化原理")
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	user := User{Name: "Alice", Email: "alice@example.com"}
	jsonStr := simpleJSON(user)
	fmt.Printf("   JSON: %s\n", jsonStr)

	// 8. 性能和最佳实践
	fmt.Println("\n8. 性能和最佳实践")
	fmt.Println("   - 反射性能开销大，避免在热路径中使用")
	fmt.Println("   - 优先使用类型安全的代码")
	fmt.Println("   - 反射主要用于库和框架")
	fmt.Println("   - 缓存reflect.Value和reflect.Type对象")
}

// 简单的JSON序列化示例
func simpleJSON(v interface{}) string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Struct {
		result := "{"
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}
			if i > 0 {
				result += ", "
			}
			result += fmt.Sprintf(`"%s":"%v"`, tag, val.Field(i).Interface())
		}
		result += "}"
		return result
	}
	return ""
}
