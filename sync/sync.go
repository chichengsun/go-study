// Package sync 演示Go语言中并发及并发安全相关的内容
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 1. Goroutine - Go的轻量级线程
// Goroutine是Go语言并发设计的核心，它比线程更轻量级，可以轻松创建成千上万个

func demonstrateGoroutine() {
	fmt.Println("=== Goroutine 演示 ===")

	// 使用go关键字启动一个goroutine
	go func() {
		fmt.Println("这是来自goroutine的消息")
	}()

	// 主goroutine继续执行
	fmt.Println("这是来自主goroutine的消息")

	// 等待一下让goroutine有时间执行完
	time.Sleep(100 * time.Millisecond)
}

// 2. WaitGroup - 等待一组goroutine完成
// WaitGroup用于等待一组goroutine执行完成

func demonstrateWaitGroup() {
	fmt.Println("\n=== WaitGroup 演示 ===")

	var wg sync.WaitGroup

	// 启动5个goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1) // 增加计数器

		go func(id int) {
			defer wg.Done() // 完成时减少计数器

			fmt.Printf("Goroutine %d 开始工作\n", id)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Goroutine %d 完成工作\n", id)
		}(i)
	}

	fmt.Println("等待所有goroutine完成...")
	wg.Wait() // 阻塞直到计数器为0
	fmt.Println("所有goroutine已完成")
}

// 3. Mutex - 互斥锁
// Mutex用于保护共享资源，确保同一时间只有一个goroutine可以访问

func demonstrateMutex() {
	fmt.Println("\n=== Mutex 演示 ===")

	var counter int
	var mutex sync.Mutex

	// 启动多个goroutine同时修改counter
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 使用mutex保护对counter的访问
			mutex.Lock()
			defer mutex.Unlock()

			// 临界区开始
			currentValue := counter
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			counter = currentValue + 1
			// 临界区结束
		}()
	}

	wg.Wait()
	fmt.Printf("最终counter值: %d\n", counter)
}

// 4. RWMutex - 读写互斥锁
// RWMutex允许多个读操作同时进行，但写操作是独占的

func demonstrateRWMutex() {
	fmt.Println("\n=== RWMutex 演示 ===")

	var data string
	var rwMutex sync.RWMutex

	// 读操作goroutine
	for i := 0; i < 3; i++ {
		go func(id int) {
			rwMutex.RLock() // 获取读锁
			defer rwMutex.RUnlock()

			fmt.Printf("读操作 %d: 读取数据 '%s'\n", id, data)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("读操作 %d: 完成读取\n", id)
		}(i)
	}

	// 写操作goroutine
	for i := 0; i < 2; i++ {
		go func(id int) {
			rwMutex.Lock() // 获取写锁
			defer rwMutex.Unlock()

			fmt.Printf("写操作 %d: 开始写入\n", id)
			data = fmt.Sprintf("数据-%d", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("写操作 %d: 完成写入\n", id)
		}(i)
	}

	// 等待所有操作完成
	time.Sleep(1 * time.Second)
}

// 5. Once - 确保操作只执行一次
// Once用于确保某个操作只执行一次，即使在多个goroutine中调用

func demonstrateOnce() {
	fmt.Println("\n=== Once 演示 ===")

	var once sync.Once
	var initialized bool

	// 启动多个goroutine尝试初始化
	for i := 0; i < 5; i++ {
		go func(id int) {
			once.Do(func() {
				fmt.Printf("初始化操作由goroutine %d 执行\n", id)
				initialized = true
			})

			fmt.Printf("Goroutine %d 检查初始化状态: %v\n", id, initialized)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

// 6. Cond - 条件变量
// Cond用于goroutine之间的同步，允许goroutine等待或通知某个条件

func demonstrateCond() {
	fmt.Println("\n=== Cond 演示 ===")

	var mutex sync.Mutex
	var cond = sync.NewCond(&mutex)
	var ready bool

	// 等待条件的goroutine
	go func() {
		mutex.Lock()
		defer mutex.Unlock()

		fmt.Println("等待条件满足...")
		for !ready {
			cond.Wait() // 等待条件满足
		}
		fmt.Println("条件已满足，继续执行")
	}()

	// 设置条件的goroutine
	go func() {
		time.Sleep(500 * time.Millisecond)

		mutex.Lock()
		defer mutex.Unlock()

		fmt.Println("设置条件为满足")
		ready = true
		cond.Broadcast() // 通知所有等待的goroutine
	}()

	time.Sleep(1 * time.Second)
}

// 7. Pool - 对象池
// Pool用于缓存和复用临时对象，减少GC压力

func demonstratePool() {
	fmt.Println("\n=== Pool 演示 ===")

	// 创建一个对象池
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("创建新对象")
			return make([]byte, 1024)
		},
	}

	// 从池中获取对象
	obj1 := pool.Get().([]byte)
	fmt.Printf("获取对象1: 长度=%d\n", len(obj1))

	// 使用对象
	obj1[0] = 1

	// 将对象放回池中
	pool.Put(obj1)

	// 再次从池中获取对象（可能会复用之前的对象）
	obj2 := pool.Get().([]byte)
	fmt.Printf("获取对象2: 长度=%d, 第一个字节=%d\n", len(obj2), obj2[0])

	// 将对象放回池中
	pool.Put(obj2)
}

// 8. Atomic - 原子操作
// Atomic提供原子操作，无需加锁即可保证并发安全

func demonstrateAtomic() {
	fmt.Println("\n=== Atomic 演示 ===")

	var counter int64
	var wg sync.WaitGroup

	// 启动多个goroutine使用原子操作增加counter
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 原子地增加counter
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("最终counter值: %d\n", atomic.LoadInt64(&counter))

	// 比较并交换操作
	var old int64 = 10
	var new int64 = 20

	// 如果counter的值等于old，则将其设置为new
	swapped := atomic.CompareAndSwapInt64(&counter, old, new)
	fmt.Printf("CAS操作: 期望值=%d, 新值=%d, 是否交换=%v, 当前值=%d\n",
		old, new, swapped, atomic.LoadInt64(&counter))
}

// 9. Channel - 通道
// Channel是goroutine之间通信的主要方式，遵循"不要通过共享内存来通信，而要通过通信来共享内存"的理念

func demonstrateChannel() {
	fmt.Println("\n=== Channel 演示 ===")

	// 创建一个无缓冲通道
	unbuffered := make(chan int)

	// 发送数据的goroutine
	go func() {
		fmt.Println("准备发送数据到无缓冲通道")
		unbuffered <- 42 // 阻塞直到有接收者
		fmt.Println("数据已发送")
	}()

	// 接收数据的goroutine
	go func() {
		time.Sleep(100 * time.Millisecond) // 确保发送者先准备
		fmt.Println("准备从无缓冲通道接收数据")
		value := <-unbuffered // 阻塞直到有数据
		fmt.Printf("接收到数据: %d\n", value)
	}()

	// 创建一个有缓冲通道
	buffered := make(chan string, 3)

	// 发送数据到有缓冲通道
	buffered <- "消息1"
	buffered <- "消息2"
	buffered <- "消息3"
	fmt.Println("已发送3条消息到有缓冲通道")

	// 接收数据
	go func() {
		time.Sleep(200 * time.Millisecond)
		for i := 0; i < 3; i++ {
			msg := <-buffered
			fmt.Printf("从有缓冲通道接收: %s\n", msg)
		}
	}()

	time.Sleep(500 * time.Millisecond)
}

// 10. Select - 多路复用
// Select允许同时等待多个通道操作

func demonstrateSelect() {
	fmt.Println("\n=== Select 演示 ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 发送数据的goroutine
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自通道1的消息"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自通道2的消息"
	}()

	// 使用select等待多个通道
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("接收到: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("接收到: %s\n", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("超时，没有接收到消息")
		}
	}
}

// 11. Context - 上下文
// Context用于控制goroutine的生命周期，传递取消信号和截止时间

func demonstrateContext() {
	fmt.Println("\n=== Context 演示 ===")

	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个goroutine，监听context的取消信号
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine收到取消信号: %v\n", ctx.Err())
		case <-time.After(2 * time.Second):
			fmt.Println("Goroutine正常完成")
		}
	}()

	// 1秒后取消context
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("取消context")
		cancel()
	}()

	time.Sleep(1500 * time.Millisecond)

	// 使用带超时的context
	ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("带超时的context: %v\n", ctx.Err())
		}
	}()

	time.Sleep(1 * time.Second)
}

// 12. 并发模式 - Worker Pool
// Worker Pool是一种常见的并发模式，用于限制并发goroutine的数量

func demonstrateWorkerPool() {
	fmt.Println("\n=== Worker Pool 演示 ===")

	const numWorkers = 3
	const numJobs = 10

	// 创建任务通道和结果通道
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动worker
	for w := 0; w < numWorkers; w++ {
		go func(id int) {
			for j := range jobs {
				fmt.Printf("Worker %d 开始处理任务 %d\n", id, j)
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				results <- j * 2 // 假设处理结果是输入的两倍
				fmt.Printf("Worker %d 完成任务 %d\n", id, j)
			}
		}(w)
	}

	// 发送任务
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for r := 0; r < numJobs; r++ {
		result := <-results
		fmt.Printf("收到结果: %d\n", result)
	}
}

// 13. 并发安全的数据结构
// 演示如何创建并发安全的数据结构

type ConcurrentCounter struct {
	mu    sync.RWMutex
	value int
}

func (c *ConcurrentCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *ConcurrentCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func demonstrateConcurrentDataStructure() {
	fmt.Println("\n=== 并发安全的数据结构演示 ===")

	counter := &ConcurrentCounter{}
	var wg sync.WaitGroup

	// 启动多个goroutine并发修改计数器
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter.Value())
}

func main() {

	demonstrateGoroutine()
	demonstrateWaitGroup()
	demonstrateMutex()
	demonstrateRWMutex()
	demonstrateOnce()
	demonstrateCond()
	demonstratePool()
	demonstrateAtomic()
	demonstrateChannel()
	demonstrateSelect()
	demonstrateContext()
	demonstrateWorkerPool()
	demonstrateConcurrentDataStructure()

	var counter int
	var wait sync.WaitGroup
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for j := 0; j < 1000; j++ {
				counter++ // 这里有竞争
			}
		}()
	}
	wait.Wait() // 等待所有goroutine完成
	fmt.Printf("counter = %d\n", counter)

	fmt.Println("\n=== 所有并发演示完成 ===")
}
