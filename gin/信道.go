package main

import "fmt"

var ch chan int = make(chan int, 5)

func insertNum() {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	//	//使用信道
	//	ch <- "chen"
	//	ch <- "asd"
	//	fmt.Println(<-ch)
	//	fmt.Println(<-ch)
	//	close(ch)\
	go insertNum()
	i1, v := <-ch
	fmt.Println(i1, v)
	i2, v1 := <-ch
	fmt.Println(i2, v1)
	i3, v2 := <-ch
	fmt.Println(i3, v2)
	fmt.Println(<-ch)
}
