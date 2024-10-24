package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	//for循环可以读取ch但是不close会发生死锁错误
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//不合法 报错
	//fmt.Println(<-ch)
}
