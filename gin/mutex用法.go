package main

import (
	"fmt"
	"sync"
)

var suo1 sync.Mutex
var i int = 0

func Add() {
	suo1.Lock()
	defer suo1.Unlock()
	fmt.Println(i)
	i++
	k.Done()
}

func Add1() {
	suo1.Lock()
	defer suo1.Unlock()
	fmt.Println(i)
	i++
	k.Done()
}

var k sync.WaitGroup

func main() {
	k.Add(2)
	go Add()
	go Add1()
	k.Wait()
}
