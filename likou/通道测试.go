package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan string, 1)
	//var a []int
	chan1 <- "123"
	//fmt.Println(<-chan1)
	close(chan1)
	fmt.Println(<-chan1)
	//close(chan1)
}
