package main

import (
	"fmt"
	"sync"
	"time"
)

// 阻塞主协程
var wg = sync.WaitGroup{}

func SendCode() {
	fmt.Println("发送验证码开始")
	time.Sleep(3 * time.Second)
	fmt.Println("发送验证码开始")
	wg.Done() //没有done会一直阻塞
}
func main() {
	wg.Add(1)
	//
	for {
		fmt.Println("用户校验完成")
		go SendCode() //用来阻塞主线程
		fmt.Println("校验码已发送，请注意查收。。")
	}
	//waitgroup
	wg.Wait()
}
