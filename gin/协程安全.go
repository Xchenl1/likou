package main

import (
	"fmt"
	"sync"
)

var sum int
var wg1 = sync.WaitGroup{}
var suo = sync.Mutex{}

func addsum() {
	//上一把锁
	suo.Lock()
	defer suo.Unlock()
	for i := 0; i < 100; i++ {
		sum++
		fmt.Println(sum)
	}
	wg1.Done()
}
func main() {
	wg1.Add(2)
	go addsum()
	go addsum()
	wg1.Wait()
	fmt.Println(sum)
}
