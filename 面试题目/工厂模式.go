package main

import (
	"fmt"
	"sync"
)

// 工厂模式
type Show struct {
	Name string
}

func NewShow(name string) *Show {
	return &Show{name}
}

func (s *Show) String() string {
	return s.Name
}

//func main() {
//	a := NewShow("阿迪达斯")
//	fmt.Println(a.String())
//}

//抽象工厂

type Fly interface {
	fly()
}

type Bird struct {
	Name  string
	Color string
}

func NewBird(name string, color string) *Bird {
	return &Bird{name, color}
}

func (b *Bird) fly() {
	fmt.Println(b.Name + "会飞")
}

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	return &Dog{name}
}

func (d *Dog) fly() {
	fmt.Println(d.name + "会飞")
}

func TestAnimal(name string) Fly {
	switch name {
	case "dog":
		return NewDog(name)

	case "bird":
		return NewBird(name, "")
	default:
		return nil
	}
}

//func main() {
//	dog := TestAnimal("dog")
//	dog.fly()
//}

//懒汉式单例模式

var bb *aa
var one sync.Once

type aa struct {
	name string
}

func NewBB(name string) {
	one.Do(func() {
		bb = &aa{name: name}
	})
}

func (c *aa) GetName() string {
	return c.name
}

// 饿汉式单例模式
var cc = &aa{name: "CCC"}

func main() {
	NewBB("111")
	fmt.Println(bb)
	fmt.Println(cc.GetName())
}
