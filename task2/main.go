package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	num := 5
	res := incrementValue(&num)
	fmt.Println(res)

	nums := []int{1, 2, 3, 4, 5}
	resarr := doubleSliceElements(&nums)
	fmt.Println(resarr)

	PrintlnOddEven()

	tasks := []func(){
		func() { fmt.Println("Task 1 executed") },
		func() { fmt.Println("Task 2 executed") },
		func() { fmt.Println("Task 3 executed") },
	}
	Scheduler(tasks...)

	ShapeExample()

	emp := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
	}
	emp.PrintInfo()

	sendMsg()

	// 演示带有缓冲的通道
	fmt.Println("=== 带有缓冲的通道演示 ===")
	bufferedChannelDemo()

	counterDemo()

	fmt.Println("=== 原子操作演示 ===")
	atomicCounterDemo()
}

/*
*
指针-题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func incrementValue(ptr *int) int {
	*ptr += 10

	return *ptr
}

/*
*
指针-题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func doubleSliceElements(slicePtr *[]int) []int {
	for i := 0; i < len(*slicePtr); i++ {
		(*slicePtr)[i] *= 2
	}
	return *slicePtr
}

/*
*
goroutine-题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

func PrintlnOddEven() {
	var wait = sync.WaitGroup{}
	wait.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("奇数: ", i)
			}
		}
		wait.Done()
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数: ", i)
			}
		}
		wait.Done()
	}()
	wait.Wait()
	fmt.Println("Odd and Even numbers printed successfully.")
}

/*
goroutine-题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func Scheduler(tasks ...func()) {
	var wait = sync.WaitGroup{}
	wait.Add(len(tasks))
	for _, task := range tasks {
		go func() {
			start := time.Now()
			task()
			time.Sleep(2 * time.Second) // 模拟任务执行时间
			duration := time.Since(start)
			fmt.Printf("Task executed in %v\n", duration)
			wait.Done()
		}()
	}
	wait.Wait()
}

/*
面向对象
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 表示一个矩形结构体，包含宽度和高度两个属性。
// Width 表示矩形的宽度。
// Height 表示矩形的高度。
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 表示一个圆形结构体，包含圆形的半径属性。
// Radius 是圆形的半径，类型为 float64。
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: Width=%.2f, Height=%.2f, Area=%.2f, Perimeter=%.2f", r.Width, r.Height, r.Area(), r.Perimeter())
}
func (c Circle) String() string {
	return fmt.Sprintf("Circle: Radius=%.2f, Area=%.2f, Perimeter=%.2f", c.Radius, c.Area(), c.Perimeter())
}

func ShapeExample() {
	var rect Shape
	var circ Shape
	rect = Rectangle{Width: 5, Height: 10}
	circ = Circle{Radius: 7}
	// fmt.Println(rect)
	// fmt.Println(circ)
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", circ.Area(), circ.Perimeter())
}

/*
面向对象
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     // 组合 Person 结构体
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee ID: %s, Name: %s, Age: %d\n", e.EmployeeID, e.Name, e.Age)
}

/*
*
Channel
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*
*/
func sendMsg() {
	var c = make(chan int, 10) //  初始化一个 有一个缓冲位的通道

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			num, ok := <-c // 从通道中接收数据
			if !ok {
				break // 如果通道已关闭，退出循环
			}
			fmt.Println("Received:", num)
		}
		close(c) // 关闭通道
	}()
	time.Sleep(2 * time.Second)
}

/*
*
Channel
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func bufferedChannelDemo() {
	// 创建一个带有缓冲的通道，缓冲大小为20
	ch := make(chan int, 20)

	// 创建一个WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 一个生产者和一个消费者

	// 生产者协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("生产者: 发送 ", i)
		}
		close(ch) // 生产完毕后关闭通道
		fmt.Println("生产者: 已发送所有数据并关闭通道")
	}()

	// 消费者协程
	go func() {
		defer wg.Done()
		// 使用range循环从通道接收数据，直到通道关闭
		for num := range ch {
			fmt.Println("消费者: 接收 ", num)
			time.Sleep(10 * time.Millisecond) // 模拟处理时间
		}
		fmt.Println("消费者: 已接收所有数据")
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("bufferedChannelDemo 完成")
}

/**
锁机制
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

var counter int64
var mu sync.Mutex

func incrementCounter() {

	for i := 0; i < 1000; i++ {

		mu.Lock() // 锁定互斥锁，保护共享资源
		counter++
		mu.Unlock() // 解锁互斥锁
	}
}

func counterDemo() {
	var wg sync.WaitGroup
	wg.Add(10) // 启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			incrementCounter()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

/**
锁机制
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

var counters int64

func atomicIncrement() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counters, 1) // 使用原子操作递增计数器
	}
}
func atomicCounterDemo() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			atomicIncrement()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counters:", atomic.LoadInt64(&counters)) // 使用 atomic.LoadInt64 获取计数器的值
}
