package main

import "sync"

var B *A

type A struct {
	Name string
}

func Sx(S string) {
	var s sync.Once
	s.Do(func() {
		B = &A{
			Name: S,
		}
	})
}
func main() {
	Sx("小红")
}
