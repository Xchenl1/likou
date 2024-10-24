package main

import (
	"fmt"
	"sync"
	"time"
)

// SlidingWindowRateLimiter
// 定义了 SlidingWindowRateLimiter 结构体，表示滑动窗口算法的限流器。
// 它包含了请求限制数量、窗口大小、时间间隔、每个时间间隔内的请求数量、上次请求时间以及一个互斥锁用于并发安全
// 滑动窗口算法的限流器结构体
type SlidingWindowRateLimiter struct {
	rateLimit   int           // 请求限制数量
	windowSize  int           // 窗口大小
	interval    time.Duration // 时间间隔
	requests    []int         // 每个时间间隔内的请求数量
	lastRequest time.Time     // 上次请求时间
	lock        sync.Mutex    // 互斥锁
}

// NewSlidingWindowRateLimiter 是一个构造函数，用于创建一个新的滑动窗口限流器。
// 它接收请求限制数量、窗口大小和时间间隔作为参数，并返回一个指向 SlidingWindowRateLimiter 结构体的指针。
// 在构造函数中，我们初始化了限流器的各个字段，并使用 make 创建了一个窗口大小为 windowSize 的切片用于存储每个时间间隔内的请求数量。
// 初始化限流器
func NewSlidingWindowRateLimiter(rateLimit int, windowSize int, interval time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		rateLimit:   rateLimit,
		windowSize:  windowSize,
		interval:    interval,
		requests:    make([]int, windowSize),
		lastRequest: time.Now(),
		lock:        sync.Mutex{},
	}
}

// HandleRequest 方法用于处理请求。在该方法中，我们首先获取互斥锁，以确保并发安全。
// 处理请求
func (limiter *SlidingWindowRateLimiter) HandleRequest() bool {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()

	//检查是否需要滑动窗口
	//我们使用 time.Since 函数计算当前时间与上次请求时间的时间间隔，如果超过了设定的时间间隔，
	//则调用 slideWindow 方法进行窗口滑动
	if time.Since(limiter.lastRequest) > limiter.interval {
		limiter.slideWindow()
	}

	// 检查请求数量是否超过限制
	// 接着，我们计算当前窗口内的请求数量，并判断是否超过了请求限制数量。
	// 如果超过了限制，则返回 false 表示请求被限流，否则将当前请求计数器加一，并返回 true 表示请求被处理。
	sum := 0
	for _, count := range limiter.requests {
		sum += count
	}
	fmt.Println("sum:", sum)
	if sum >= limiter.rateLimit {
		return false
	}

	// 处理请求并增加计数器
	limiter.requests[limiter.windowSize-1]++
	fmt.Println(limiter.requests)
	return true
}

// slideWindow 方法用于进行窗口滑动操作。在该方法中，我们首先计算当前时间与上次请求时间的时间间隔对应的窗口位置 currentWindow。
// 如果 currentWindow 大于等于窗口大小，则表示需要重置所有时间间隔内的请求数量为 0。我们通过循环将 requests 切片中的所有元素设置为 0。
// 否则，我们使用 copy 函数将请求数量从当前窗口位置开始复制到切片的开头，并将剩余的位置设置为 0，以保证窗口大小始终为固定大小。
// 滑动窗口
func (limiter *SlidingWindowRateLimiter) slideWindow() {
	currentWindow := int(time.Since(limiter.lastRequest) / limiter.interval)
	fmt.Println("currentWindow:", currentWindow, limiter.windowSize)
	if currentWindow >= limiter.windowSize {
		// 重置所有时间间隔的请求数量为0
		for i := 0; i < limiter.windowSize; i++ {
			limiter.requests[i] = 0
		}
	} else {
		// 向前滑动窗口，并丢弃最旧的时间间隔内的请求数量
		copy(limiter.requests, limiter.requests[currentWindow:])
		fmt.Println(limiter.requests)
		for i := limiter.windowSize - currentWindow; i < limiter.windowSize; i++ {
			limiter.requests[i] = 0
		}
		fmt.Println(limiter.requests)
	}
	limiter.lastRequest = time.Now()
}

func main() {
	//创建一个每秒最多处理3个请求的限流器，窗口大小为5，时间间隔为200毫秒
	limiter := NewSlidingWindowRateLimiter(3, 5, time.Millisecond*200)

	// 模拟处理10个请求a
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * 200) // 模拟请求间隔200毫秒
		if limiter.HandleRequest() {
			fmt.Println("处理请求", i+1)
		} else {
			fmt.Println("限流请求", i+1)
		}
	}
	//a := []int{1, 2, 3}
	//copy(a, a[1:])
	//fmt.Println(a)
}
