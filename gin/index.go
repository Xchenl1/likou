package main

import (
	"fmt"
)

type student struct {
	name   string
	age    int
	weight float32
}
type point struct {
	x, y int
}

func main() {
	s := student{name: "孙悟空", age: 500, weight: 1000.89}
	fmt.Printf("student struct %% %v\n", s)
	fmt.Printf("student struct %%+v %+v\n", s)
	fmt.Printf("student struct %%#v %#v\n", s)
	fmt.Printf("student struct %%T %T\n", s)
	fmt.Printf("student age %%b %b\n", s.age)
	//sp := &student{name: "猪八戒", age: 500, weight: 800.99}
	//fmt.Printf("student struct %%v %v\n", sp)
	//fmt.Printf("student struct %%+v %+v\n", sp)
	//fmt.Printf("student struct %%#v %#v\n", sp)
	//fmt.Printf("student struct %%p %p\n", sp)
	//fmt.Printf("student struct %%T %T\n", sp)
	//p := point{1, 2}
	//fmt.Printf("%v\n", p)
	//fmt.Printf("%+v\n", p)
	//fmt.Printf("%#v\n", p)
	//fmt.Printf("%T\n", p)
	//fmt.Printf("%t\n", true)
	//fmt.Printf("%d\n", 123)
	//fmt.Printf("%b\n", 14)
	//fmt.Printf("%c\n", 33)
	//fmt.Printf("%x\n", 456)
	//fmt.Printf("%f\n", 78.9)
	//fmt.Printf("%e\n", 123400000.0)
	//fmt.Printf("%E\n", 123400000.0)
	//fmt.Printf("%s\n", "\"string\"")
	//fmt.Printf("%q\n", "\"string\"")
	//fmt.Printf("%x\n", "hex this")
	//fmt.Printf("%p\n", &p)
	//fmt.Printf("|%6d|%6d|\n", 12, 345)
	//fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	//fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//fmt.Printf("|%6s|%6s|\n", "foo", "b")
	//fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	//ss := fmt.Sprintf("a %s", "string")
	//fmt.Println(ss)
	//fmt.Fprintf(os.Stderr, "an %s\n", "error")
	//for _, r := range "123qbc中午" {
	//	cvt := string(r)
	//	if r >= 128 {
	//		cvt = fmt.Sprintf("\\u%04x", int64(r))
	//	}
	//	println(cvt)
	//}
}
