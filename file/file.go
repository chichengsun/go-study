// Package file 演示Go语言中的文件操作
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// 1. 基本文件操作
func demonstrateBasicFileOperations() {
	fmt.Println("=== 基本文件操作 ===")

	// 创建文件
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close() // 确保文件被关闭

	// 写入文件
	content := "Hello, Go文件操作!\n这是第二行内容。"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}

	fmt.Println("文件创建并写入成功")

	// 读取文件
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	fmt.Printf("文件内容: %s\n", string(data))

	// 删除文件
	err = os.Remove("test.txt")
	if err != nil {
		fmt.Printf("删除文件失败: %v\n", err)
		return
	}

	fmt.Println("文件删除成功")
}

// 2. 文件信息获取
func demonstrateFileInfo() {
	fmt.Println("\n=== 文件信息获取 ===")

	// 创建测试文件
	file, err := os.Create("info_test.txt")
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	err = file.Close()
	if err != nil {
		return
	}

	// 获取文件信息
	info, err := os.Stat("info_test.txt")
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fmt.Printf("文件名: %s\n", info.Name())
	fmt.Printf("文件大小: %d 字节\n", info.Size())
	fmt.Printf("文件权限: %v\n", info.Mode())
	fmt.Printf("最后修改时间: %v\n", info.ModTime())
	fmt.Printf("是否是目录: %v\n", info.IsDir())
	fmt.Printf("系统接口类型: %T\n", info.Sys())

	// 清理
	err = os.Remove("info_test.txt")
	if err != nil {
		return
	}
}

// 3. 文件权限操作
func demonstrateFilePermissions() {
	fmt.Println("\n=== 文件权限操作 ===")

	// 创建具有特定权限的文件
	file, err := os.OpenFile("perm_test.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	file.Close()

	// 获取文件权限
	info, err := os.Stat("perm_test.txt")
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fmt.Printf("文件权限: %v\n", info.Mode().Perm())

	// 修改文件权限
	err = os.Chmod("perm_test.txt", 0755)
	if err != nil {
		fmt.Printf("修改文件权限失败: %v\n", err)
		return
	}

	// 再次获取文件权限
	info, _ = os.Stat("perm_test.txt")
	fmt.Printf("修改后的文件权限: %v\n", info.Mode().Perm())

	// 清理
	os.Remove("perm_test.txt")
}

// 4. 目录操作
func demonstrateDirectoryOperations() {
	fmt.Println("\n=== 目录操作 ===")

	// 创建目录
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}

	// 创建多级目录
	err = os.MkdirAll("testdir/subdir1/subdir2", 0755)
	if err != nil {
		fmt.Printf("创建多级目录失败: %v\n", err)
		return
	}

	// 在目录中创建文件
	file, err := os.Create("testdir/file.txt")
	if err != nil {
		fmt.Printf("在目录中创建文件失败: %v\n", err)
		return
	}
	file.WriteString("目录中的文件")
	file.Close()

	// 读取目录内容
	entries, err := os.ReadDir("testdir")
	if err != nil {
		fmt.Printf("读取目录内容失败: %v\n", err)
		return
	}

	fmt.Println("目录内容:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  目录: %s\n", entry.Name())
		} else {
			fmt.Printf("  文件: %s\n", entry.Name())
		}
	}

	// 删除目录及其内容
	err = os.RemoveAll("testdir")
	if err != nil {
		fmt.Printf("删除目录失败: %v\n", err)
		return
	}

	fmt.Println("目录删除成功")
}

// 5. 文件路径操作
func demonstratePathOperations() {
	fmt.Println("\n=== 文件路径操作 ===")

	// 构建路径
	path := filepath.Join("dir", "subdir", "file.txt")
	fmt.Printf("构建的路径: %s\n", path)

	// 获取路径的目录部分
	dir := filepath.Dir(path)
	fmt.Printf("目录部分: %s\n", dir)

	// 获取路径的文件名部分
	filename := filepath.Base(path)
	fmt.Printf("文件名部分: %s\n", filename)

	// 获取文件扩展名
	ext := filepath.Ext(path)
	fmt.Printf("文件扩展名: %s\n", ext)

	// 分割路径的目录和文件名
	dir, file := filepath.Split(path)
	fmt.Printf("分割结果 - 目录: %s, 文件: %s\n", dir, file)

	// 获取绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("获取绝对路径失败: %v\n", err)
		return
	}
	fmt.Printf("绝对路径: %s\n", absPath)

	// 清理路径
	cleanPath := filepath.Clean("dir/../dir/./subdir/file.txt")
	fmt.Printf("清理后的路径: %s\n", cleanPath)

	// 匹配模式
	matches, err := filepath.Glob("*.go")
	if err != nil {
		fmt.Printf("匹配模式失败: %v\n", err)
		return
	}

	fmt.Println("匹配的Go文件:")
	for _, match := range matches {
		fmt.Printf("  %s\n", match)
	}
}

// 6. 文件复制和移动
func demonstrateFileCopyAndMove() {
	fmt.Println("\n=== 文件复制和移动 ===")

	// 创建源文件
	sourceFile, err := os.Create("source.txt")
	if err != nil {
		fmt.Printf("创建源文件失败: %v\n", err)
		return
	}
	sourceFile.WriteString("这是源文件的内容")
	sourceFile.Close()

	// 方法1: 使用io.Copy复制文件
	source, err := os.Open("source.txt")
	if err != nil {
		fmt.Printf("打开源文件失败: %v\n", err)
		return
	}
	defer source.Close()

	dest, err := os.Create("dest.txt")
	if err != nil {
		fmt.Printf("创建目标文件失败: %v\n", err)
		return
	}
	defer dest.Close()

	_, err = io.Copy(dest, source)
	if err != nil {
		fmt.Printf("复制文件失败: %v\n", err)
		return
	}

	fmt.Println("文件复制成功 (使用io.Copy)")

	// 方法2: 使用ReadFile和WriteFile复制文件
	data, err := os.ReadFile("source.txt")
	if err != nil {
		fmt.Printf("读取源文件失败: %v\n", err)
		return
	}

	err = os.WriteFile("dest2.txt", data, 0644)
	if err != nil {
		fmt.Printf("写入目标文件失败: %v\n", err)
		return
	}

	fmt.Println("文件复制成功 (使用ReadFile/WriteFile)")

	// 移动文件（重命名）
	err = os.Rename("dest2.txt", "moved.txt")
	if err != nil {
		fmt.Printf("移动文件失败: %v\n", err)
		return
	}

	fmt.Println("文件移动成功")

	// 清理
	os.Remove("source.txt")
	os.Remove("dest.txt")
	os.Remove("moved.txt")
}

// 7. 文件读取方式
func demonstrateFileReadingMethods() {
	fmt.Println("\n=== 文件读取方式 ===")

	// 创建测试文件
	content := `第一行内容
第二行内容
第三行内容
第四行内容`

	err := os.WriteFile("read_test.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("创建测试文件失败: %v\n", err)
		return
	}

	// 方法1: 一次性读取整个文件
	data, err := os.ReadFile("read_test.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}
	fmt.Printf("1. 一次性读取:\n%s\n", string(data))

	// 方法2: 使用bufio逐行读取
	file, err := os.Open("read_test.txt")
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("2. 使用bufio逐行读取:")
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  行%d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("扫描错误: %v\n", err)
	}

	// 方法3: 使用bufio按字节读取
	file, _ = os.Open("read_test.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Println("3. 使用bufio按字节读取:")
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Printf("读取错误: %v\n", err)
			break
		}
		if n == 0 {
			break
		}
		fmt.Printf("  读取到: %q\n", buffer[:n])
	}

	// 清理
	os.Remove("read_test.txt")
}

// 8. 文件写入方式
func demonstrateFileWritingMethods() {
	fmt.Println("\n=== 文件写入方式 ===")

	// 方法1: 使用WriteFile一次性写入
	content := "这是使用WriteFile写入的内容"
	err := os.WriteFile("write_test1.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}
	fmt.Println("1. 使用WriteFile写入成功")

	// 方法2: 使用File.WriteString
	file, err := os.Create("write_test2.txt")
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("这是使用File.WriteString写入的内容\n")
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}
	fmt.Println("2. 使用File.WriteString写入成功")

	// 方法3: 使用bufio.Writer
	file, err = os.Create("write_test3.txt")
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("这是使用bufio.Writer写入的内容\n")
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}

	// 确保所有缓冲数据都写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Printf("刷新缓冲区失败: %v\n", err)
		return
	}
	fmt.Println("3. 使用bufio.Writer写入成功")

	// 清理
	os.Remove("write_test1.txt")
	os.Remove("write_test2.txt")
	os.Remove("write_test3.txt")
}

// 9. 文件追加内容
func demonstrateFileAppend() {
	fmt.Println("\n=== 文件追加内容 ===")

	// 创建初始文件
	err := os.WriteFile("append_test.txt", []byte("初始内容\n"), 0644)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}

	// 方法1: 使用os.OpenFile以追加模式打开
	file, err := os.OpenFile("append_test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("追加的内容1\n")
	if err != nil {
		fmt.Printf("追加内容失败: %v\n", err)
		return
	}

	// 方法2: 使用bufio.Writer以追加模式打开
	file, err = os.OpenFile("append_test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("追加的内容2\n")
	if err != nil {
		fmt.Printf("追加内容失败: %v\n", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("刷新缓冲区失败: %v\n", err)
		return
	}

	// 读取并显示最终内容
	data, err := os.ReadFile("append_test.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	fmt.Printf("最终文件内容:\n%s", string(data))

	// 清理
	os.Remove("append_test.txt")
}

// 10. 临时文件操作
func demonstrateTempFileOperations() {
	fmt.Println("\n=== 临时文件操作 ===")

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Printf("创建临时文件失败: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // 确保临时文件被删除
	defer tempFile.Close()

	fmt.Printf("临时文件路径: %s\n", tempFile.Name())

	// 写入临时文件
	_, err = tempFile.WriteString("这是临时文件的内容")
	if err != nil {
		fmt.Printf("写入临时文件失败: %v\n", err)
		return
	}

	// 读取临时文件
	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Printf("读取临时文件失败: %v\n", err)
		return
	}

	fmt.Printf("临时文件内容: %s\n", string(data))

	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "example-dir-*")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // 确保临时目录被删除

	fmt.Printf("临时目录路径: %s\n", tempDir)

	// 在临时目录中创建文件
	tempFileInDir, err := os.CreateTemp(tempDir, "file-*.txt")
	if err != nil {
		fmt.Printf("在临时目录中创建文件失败: %v\n", err)
		return
	}
	defer tempFileInDir.Close()

	fmt.Printf("临时目录中的文件路径: %s\n", tempFileInDir.Name())
}

// 11. 文件搜索和遍历
func demonstrateFileSearchAndTraversal() {
	fmt.Println("\n=== 文件搜索和遍历 ===")

	// 创建测试目录结构
	err := os.MkdirAll("search_test/dir1/subdir1", 0755)
	if err != nil {
		fmt.Printf("创建测试目录结构失败: %v\n", err)
		return
	}

	err = os.MkdirAll("search_test/dir2", 0755)
	if err != nil {
		fmt.Printf("创建测试目录结构失败: %v\n", err)
		return
	}

	// 创建测试文件
	os.WriteFile("search_test/file1.txt", []byte("内容1"), 0644)
	os.WriteFile("search_test/file2.go", []byte("内容2"), 0644)
	os.WriteFile("search_test/dir1/file3.txt", []byte("内容3"), 0644)
	os.WriteFile("search_test/dir1/file4.go", []byte("内容4"), 0644)
	os.WriteFile("search_test/dir2/file5.txt", []byte("内容5"), 0644)

	// 方法1: 使用filepath.Walk遍历目录
	fmt.Println("1. 使用filepath.Walk遍历目录:")
	err = filepath.Walk("search_test", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Printf("  目录: %s\n", path)
		} else {
			fmt.Printf("  文件: %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}

	// 方法2: 使用filepath.WalkDir遍历目录（更高效）
	fmt.Println("\n2. 使用filepath.WalkDir遍历目录:")
	err = filepath.WalkDir("search_test", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Printf("  目录: %s\n", path)
		} else {
			fmt.Printf("  文件: %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}

	// 方法3: 查找特定类型的文件
	fmt.Println("\n3. 查找Go文件:")
	var goFiles []string

	err = filepath.Walk("search_test", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("查找Go文件失败: %v\n", err)
	} else {
		fmt.Println("找到的Go文件:")
		for _, file := range goFiles {
			fmt.Printf("  %s\n", file)
		}
	}

	// 清理
	os.RemoveAll("search_test")
}

// 12. 文件监控
func demonstrateFileWatcher() {
	fmt.Println("\n=== 文件监控 ===")

	// 创建测试目录
	err := os.Mkdir("watch_test", 0755)
	if err != nil {
		fmt.Printf("创建测试目录失败: %v\n", err)
		return
	}
	defer os.RemoveAll("watch_test")

	// 使用fsnotify包监控文件变化（需要先安装：go get github.com/fsnotify/fsnotify）
	// 这里只演示基本概念，实际使用需要导入fsnotify包
	fmt.Println("文件监控需要使用fsnotify包:")
	fmt.Println("import \"github.com/fsnotify/fsnotify\"")
	fmt.Println(`
// 创建监控器
watcher, err := fsnotify.NewWatcher()
if err != nil {
    log.Fatal(err)
}
defer watcher.Close()

// 添加要监控的目录
err = watcher.Add("watch_test")
if err != nil {
    log.Fatal(err)
}

// 启动goroutine处理事件
go func() {
    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                return
            }
            fmt.Println("event:", event)
            if event.Op&fsnotify.Write == fsnotify.Write {
                fmt.Println("modified file:", event.Name)
            }
        case err, ok := <-watcher.Errors:
            if !ok {
                return
            }
            fmt.Println("error:", err)
        }
    }
}()

// 创建文件触发事件
time.Sleep(100 * time.Millisecond)
ioutil.WriteFile("watch_test/test.txt", []byte("hello"), 0644)
time.Sleep(100 * time.Millisecond)
ioutil.WriteFile("watch_test/test.txt", []byte("world"), 0644)
time.Sleep(100 * time.Millisecond)
`)
}

// 13. 文件压缩和解压
func demonstrateFileCompression() {
	fmt.Println("\n=== 文件压缩和解压 ===")

	// 创建测试文件
	content := "这是需要压缩的内容"
	err := os.WriteFile("compress_test.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("创建测试文件失败: %v\n", err)
		return
	}
	defer os.Remove("compress_test.txt")

	// 使用compress/gzip包进行压缩（需要导入）
	fmt.Println("文件压缩需要使用compress/gzip或compress/zip包:")
	fmt.Println("import \"compress/gzip\"")
	fmt.Println("import \"os\"")
	fmt.Println(`
// 压缩文件
func compressFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()
    
    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()
    
    writer := gzip.NewWriter(dstFile)
    defer writer.Close()
    
    _, err = io.Copy(writer, srcFile)
    return err
}

// 解压文件
func decompressFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()
    
    reader, err := gzip.NewReader(srcFile)
    if err != nil {
        return err
    }
    defer reader.Close()
    
    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()
    
    _, err = io.Copy(dstFile, reader)
    return err
}
`)
}

// 14. JSON文件操作
func demonstrateJSONFileOperations() {
	fmt.Println("\n=== JSON文件操作 ===")

	// 定义结构体
	type Person struct {
		Name    string    `json:"name"`
		Age     int       `json:"age"`
		Email   string    `json:"email"`
		Created time.Time `json:"created"`
	}

	// 创建实例
	person := Person{
		Name:    "张三",
		Age:     30,
		Email:   "zhangsan@example.com",
		Created: time.Now(),
	}

	// 序列化为JSON并写入文件
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("序列化JSON失败: %v\n", err)
		return
	}

	err = os.WriteFile("person.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("写入JSON文件失败: %v\n", err)
		return
	}

	fmt.Println("JSON文件写入成功")

	// 从文件读取JSON并反序列化
	data, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Printf("读取JSON文件失败: %v\n", err)
		return
	}

	var loadedPerson Person
	err = json.Unmarshal(data, &loadedPerson)
	if err != nil {
		fmt.Printf("反序列化JSON失败: %v\n", err)
		return
	}

	fmt.Printf("从JSON文件加载的数据:\n")
	fmt.Printf("  姓名: %s\n", loadedPerson.Name)
	fmt.Printf("  年龄: %d\n", loadedPerson.Age)
	fmt.Printf("  邮箱: %s\n", loadedPerson.Email)
	fmt.Printf("  创建时间: %s\n", loadedPerson.Created.Format("2006-01-02 15:04:05"))

	// 清理
	os.Remove("person.json")
}

// 15. 文件锁操作
func demonstrateFileLocking() {
	fmt.Println("\n=== 文件锁操作 ===")

	// 文件锁需要使用syscall包（平台相关）
	fmt.Println("文件锁需要使用syscall包:")
	fmt.Println("import \"syscall\"")
	fmt.Println(`
// 获取文件锁
func lockFile(file *os.File) error {
    return syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
}

// 释放文件锁
func unlockFile(file *os.File) error {
    return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}

// 使用示例
file, err := os.Create("lock_test.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// 获取排他锁
err = lockFile(file)
if err != nil {
    log.Fatal(err)
}
defer unlockFile(file)

// 在锁保护下进行文件操作
file.WriteString("这是在锁保护下写入的内容")
`)
}

func main() {
	demonstrateBasicFileOperations()
	demonstrateFileInfo()
	demonstrateFilePermissions()
	demonstrateDirectoryOperations()
	demonstratePathOperations()
	demonstrateFileCopyAndMove()
	demonstrateFileReadingMethods()
	demonstrateFileWritingMethods()
	demonstrateFileAppend()
	demonstrateTempFileOperations()
	demonstrateFileSearchAndTraversal()
	demonstrateFileWatcher()
	demonstrateFileCompression()
	demonstrateJSONFileOperations()
	demonstrateFileLocking()
}
