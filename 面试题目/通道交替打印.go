package main

import (
	"fmt"
	"sync"
)

var c1 chan int
var c2 chan int

//var mutex sync.Mutex

func Print1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i += 2 {
		<-c1
		fmt.Println(i)
		c2 <- 1
	}
}

func Print2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 26; i += 2 {
		<-c2
		fmt.Println("print2:", i)
		c1 <- 1
	}
}

func main() {
	c1, c2 = make(chan int), make(chan int)
	//c1 = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go Print1(&wg)
	go Print2(&wg)
	c1 <- 1
	wg.Wait()
}
