# Go反射快速参考卡

## 基础操作

```go
import "reflect"

var x int = 42
v := reflect.ValueOf(x)  // 获取Value
t := reflect.TypeOf(x)   // 获取Type
```

## 修改值（必须用指针）

```go
num := 10
v := reflect.ValueOf(&num).Elem()  // 关键：指针 + Elem()
v.SetInt(100)  // num 现在是 100
```

## 结构体操作

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "Alice", Age: 30}
v := reflect.ValueOf(p)

// 获取字段数量
v.NumField()  // 2

// 遍历字段
for i := 0; i < v.NumField(); i++ {
    field := v.Field(i)
    fieldType := v.Type().Field(i)
    fmt.Printf("%s = %v\n", fieldType.Name, field.Interface())
}

// 通过字段名访问
nameField := v.FieldByName("Name")

// 修改字段（需要指针）
p2 := Person{Name: "Bob", Age: 25}
v2 := reflect.ValueOf(&p2).Elem()
v2.FieldByName("Name").SetString("Charlie")
```

## 方法调用

```go
type Calculator struct{}
func (c Calculator) Add(a, b int) int { return a + b }

c := Calculator{}
v := reflect.ValueOf(c)
method := v.MethodByName("Add")

args := []reflect.Value{
    reflect.ValueOf(5),
    reflect.ValueOf(3),
}
results := method.Call(args)
fmt.Println(results[0].Int())  // 8
```

## 动态创建

```go
// 创建整数
intVal := reflect.New(reflect.TypeOf(0)).Elem()
intVal.SetInt(42)

// 创建切片
sliceType := reflect.TypeOf([]int{})
sliceVal := reflect.MakeSlice(sliceType, 0, 10)
sliceVal = reflect.Append(sliceVal, reflect.ValueOf(1))

// 创建Map
mapType := reflect.TypeOf(map[string]int{})
mapVal := reflect.MakeMap(mapType)
mapVal.SetMapIndex(reflect.ValueOf("key"), reflect.ValueOf(42))

// 创建Channel
chanType := reflect.TypeOf(make(chan int))
chanVal := reflect.MakeChan(chanType, 0)
```

## Kind类型

```go
// Kind是比Type更基础的类型分类
v.Kind()  // 返回 reflect.Kind 类型

// 常见Kind：
// reflect.Int, reflect.Int32, reflect.Int64
// reflect.String, reflect.Bool, reflect.Float64
// reflect.Slice, reflect.Map, reflect.Struct
// reflect.Ptr, reflect.Chan, reflect.Func
```

## 标签解析

```go
type User struct {
    Name string `json:"name" db:"name_column"`
}

t := reflect.TypeOf(User{})
field := t.Field(0)
jsonTag := field.Tag.Get("json")  // "name"
dbTag := field.Tag.Get("db")      // "name_column"
```

## 类型判断

```go
var i interface{} = "hello"
v := reflect.ValueOf(i)

switch v.Kind() {
case reflect.String:
    fmt.Println("是字符串")
case reflect.Int:
    fmt.Println("是整数")
case reflect.Slice:
    fmt.Println("是切片")
}
```

## 接口实现检查

```go
type MyWriter struct{}
func (w MyWriter) Write(p []byte) (n int, err error) { return 0, nil }

writerType := reflect.TypeOf((*interface{ Write([]byte) (int, error) })(nil)).Elem()
myWriterType := reflect.TypeOf(MyWriter{})
fmt.Println(myWriterType.Implements(writerType))  // true
```

## 性能提示

```go
// ❌ 避免：在循环中重复创建reflect.Value
for i := 0; i < 1000; i++ {
    v := reflect.ValueOf(data[i])  // 每次都创建
}

// ✅ 推荐：缓存reflect.Type
var cachedType reflect.Type
func init() {
    cachedType = reflect.TypeOf(MyStruct{})
}

// ✅ 推荐：优先使用接口
func process(data interface{}) {
    // 反射是最后的选择
}
```

## 常见错误

```go
// ❌ 错误：无法修改值
v := reflect.ValueOf(10)
v.SetInt(100)  // panic: can't set value

// ✅ 正确：使用指针
v := reflect.ValueOf(&x).Elem()
v.SetInt(100)

// ❌ 错误：访问未导出字段
type S struct {
    private string
}
v := reflect.ValueOf(S{})
v.FieldByName("private").SetString("test")  // panic

// ✅ 正确：只访问导出字段
type S struct {
    Public string
}
v := reflect.ValueOf(S{})
v.FieldByName("Public").SetString("test")
```

## 与Java对比

| Go反射 | Java反射 | 说明 |
|--------|----------|------|
| `reflect.Value` | `Object` | 值容器 |
| `reflect.Type` | `Class` | 类型信息 |
| `v.MethodByName()` | `clazz.getMethod()` | 方法查找 |
| `v.FieldByName()` | `clazz.getField()` | 字段查找 |
| `v.Call()` | `method.invoke()` | 方法调用 |
| `v.Set()` | `field.set()` | 设置值 |

## 使用场景

1. **JSON/XML序列化** - 标准库`encoding/json`
2. **数据库ORM** - 如GORM
3. **依赖注入** - 框架中的组件创建
4. **单元测试** - Mock对象创建
5. **命令行参数解析** - 如Cobra库

## 记住的原则

1. **反射是最后的手段** - 优先使用类型安全的代码
2. **性能开销大** - 避免在热路径中使用
3. **只能访问导出成员** - 大写开头的字段/方法
4. **需要指针才能修改** - `reflect.ValueOf(&x).Elem()`
5. **缓存reflect对象** - 避免重复创建