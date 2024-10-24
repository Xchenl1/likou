package main

import "fmt"

type User struct {
	name    string
	nailing int
}
type s struct {
	User
	xingbie string
}

func (u User) a() string {
	fmt.Printf("%p\n", &u)
	return u.name
}

// u是拷贝的 改变值无效
func (u User) a1(name string) {
	u.name = name
}

// 由于是结构体所以指向的是地址 可以改变结构体的值
func (u *User) setName(name string) {
	u.name = name
}

func main() {
	//1.
	var a1 = User{"CHEN", 11}
	fmt.Println(a1)
	//2不传入数据就是默认值
	var a2 = User{
		"CHEN1",
		11,
	}
	fmt.Println(a2)
	//3.
	var a3 User
	a3.name = "chen3"
	a3.nailing = 111
	//4.
	a4 := User{"asd", 12}
	fmt.Println(a4)
	//结构体继承
	var a = s{User{
		"王",
		18,
	}, "男"}
	fmt.Println(a)
	//结构体方法 会继承方法 但是传递是值传递 而不是地址传递
	//这叫实例化
	var user = User{"chen", 12}
	fmt.Printf("%p\n", &user)
	user.a()
	//结构体指针
	user.setName("lisi")
	user.a1("wangwu")
	fmt.Println(user)
	//结构体标签

}
