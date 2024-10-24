package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	i := 0
	k := make([]string, 0)
	for i < n {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			k = append(k, "FizzBuzz")
		} else if (i+1)%3 == 0 {
			k = append(k, "Fizz")
		} else if (i+1)%5 == 0 {
			k = append(k, "Buzz")
		} else {
			k = append(k, strconv.Itoa(i+1))
		}
		i++
	}
	fmt.Println(k)
	return k
}

func main() {
	fizzBuzz(15)
}
