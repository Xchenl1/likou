package main

// import "fmt"
//type MyStack struct {
//	//定义两个切片
//	duilie1, duilie2 []int
//}
//
//func Constructor() MyStack {
//	//直接空切片
//	a := MyStack{}
//	return a
//}
//
//// 入栈
//func (this *MyStack) Push(x int) {
//	//先添加到第二个队列中
//	this.duilie2 = append(this.duilie2, x)
//	//再将队列1中的数据添加过去 如果队列1为空 就不需要向队列2添加数据
//	for len(this.duilie1) != 0 {
//		this.duilie2 = append(this.duilie2, this.duilie1[0])
//		//这步是移向后一个元素
//		this.duilie1 = this.duilie1[1:]
//	}
//	b := this.duilie1
//	this.duilie1 = this.duilie2
//	this.duilie2 = b
//
//}
//
//// 出栈
//func (this *MyStack) Pop() int {
//	if this.Empty() {
//		return -1
//	}
//	v := this.duilie1[0]
//	this.duilie1 = this.duilie1[1:]
//	return v
//}
//
//// 取队头元素
//func (this *MyStack) Top() int {
//	return this.duilie1[0]
//}
//
//// 判断是否为空
//func (this *MyStack) Empty() bool {
//	return len(this.duilie1) == 0
//}

//// 方法是用两个队列来实现栈的操作
//func main() {
//	obj := Constructor()
//	obj.Push(2)
//	obj.Push(3)
//	param5 := obj.Empty()
//	fmt.Println(param5)
//	param2 := obj.Pop()
//	fmt.Println(param2)
//	param3 := obj.Top()
//	fmt.Println(param3)
//	obj.Pop()
//	param4 := obj.Empty()
//	fmt.Println(param4)
//}
