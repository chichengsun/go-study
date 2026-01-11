package main

import (
	"fmt"
)

// ==================== 结构体基础知识 ====================

// Person 1. 基本结构体定义
// 结构体是将多个不同类型的数据组合在一起的复合数据类型
// 类似于其他语言中的类（class），但Go中的结构体没有继承
type Person struct {
	Name string // 字段名（大写开头表示导出，包外可访问）
	Age  int    // 字段类型
	Sex  string // 另一个字段
}

// 2. 匿名结构体
// 可以直接定义结构体而不给它命名，常用于临时数据结构
var anonymous struct {
	ID   int
	Name string
}

// Student 3. 结构体标签（Tag）
// 结构体字段可以附加标签，用于反射、JSON序列化等
type Student struct {
	Name    string `json:"name"`    // JSON序列化时字段名为name
	Grade   int    `json:"grade"`   // JSON序列化时字段名为grade
	Score   int    `json:"score"`   // JSON序列化时字段名为score
	Hobbies string `json:"hobbies"` // JSON序列化时字段名为hobbies
}

// Address 4. 嵌套结构体
// 结构体可以包含其他结构体作为字段
type Address struct {
	City    string
	Street  string
	ZipCode string
}

type Employee struct {
	Name    string
	Age     int
	Address Address // 嵌套结构体
}

// Human 5. 结构体组合（匿名字段）
// 使用匿名字段实现类似继承的效果（组合优于继承）
type Human struct {
	Name string
	Age  int
}

type Teacher struct {
	Human      // 匿名字段，直接嵌入
	Subject    string
	Experience int
}

// 6. 结构体指针
// 结构体指针的使用和操作

// ==================== 方法和函数 ====================

// Introduce 7. 结构体方法 - 值接收者
// 方法是与结构体关联的函数
// 值接收者：方法内部对结构体的修改不会影响原结构体
func (p Person) Introduce() string {
	return fmt.Sprintf("大家好，我叫%s，今年%d岁", p.Name, p.Age)
}

// Birthday 8. 结构体方法 - 指针接收者
// 指针接收者：方法内部可以修改原结构体
func (p *Person) Birthday() {
	p.Age++ // 年龄加1
}

// 9. 结构体方法 - 修改字段值
func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

// ==================== 主函数，演示所有知识点 ====================

// 自定义类型指的是使用type关键字定义的新类型，它可以是基本类型的别名，也可以是结构体、函数等组合而成的新类型。
// 自定义类型可以帮助我们更好地抽象和封装数据，让代码更加易读、易懂和易维护

// 其中类型别名和自定义类型有很大的区别type AliasCode = int
// 1. 不能绑定方法
// 2. 打印类型还是原始类型
// 3. 和原始类型比较，类型别名不用转换

func main() {
	fmt.Println("========== Go结构体详细教程 ==========\n")

	// 1. 结构体的创建和初始化
	fmt.Println("1. 结构体的创建和初始化:")

	// 方式1：按顺序初始化
	p1 := Person{"张三", 25, "男"}
	fmt.Printf("  按顺序初始化: %+v\n", p1)

	// 方式2：按字段名初始化（推荐）
	p2 := Person{Name: "李四", Age: 30, Sex: "女"}
	fmt.Printf("  按字段名初始化: %+v\n", p2)

	// 方式3：使用var声明（零值）
	var p3 Person
	p3.Name = "王五"
	p3.Age = 35
	p3.Sex = "男"
	fmt.Printf("  var声明后赋值: %+v\n", p3)

	// 方式4：使用new关键字
	p4 := new(Person)
	fmt.Printf("  new关键字创建: %+v\n", *p4)
	p4.Name = "赵六"
	p4.Age = 28
	p4.Sex = "男"
	fmt.Printf("  new关键字创建: %+v\n", *p4)

	// 方式5：匿名结构体
	anonymousPerson := struct {
		Name string
		Age  int
	}{
		Name: "临时人物",
		Age:  99,
	}
	fmt.Printf("  匿名结构体: %+v\n", anonymousPerson)
	fmt.Println()

	// 2. 结构体字段访问和修改
	fmt.Println("2. 结构体字段访问和修改:")
	fmt.Printf("  p1.Name = %s, p1.Age = %d\n", p1.Name, p1.Age)

	// 修改字段值
	p1.Name = "张三丰"
	p1.Age = 100
	fmt.Printf("  修改后: %+v\n", p1)

	// 结构体指针的字段访问
	p5 := &Person{Name: "钱七", Age: 40, Sex: "男"}
	fmt.Printf("  指针访问: Name=%s, Age=%d\n", p5.Name, p5.Age)
	// 等价于 (*p5).Name，但Go允许简写
	fmt.Println()

	// 3. 结构体作为函数参数
	fmt.Println("3. 结构体作为函数参数:")

	// 值传递（复制）
	passByValue := func(p Person) {
		p.Name = "修改后的名字"
		fmt.Printf("    函数内部: %+v\n", p)
	}

	testP := Person{Name: "原始名字", Age: 20, Sex: "男"}
	fmt.Printf("  调用前: %+v\n", testP)
	passByValue(testP)
	fmt.Printf("  调用后: %+v (值传递不会修改原结构体)\n", testP)

	// 引用传递（指针）
	passByReference := func(p *Person) {
		p.Name = "修改后的名字"
		fmt.Printf("    函数内部: %+v\n", *p)
	}

	fmt.Printf("  调用前: %+v\n", testP)
	passByReference(&testP)
	fmt.Printf("  调用后: %+v (指针传递会修改原结构体)\n", testP)
	fmt.Println()

	// 4. 结构体方法演示
	fmt.Println("4. 结构体方法演示:")

	methodDemo := Person{Name: "方法演示", Age: 25, Sex: "男"}

	// 调用值接收者方法
	intro := methodDemo.Introduce()
	fmt.Printf("  介绍: %s\n", intro)

	// 调用指针接收者方法（修改年龄）
	fmt.Printf("  原始年龄: %d\n", methodDemo.Age)
	methodDemo.Birthday()
	fmt.Printf("  生日后年龄: %d\n", methodDemo.Age)

	// 修改名字
	methodDemo.ChangeName("新名字")
	fmt.Printf("  修改名字后: %+v\n", methodDemo)
	fmt.Println()

	// 5. 嵌套结构体
	fmt.Println("5. 嵌套结构体:")

	emp := Employee{
		Name: "张员工",
		Age:  30,
		Address: Address{
			City:    "北京",
			Street:  "长安街1号",
			ZipCode: "100000",
		},
	}

	fmt.Printf("  员工信息: %+v\n", emp)
	fmt.Printf("  城市: %s\n", emp.Address.City)
	fmt.Printf("  详细地址: %s, %s\n", emp.Address.Street, emp.Address.City)
	fmt.Println()

	// 6. 结构体组合（匿名字段）
	fmt.Println("6. 结构体组合（匿名字段）:")

	teacher := Teacher{
		Human: Human{
			Name: "王老师",
			Age:  45,
		},
		Subject:    "数学",
		Experience: 20,
	}

	fmt.Printf("  老师信息: %+v\n", teacher)
	// 可以直接访问匿名字段的字段
	fmt.Printf("  直接访问: Name=%s, Age=%d\n", teacher.Name, teacher.Age)
	fmt.Printf("  自身字段: Subject=%s, Experience=%d\n", teacher.Subject, teacher.Experience)
	fmt.Println()

	// 7. 结构体切片
	fmt.Println("7. 结构体切片:")

	people := []Person{
		{Name: "小明", Age: 18, Sex: "男"},
		{Name: "小红", Age: 17, Sex: "女"},
		{Name: "小刚", Age: 19, Sex: "男"},
	}

	fmt.Println("  学生列表:")
	for i, p := range people {
		fmt.Printf("    %d. %s (%d岁, %s)\n", i+1, p.Name, p.Age, p.Sex)
	}
	fmt.Println()

	// 8. 结构体数组
	fmt.Println("8. 结构体数组:")

	var studentArray [3]Student
	studentArray[0] = Student{Name: "张三", Grade: 3, Score: 85, Hobbies: "篮球"}
	studentArray[1] = Student{Name: "李四", Grade: 2, Score: 92, Hobbies: "画画"}
	studentArray[2] = Student{Name: "王五", Grade: 1, Score: 78, Hobbies: "音乐"}

	fmt.Println("  学生数组:")
	for i, s := range studentArray {
		fmt.Printf("    %d. %s - 年级:%d, 分数:%d, 爱好:%s\n",
			i+1, s.Name, s.Grade, s.Score, s.Hobbies)
	}
	fmt.Println()

	// 9. 结构体比较
	fmt.Println("9. 结构体比较:")

	p6 := Person{Name: "测试", Age: 20, Sex: "男"}
	p7 := Person{Name: "测试", Age: 20, Sex: "男"}
	p8 := Person{Name: "测试", Age: 21, Sex: "男"}

	fmt.Printf("  p6 == p7: %v (相同内容)\n", p6 == p7)
	fmt.Printf("  p6 == p8: %v (不同年龄)\n", p6 == p8)
	fmt.Println()

	// 10. 结构体复制
	fmt.Println("10. 结构体复制:")

	original := Person{Name: "原版", Age: 30, Sex: "男"}
	copy := original // 值复制

	fmt.Printf("  原版: %+v\n", original)
	fmt.Printf("  复制: %+v\n", copy)

	original.Name = "修改原版"
	fmt.Printf("  修改原版后 - 原版: %+v\n", original)
	fmt.Printf("  修改原版后 - 复制: %+v (不受影响)\n", copy)
	fmt.Println()

	// 11. 空结构体
	fmt.Println("11. 空结构体:")

	empty := struct{}{}
	fmt.Printf("  空结构体: %+v (占用0字节内存)\n", empty)
	fmt.Printf("  空结构体大小: %d bytes\n", 0) // 空结构体不占用内存空间
	fmt.Println()

	// 12. 结构体标签的实际应用
	fmt.Println("12. 结构体标签的实际应用:")

	s1 := Student{
		Name:    "张学生",
		Grade:   5,
		Score:   95,
		Hobbies: "编程,阅读",
	}

	fmt.Printf("  结构体: %+v\n", s1)
	fmt.Printf("  标签说明: json标签用于序列化时的字段名映射\n")
	fmt.Printf("  例如: Name字段的json标签是'name'，序列化后会变成{\"name\":\"张学生\",...}\n")
	fmt.Println()

	fmt.Println("========== 结构体教程结束 ==========")
}

// ==================== 额外的练习函数 ====================

// 练习1：创建一个表示"书籍"的结构体，包含标题、作者、价格、页数
type Book struct {
	Title  string
	Author string
	Price  float64
	Pages  int
}

// 练习2：创建一个表示"矩形"的结构体，包含长和宽，以及计算面积和周长的方法
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 练习3：创建一个表示"汽车"的结构体，包含品牌、型号、年份，以及一个表示"驾驶"的方法
type Car struct {
	Brand string
	Model string
	Year  int
}

func (c *Car) Drive(kilometers int) {
	fmt.Printf("驾驶 %s %s 行驶了 %d 公里\n", c.Brand, c.Model, kilometers)
}

// 练习函数 - 可以在main中调用来练习
func practiceExercises() {
	fmt.Println("\n========== 结构体练习 ==========")

	// 练习1：书籍
	book := Book{
		Title:  "Go语言编程",
		Author: "张三",
		Price:  68.00,
		Pages:  350,
	}
	fmt.Printf("书籍: %+v\n", book)

	// 练习2：矩形
	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", rect.Area(), rect.Perimeter())

	// 练习3：汽车
	car := Car{Brand: "特斯拉", Model: "Model 3", Year: 2023}
	car.Drive(100)

	fmt.Println("================================")
}
