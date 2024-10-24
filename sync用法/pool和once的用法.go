package main

import (
	"fmt"
	"sync"
)

type Object struct {
	Name string
}

var once sync.Once

var pool = sync.Pool{
	New: func() interface{} {
		return &Object{}
	},
}

func getObject() *Object {
	obj := pool.Get().(*Object) //从对象池中获取一个对象
	if obj.Name == "" {
		once.Do(func() {
			obj.Name = "Alice"
		})
	}
	return obj
}

func putObject(obj *Object) {
	obj.Name = ""
	pool.Put(obj)
}

func main() {
	// 获取对象并输出名称
	obj1 := getObject()
	fmt.Println(obj1.Name)

	// 放回对象池并获取另一个对象
	putObject(obj1)
	obj2 := getObject()
	fmt.Println(obj2.Name)
}
