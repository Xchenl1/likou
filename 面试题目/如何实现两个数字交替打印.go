package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // 设置等待组的计数器为2

	ch := make(chan int) // 创建一个用于通信的整数通道

	// 协程1：打印奇数
	go func() {
		defer wg.Done() // 减少等待组的计数器
		for {
			fmt.Println(1)
			ch <- 0 // 发送信号给协程2，表示可以打印偶数
			<-ch    // 等待协程2发送信号
		}
	}()
	// 协程2：打印偶数
	go func() {
		defer wg.Done() // 减少等待组的计数器
		for {
			<-ch // 等待协程1发送信号
			fmt.Println(2)
			ch <- 0 // 发送信号给协程1，表示可以打印奇数
		}
	}()

	wg.Wait() // 等待两个协程结束

}
