package main

import (
	"fmt"
	"sync"
	"time"
)

func Slice1(a []int) {
	a[0] = 1
	fmt.Printf("%v %p\n", a, &a)
}

func Qiepian(a map[int]int) {
	a[1] = 2
	fmt.Printf("%v %p\n", a, &a)
}
func Xicheng() {
	var k chan []int
	fmt.Printf("%v %T\n", k, k)
	//var p chan []int
	p := <-k
	fmt.Println(p)
}

//func main() {
//	//显然对于切片就是传递指针的副本
//	slice1 := make([]int, 0)
//	slice1 = append(slice1, 12)
//	fmt.Printf("%v %p\n", slice1, &slice1)
//	Slice1(slice1)
//	//fmt.Println(slice1)
//	//对于map 传递的是指针的副本
//	qiepian := make(map[int]int)
//	qiepian[1] = 1
//	fmt.Printf("%v %p\n", qiepian, &qiepian)
//	Qiepian(qiepian)
//	s := make(chan int, 2)
//	fmt.Println(s)
//	go Xicheng()
//}

func main() {
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	wg1 := sync.WaitGroup{}
	wg1.Add(2)
	var ch1 chan int
	var ch2 chan int
	// 向 ch1 中写入数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	// 向 ch2 中写入数据
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	// 监听 ch1 和 ch2 的读取操作
	for {
		select {
		case data := <-ch1:
			fmt.Println("read from ch1:", data)
		case data := <-ch2:
			fmt.Println("read from ch2:", data)
		}
	}
	wg1.Wait()
}
