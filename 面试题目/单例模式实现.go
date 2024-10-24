package main

import "sync"

var SX *A

type A struct {
	Name string
}

// 懒汉式
func NewA() *A {
	a := A{
		Name: "小陈",
	}
	SX = &a
	return SX
}

// 饿汉式
func NewA1() *A {
	var s sync.Once
	s.Do(func() {
		SX = &A{Name: "小陈"}
	})
	return SX
}

func main() {

	NewA1()
	_ = NewA()

}
