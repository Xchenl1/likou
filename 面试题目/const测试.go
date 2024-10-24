package main

const (
	x = iota
	_
	y
	z = "zz"
	k
	j
	p = iota
)
const (
	x2 = 1
	_
	u
	x1 = iota
	s
	p1
)

//关于const中的iota中的关键字  iota是枚举类型 从第一行开始 如果定义变量没有显示调用 默认+1或者取上方的值
