// Package reflect_demo 详细演示Go语言中reflect包的使用
// 通过丰富的代码示例展示反射的核心概念和应用场景
package main

import (
	"fmt"
	"reflect"
)

// Person 用于演示结构体反射
type Person struct {
	Name string `json:"name" db:"name_column"`
	Age  int    `json:"age" db:"age_column"`
}

// Student 继承Person，演示嵌套结构体
type Student struct {
	Person
	StudentID string `json:"student_id"`
	Score     float64
}

// MethodDemo 用于演示方法反射
type MethodDemo struct {
	Value int
}

// Add 方法
func (m MethodDemo) Add(a, b int) int {
	return a + b + m.Value
}

// Print 方法
func (m *MethodDemo) Print(s string) {
	fmt.Printf("Value: %d, String: %s\n", m.Value, s)
}

func main() {
	fmt.Println("=== Go Reflect 深度解析 ===\n")

	// 1. 基础：Type和Value
	fmt.Println("1. 基础：Type和Value")
	basicTypeAndValue()

	// 2. Kind的详细分类
	fmt.Println("\n2. Kind的详细分类")
	kindDemo()

	// 3. 基本类型反射
	fmt.Println("\n3. 基本类型反射")
	basicTypeReflection()

	// 4. 结构体反射
	fmt.Println("\n4. 结构体反射")
	structReflection()

	// 5. 结构体标签解析
	fmt.Println("\n5. 结构体标签解析")
	structTagDemo()

	// 6. 方法反射
	fmt.Println("\n6. 方法反射")
	methodReflection()

	// 7. 动态创建和修改
	fmt.Println("\n7. 动态创建和修改")
	dynamicCreationAndModification()

	// 8. 切片和Map操作
	fmt.Println("\n8. 切片和Map操作")
	sliceAndMapOperations()

	// 9. 接口和类型断言
	fmt.Println("\n9. 接口和类型断言")
	interfaceAndTypeAssertion()

	// 10. 实际应用场景
	fmt.Println("\n10. 实际应用场景")
	practicalApplications()
}

// 1. 基础：Type和Value
func basicTypeAndValue() {
	// Type：描述类型信息，类似于Java的Class
	// Value：包含具体值和类型信息

	var x int = 42
	v := reflect.ValueOf(x) // 获取Value
	t := reflect.TypeOf(x)  // 获取Type

	fmt.Printf("值: %v\n", v)
	fmt.Printf("类型: %v\n", t)
	fmt.Printf("值的Kind: %v\n", v.Kind())
	fmt.Printf("类型的Kind: %v\n", t.Kind())

	// Value和Type的关系
	fmt.Printf("v.Type() == t: %v\n", v.Type() == t)

	// 通过Value获取Type的几种方式
	fmt.Printf("v.Type(): %v\n", v.Type())
	fmt.Printf("reflect.TypeOf(v.Interface()): %v\n", reflect.TypeOf(v.Interface()))
}

// 2. Kind的详细分类
func kindDemo() {
	// Kind表示底层类型分类，比Type更基础

	values := []interface{}{
		42,                     // Int
		3.14,                   // Float64
		"hello",                // String
		true,                   // Bool
		[]int{1, 2},            // Slice
		map[string]int{"a": 1}, // Map
		struct{}{},             // Struct
		(*int)(nil),            // Ptr
		[3]int{1, 2, 3},        // Array
		make(chan int),         // Chan
		func() {},              // Func
	}

	for _, val := range values {
		v := reflect.ValueOf(val)
		fmt.Printf("%-20v -> Kind: %-10v, Type: %v\n",
			val, v.Kind(), v.Type())
	}

	// Kind和Type的区别
	var myInt int = 42
	var myInt32 int32 = 42

	v1 := reflect.ValueOf(myInt)
	v2 := reflect.ValueOf(myInt32)

	fmt.Printf("\nKind比较: myInt(%v) vs myInt32(%v)\n", v1.Kind(), v2.Kind())
	fmt.Printf("Type比较: %v vs %v\n", v1.Type(), v2.Type())
}

// 3. 基本类型反射
func basicTypeReflection() {
	// 3.1 整数类型
	var i int = 42
	v := reflect.ValueOf(&i).Elem() // 使用指针来修改

	fmt.Printf("原始值: %d\n", i)
	fmt.Printf("反射值: %d\n", v.Int())

	v.SetInt(100) // 修改值
	fmt.Printf("修改后: %d\n", i)

	// 3.2 字符串
	var s string = "hello"
	vs := reflect.ValueOf(&s).Elem()

	fmt.Printf("原始字符串: %s\n", s)
	fmt.Printf("反射字符串: %s\n", vs.String())

	vs.SetString("world")
	fmt.Printf("修改后: %s\n", s)

	// 3.3 浮点数
	var f float64 = 3.14
	vf := reflect.ValueOf(&f).Elem()

	fmt.Printf("原始浮点: %f\n", f)
	vf.SetFloat(6.28)
	fmt.Printf("修改后: %f\n", f)

	// 3.4 布尔值
	var b bool = true
	vb := reflect.ValueOf(&b).Elem()

	fmt.Printf("原始布尔: %t\n", b)
	vb.SetBool(false)
	fmt.Printf("修改后: %t\n", b)
}

// 4. 结构体反射
func structReflection() {
	p := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(p)

	fmt.Printf("结构体: %+v\n", p)
	fmt.Printf("字段数量: %d\n", v.NumField())

	// 遍历所有字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		fmt.Printf("  字段%d: %s = %v (类型: %v)\n",
			i, fieldType.Name, field.Interface(), field.Type())
	}

	// 通过字段名访问
	nameField := v.FieldByName("Name")
	if nameField.IsValid() {
		fmt.Printf("Name字段值: %v\n", nameField.String())
	}

	// 修改结构体字段（需要指针）
	p2 := Person{Name: "Bob", Age: 25}
	v2 := reflect.ValueOf(&p2).Elem()

	nameField2 := v2.FieldByName("Name")
	if nameField2.CanSet() {
		nameField2.SetString("Charlie")
	}

	ageField := v2.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(35)
	}

	fmt.Printf("修改后的结构体: %+v\n", p2)

	// 嵌套结构体
	s := Student{
		Person:    Person{Name: "David", Age: 20},
		StudentID: "S001",
		Score:     95.5,
	}

	vs := reflect.ValueOf(s)
	fmt.Printf("\n嵌套结构体: %+v\n", s)

	// 访问嵌套字段
	personField := vs.FieldByName("Person")
	if personField.IsValid() {
		nameField := personField.FieldByName("Name")
		fmt.Printf("嵌套字段Person.Name: %v\n", nameField.String())
	}

	// 访问直接字段
	studentIDField := vs.FieldByName("StudentID")
	fmt.Printf("直接字段StudentID: %v\n", studentIDField.String())
}

// 5. 结构体标签解析
func structTagDemo() {
	p := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(p)
	t := v.Type()

	fmt.Printf("结构体: %+v\n", p)

	// 遍历字段和标签
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag

		fmt.Printf("\n字段: %s\n", field.Name)
		fmt.Printf("  JSON标签: %s\n", tag.Get("json"))
		fmt.Printf("  DB标签: %s\n", tag.Get("db"))

		// 解析标签的完整信息
		jsonTag, ok := tag.Lookup("json")
		if ok {
			fmt.Printf("  JSON标签存在: %s\n", jsonTag)
		}

		// 标签包含多个值的情况
		// 例如: `json:"name,omitempty" db:"name_column"`
		// 可以使用 strings.Split(tag.Get("json"), ",") 来处理
	}

	// 实际应用：根据标签进行字段映射
	fmt.Println("\n实际应用：根据标签映射字段")
	mapFieldsByTag(p)
}

// 根据标签映射字段的示例函数
func mapFieldsByTag(p interface{}) {
	v := reflect.ValueOf(p)
	t := v.Type()

	fieldMap := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag != "" {
			fieldMap[jsonTag] = v.Field(i).Interface()
		}
	}

	fmt.Printf("映射结果: %+v\n", fieldMap)
}

// 6. 方法反射
func methodReflection() {
	m := MethodDemo{Value: 10}
	v := reflect.ValueOf(m)

	fmt.Printf("对象: %+v\n", m)
	fmt.Printf("方法数量: %d\n", v.NumMethod())

	// 遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Type().Method(i)
		fmt.Printf("方法%d: %s, 参数: %v, 返回值: %v\n",
			i, method.Name, method.Type.In(0), method.Type.Out(0))
	}

	// 调用方法
	addMethod := v.MethodByName("Add")
	if addMethod.IsValid() {
		// 准备参数
		args := []reflect.Value{
			reflect.ValueOf(5),
			reflect.ValueOf(3),
		}

		// 调用方法
		results := addMethod.Call(args)
		fmt.Printf("调用Add(5, 3)结果: %v\n", results[0].Int())
	}

	// 调用指针接收者的方法
	vm := reflect.ValueOf(&m)
	printMethod := vm.MethodByName("Print")
	if printMethod.IsValid() {
		args := []reflect.Value{reflect.ValueOf("反射调用")}
		printMethod.Call(args)
	}

	// 动态调用未知方法
	fmt.Println("\n动态调用示例:")
	dynamicCallMethod(m, "Add", 20, 30)
	dynamicCallMethod(&m, "Print", "Hello from dynamic call")
}

func dynamicCallMethod(obj interface{}, methodName string, args ...interface{}) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		fmt.Printf("方法 %s 不存在\n", methodName)
		return
	}

	// 转换参数
	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	// 调用
	results := method.Call(reflectArgs)

	// 处理返回值
	if len(results) > 0 {
		fmt.Printf("方法 %s 调用结果: ", methodName)
		for i, result := range results {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%v", result.Interface())
		}
		fmt.Println()
	} else {
		fmt.Printf("方法 %s 调用完成\n", methodName)
	}
}

// 7. 动态创建和修改
func dynamicCreationAndModification() {
	// 7.1 创建基本类型
	intType := reflect.TypeOf(0)
	intVal := reflect.New(intType).Elem()
	intVal.SetInt(42)
	fmt.Printf("动态创建int: %v\n", intVal.Interface())

	// 7.2 创建结构体
	personType := reflect.TypeOf(Person{})
	personVal := reflect.New(personType).Elem()

	// 设置字段
	personVal.FieldByName("Name").SetString("Dynamic")
	personVal.FieldByName("Age").SetInt(28)

	fmt.Printf("动态创建结构体: %+v\n", personVal.Interface())

	// 7.3 创建切片
	sliceType := reflect.TypeOf([]int{})
	sliceVal := reflect.MakeSlice(sliceType, 0, 5)

	// 添加元素
	sliceVal = reflect.Append(sliceVal, reflect.ValueOf(1))
	sliceVal = reflect.Append(sliceVal, reflect.ValueOf(2))
	sliceVal = reflect.Append(sliceVal, reflect.ValueOf(3))

	fmt.Printf("动态创建切片: %v\n", sliceVal.Interface())

	// 7.4 创建Map
	mapType := reflect.TypeOf(map[string]int{})
	mapVal := reflect.MakeMap(mapType)

	// 添加键值对
	mapVal.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(1))
	mapVal.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(2))

	fmt.Printf("动态创建Map: %v\n", mapVal.Interface())

	// 7.5 创建Channel
	chanType := reflect.TypeOf(make(chan int))
	chanVal := reflect.MakeChan(chanType, 0)

	fmt.Printf("动态创建Channel: %v\n", chanVal.Type())

	// 7.6 创建函数
	funcType := reflect.TypeOf(func(a, b int) int { return a + b })
	funcVal := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		a := int(args[0].Int())
		b := int(args[1].Int())
		// 必须创建正确类型的Value
		result := a + b
		return []reflect.Value{reflect.ValueOf(result)}
	})

	result := funcVal.Call([]reflect.Value{
		reflect.ValueOf(int(10)),
		reflect.ValueOf(int(20)),
	})
	fmt.Printf("动态创建函数调用: %v\n", result[0].Int())
}

// 8. 切片和Map操作
func sliceAndMapOperations() {
	// 8.1 切片操作
	slice := []int{1, 2, 3}
	v := reflect.ValueOf(&slice).Elem()

	fmt.Printf("原始切片: %v\n", slice)
	fmt.Printf("长度: %d, 容量: %d\n", v.Len(), v.Cap())

	// 修改元素
	v.Index(0).SetInt(100)
	fmt.Printf("修改后: %v\n", slice)

	// 添加元素（使用append需要特殊处理）
	newVal := reflect.Append(v, reflect.ValueOf(4))
	v.Set(newVal)
	fmt.Printf("添加元素后: %v\n", slice)

	// 切片操作
	subSlice := v.Slice(1, 3)
	fmt.Printf("切片[1:3]: %v\n", subSlice.Interface())

	// 8.2 Map操作
	m := map[string]int{"a": 1, "b": 2}
	vm := reflect.ValueOf(&m).Elem()

	fmt.Printf("原始Map: %v\n", m)

	// 获取值
	val := vm.MapIndex(reflect.ValueOf("a"))
	fmt.Printf("Map['a']: %v\n", val.Int())

	// 设置值
	vm.SetMapIndex(reflect.ValueOf("c"), reflect.ValueOf(3))
	fmt.Printf("添加后: %v\n", m)

	// 删除值
	vm.SetMapIndex(reflect.ValueOf("b"), reflect.Value{}) // 设置为空值删除
	fmt.Printf("删除后: %v\n", m)

	// 遍历Map
	fmt.Println("遍历Map:")
	iter := vm.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("  %v: %v\n", k.Interface(), v.Interface())
	}
}

// 9. 接口和类型断言
func interfaceAndTypeAssertion() {
	// 9.1 接口的动态类型
	var i interface{} = "hello"
	v := reflect.ValueOf(i)

	fmt.Printf("接口值: %v\n", v.Interface())
	fmt.Printf("动态类型: %v\n", v.Type())
	fmt.Printf("Kind: %v\n", v.Kind())

	// 9.2 类型断言
	i = 42
	v = reflect.ValueOf(i)

	// 使用Kind进行类型判断
	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("是整数: %d\n", v.Int())
	case reflect.String:
		fmt.Printf("是字符串: %s\n", v.String())
	case reflect.Slice:
		fmt.Printf("是切片: %v\n", v.Interface())
	}

	// 9.3 Type Switch模拟
	typeSwitch := func(val interface{}) {
		v := reflect.ValueOf(val)
		fmt.Printf("值: %v, 类型: %v, Kind: %v\n", val, v.Type(), v.Kind())
	}

	typeSwitch(42)
	typeSwitch("hello")
	typeSwitch([]int{1, 2, 3})
	typeSwitch(map[string]int{"a": 1})

	// 9.4 检查类型实现接口
	fmt.Printf("MyWriter是否实现了Writer接口: %v\n", checkInterfaceImplementation())
}

// 检查接口实现的辅助函数
func checkInterfaceImplementation() bool {
	type MyWriter struct{}

	writerType := reflect.TypeOf((*interface{ Write([]byte) (int, error) })(nil)).Elem()
	myWriterType := reflect.TypeOf(MyWriter{})

	return myWriterType.Implements(writerType)
}

// 10. 实际应用场景
func practicalApplications() {
	fmt.Println("10.1 简单的JSON序列化器")
	simpleJSONDemo()

	fmt.Println("\n10.2 通用的表单验证器")
	formValidationDemo()

	fmt.Println("\n10.3 依赖注入模拟")
	dependencyInjectionDemo()
}

// 简单的JSON序列化器
func simpleJSONDemo() {
	p := Person{Name: "Alice", Age: 30}
	jsonStr := toJSON(p)
	fmt.Printf("序列化结果: %s\n", jsonStr)
}

func toJSON(v interface{}) string {
	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Struct:
		result := "{"
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldVal := val.Field(i)

			if i > 0 {
				result += ", "
			}

			// 使用json标签
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}

			result += fmt.Sprintf(`"%s":%v`, tag, fieldVal.Interface())
		}
		result += "}"
		return result
	case reflect.Slice:
		result := "["
		for i := 0; i < val.Len(); i++ {
			if i > 0 {
				result += ", "
			}
			result += fmt.Sprintf("%v", val.Index(i).Interface())
		}
		result += "]"
		return result
	default:
		return fmt.Sprintf("%v", v)
	}
}

// 通用的表单验证器
func formValidationDemo() {
	type UserForm struct {
		Username string `validate:"required,min=3,max=20"`
		Email    string `validate:"required,email"`
		Age      int    `validate:"min=18,max=100"`
	}

	form := UserForm{
		Username: "Al",
		Email:    "invalid-email",
		Age:      15,
	}

	errors := validateForm(form)
	fmt.Printf("验证结果: %v\n", errors)
}

func validateForm(v interface{}) []string {
	val := reflect.ValueOf(v)
	t := val.Type()

	var errors []string

	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)
		tag := field.Tag.Get("validate")

		if tag == "" {
			continue
		}

		// 简单的验证规则解析
		// 实际项目中可以使用更复杂的解析器
		if tag == "required" {
			if fieldVal.IsZero() {
				errors = append(errors, fmt.Sprintf("%s是必填字段", field.Name))
			}
		}

		// 根据类型进行验证
		switch fieldVal.Kind() {
		case reflect.String:
			str := fieldVal.String()
			if len(str) < 3 {
				errors = append(errors, fmt.Sprintf("%s长度不能小于3", field.Name))
			}
		case reflect.Int:
			age := fieldVal.Int()
			if age < 18 {
				errors = append(errors, fmt.Sprintf("%s必须大于等于18", field.Name))
			}
		}
	}

	if len(errors) == 0 {
		errors = append(errors, "验证通过")
	}

	return errors
}

// 依赖注入模拟
func dependencyInjectionDemo() {
	// 模拟一个简单的依赖注入容器
	container := make(map[reflect.Type]interface{})

	// 注册服务
	type Database struct{}
	container[reflect.TypeOf(Database{})] = &Database{}

	type Service struct {
		DB *Database // 使用导出字段
	}

	// 创建Service并注入依赖
	serviceType := reflect.TypeOf(Service{})
	serviceVal := reflect.New(serviceType).Elem()

	// 查找依赖
	dbType := reflect.TypeOf((*Database)(nil)).Elem()
	dbVal := reflect.ValueOf(container[dbType])

	// 注入
	serviceVal.FieldByName("DB").Set(dbVal)

	fmt.Printf("依赖注入结果: %+v\n", serviceVal.Interface())
}
