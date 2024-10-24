package main

import "fmt"

func main() {
	//	var ch chan int = make(chan int, 2)
	//	ch <- 2
	//	ch <- 3
	//exit:
	//	for {
	//		select {
	//		case v := <-ch:
	//			fmt.Println(v)
	//		default:
	//			fmt.Println("没有数据")
	//			break exit
	//		}
	//	}
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//ch1 <- 1
	//
	//// ...
	//for {
	//	select {
	//	case v := <-ch1:
	//		// 处理v
	//		fmt.Println(v)
	//		ch1 = nil // 屏蔽ch1分支
	//	case v := <-ch2:
	//		// 处理v
	//		fmt.Println(v)
	//		ch2 = nil // 屏蔽ch2分支
	//	}
	//}
	//	var ch chan int // 未初始化的channel
	//	ch = make(chan int, 1)
	//	go func() {
	//		ch <- 1 // 发送操作将永久阻塞
	//	}()
	//info:
	//	for {
	//		select {
	//		case v, ok := <-ch: //通道关闭并且通道数据为空
	//			if !ok {
	//				break info
	//			}
	//			fmt.Println(v)
	//			close(ch)
	//		}
	//	}
	//	//<-ch // 接收操作将永久阻塞
	ch := make(chan float64, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- float64(i)
		}
		close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	fmt.Println(<-ch)
}
