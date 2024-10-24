package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	f1()
	fmt.Println("===========")
	wg.Add(1)
	go f2(&wg)
	wg.Wait()
}

func f1() {
	intSlice := []int{1, 2, 3, 4, 5}
	for index, value := range intSlice {
		intSlice = append(intSlice, value*10)
		fmt.Println(index, value)
	}
}

func f2(wg *sync.WaitGroup) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	for k, v := range m {
		m[k*10] = v * 10
		fmt.Println(k, v)
	}
	wg.Done()
}
