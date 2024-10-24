package main

import "fmt"

type aaa struct {
	name    string
	xingbie int
}

func add(n *int) {
	*n += 1
	fmt.Println("add函数结束时：", n, *n)
}
func shuchu(a *aaa) {
	fmt.Printf("%v |", a)
	//fmt.Println(a.xingbie)
	fmt.Printf("%v", *a)
}
func add1(n int) {
	n++
	fmt.Println("add1函数结束时：", n)
}

type person struct { // 定义一个结构体类型person，包含name和age两个字段
	name string
	age  int
}

func main() {
	//xin := aaa{
	//	name:    "asd",
	//	xingbie: 1,
	//}
	////也就是说*（&xin）==xin
	//fmt.Println(xin)
	//fmt.Println(*(&xin))
	//fmt.Println(xin == *(&xin))
	//fmt.Println(&xin)
	//shuchu(&xin)
	//var y = 2022
	////yy是地址
	//var yy = &y
	//fmt.Printf("%v", yy)
	//*yy++
	//add1(y)
	//fmt.Println(y)
	//add(yy)
	//fmt.Println("调用add函数之后：", y, &y)
	//var num int = 5     // 声明一个变量num，类型为int，值为5
	//var ptr *int = &num // 声明一个指针ptr，类型为*int，值为num的地址
	//fmt.Println(num)    // 输出num的值，结果为5
	//fmt.Println(&num)   // 输出num的地址，结果为一个十六进制的数，如0xc0000160a8
	//fmt.Println(ptr)    // 输出ptr的值，结果与num的地址相同，如0xc0000160a8
	//fmt.Println(*ptr)   // 输出ptr指向的值，结果与num的值相同，如5

	var person1 person = person{"John", 25} // 声明一个变量person1，类型为person，值为{"John", 25}
	var ptr *person = &person1              // 声明一个指针ptr，类型为*person，值为person1的地址
	fmt.Println(person1)                    // 输出person1的值，结果为{John 25}
	fmt.Println(&person1)                   // 输出person1的地址，结果为一个十六进制的数，如0xc00000a0c0
	fmt.Println(ptr)                        // 输出ptr的值，结果与person1的地址相同，如0xc00000a0c0
	fmt.Println(*ptr)                       // 输出ptr指向的值，结果与person1的值相同，如{John 25}
	fmt.Println(person1.name)               // 输出person1的name字段的值，结果为John
	fmt.Println(&person1.name)              // 输出person1的name字段的地址，结果为一个十六进制的数，如0xc000010200
	fmt.Println(ptr.name)                   // 输出ptr指向的name字段的值，结果与person1的name字段的值相同，如John
	fmt.Println((*ptr).name)                // 输出ptr指向的name字段的值，结果与person1的name字段的值相同，如John
}
