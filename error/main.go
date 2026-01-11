// Package error 演示Go语言中的异常处理机制
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
)

// 1. Go中的错误处理概述
// Go没有传统意义上的异常机制（如try-catch），而是使用显式的错误返回
// 函数可以返回多个值，通常最后一个值是error类型

// 2. 基本错误处理
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零") // 创建一个简单的错误
	}
	return a / b, nil
}

func demonstrateBasicErrorHandling() {
	fmt.Println("=== 基本错误处理 ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}
}

// 3. 自定义错误类型
// 通过实现Error()方法来创建自定义错误类型

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("错误代码: %d, 消息: %s", e.Code, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return MyError{Code: 1001, Message: "年龄不能为负数"}
	}
	if age > 150 {
		return MyError{Code: 1002, Message: "年龄不能超过150"}
	}
	return nil
}

func demonstrateCustomError() {
	fmt.Println("\n=== 自定义错误类型 ===")

	ages := []int{25, -5, 200}

	for _, age := range ages {
		err := validateAge(age)
		if err != nil {
			// 类型断言检查错误类型
			var myErr MyError
			if errors.As(err, &myErr) {
				fmt.Printf("自定义错误: 代码=%d, 消息=%s\n", myErr.Code, myErr.Message)
			}
		} else {
			fmt.Printf("年龄 %d 验证通过\n", age)
		}
	}
}

// 4. 错误包装和错误链
// 使用fmt.Errorf和%w动词来包装错误，保留原始错误信息

func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取文件 %s 失败: %w", filename, err)
	}
	return data, nil
}

func processFile(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("处理文件时出错: %v", err)
	}

	// 处理文件内容...
	fmt.Printf("文件大小: %d 字节\n", len(data))
	return nil
}

func demonstrateErrorWrapping() {
	fmt.Println("\n=== 错误包装和错误链 ===")

	err := processFile("不存在的文件.txt")
	if err != nil {
		fmt.Printf("错误: %v\n", err)

		// 使用errors.Unwrap获取原始错误
		unwrapped := errors.Unwrap(err)
		fmt.Printf("解包后的错误: %v\n", unwrapped)

		// 使用errors.Is检查错误链中是否包含特定错误
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("错误原因是文件不存在")
		}

		// 使用errors.As检查错误链中是否包含特定类型的错误
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Printf("路径错误: 路径=%s, 操作=%s\n", pathError.Path, pathError.Op)
		}
	}
}

// 5. panic和recover
// panic用于处理不可恢复的错误，类似于其他语言的异常
// recover用于捕获panic，只能在defer函数中使用

func riskyOperation(shouldPanic bool) {
	fmt.Println("开始执行可能panic的操作")

	if shouldPanic {
		panic("发生了不可恢复的错误！")
	}

	fmt.Println("操作成功完成")
}

func demonstratePanicAndRecover() {
	fmt.Println("\n=== panic和recover ===")

	// 不使用recover的情况
	fmt.Println("1. 不使用recover的情况:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("捕获到panic: %v\n", r)
				fmt.Println("堆栈跟踪:")
				fmt.Println(string(debug.Stack()))
			}
		}()

		riskyOperation(true)
		fmt.Println("这行代码不会执行")
	}()

	// 使用recover的情况
	fmt.Println("\n2. 使用recover的情况:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("捕获到panic: %v\n", r)
			}
		}()

		riskyOperation(false)
		fmt.Println("操作正常完成")
	}()
}

// 6. defer语句
// defer用于延迟执行函数，通常用于资源清理和错误处理

func demonstrateDefer() {
	fmt.Println("\n=== defer语句 ===")

	fmt.Println("1. defer的执行顺序:")
	func() {
		defer fmt.Println("第一个defer")
		defer fmt.Println("第二个defer")
		defer fmt.Println("第三个defer")
		fmt.Println("函数体")
	}()

	fmt.Println("\n2. defer在资源清理中的应用:")
	func() {
		file, err := os.Create("test.txt")
		if err != nil {
			fmt.Printf("创建文件失败: %v\n", err)
			return
		}

		// 使用defer确保文件被关闭
		defer func() {
			fmt.Println("关闭文件")
			file.Close()
			// 删除测试文件
			err := os.Remove("test.txt")
			if err != nil {
				return
			}
			fmt.Printf("删除文件失败，%v\n", err)
		}()

		// 写入文件
		_, err = file.WriteString("测试内容")
		if err != nil {
			fmt.Printf("写入文件失败: %v\n", err)
			return
		}

		fmt.Println("文件操作成功")
	}()
}

// 7. 错误处理的最佳实践
func demonstrateBestPractices() {
	fmt.Println("\n=== 错误处理最佳实践 ===")

	// 1. 立即检查错误
	fmt.Println("1. 立即检查错误:")
	data, err := os.ReadFile("不存在的文件.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
	} else {
		fmt.Printf("文件内容: %s\n", string(data))
	}

	// 2. 提供有意义的错误信息
	fmt.Println("\n2. 提供有意义的错误信息:")
	_, err = os.Open("不存在的文件.txt")
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
	}

	// 3. 避免忽略错误
	fmt.Println("\n3. 避免忽略错误:")
	// 不好的做法:
	// os.Remove("file.txt") // 忽略错误

	// 好的做法:
	err = os.Remove("不存在的文件.txt")
	if err != nil {
		fmt.Printf("删除文件失败: %v\n", err)
	}

	// 4. 使用哨兵错误
	fmt.Println("\n4. 使用哨兵错误:")
	var ErrNotFound = errors.New("未找到")

	searchItem := func(id int) (string, error) {
		if id == 0 {
			return "", ErrNotFound
		}
		return fmt.Sprintf("项目%d", id), nil
	}

	item, err := searchItem(0)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("项目未找到")
	} else if err != nil {
		fmt.Printf("搜索出错: %v\n", err)
	} else {
		fmt.Printf("找到项目: %s\n", item)
	}
}

func main() {
	demonstrateBasicErrorHandling()
	demonstrateCustomError()
	demonstrateErrorWrapping()
	demonstratePanicAndRecover()
	demonstrateDefer()
	demonstrateBestPractices()
}
