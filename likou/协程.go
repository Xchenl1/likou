package main

import "time"

func main() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("panic")
	}()
	time.Sleep(1 * time.Second)
}
