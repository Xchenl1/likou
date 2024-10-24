package main

import (
	"fmt"
	"sync"
)

//type A11 int
//type B11 = int

func main() {
	var wg sync.WaitGroup
	sl := []int{1, 2, 3, 4}
	for _, v := range sl {
		wg.Add(1)
		go print(&v, &wg)
	}
	wg.Wait()
	//var i int
	//var l A11 = A11(i) //go语言是强类型语言 可以强制转化
	//var l1 B11 = i
}

func print(s *int, wg *sync.WaitGroup) {
	fmt.Println(*s)
	wg.Done()
}
