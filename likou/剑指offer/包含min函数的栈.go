package main

import "fmt"

// 先进后出
type MinStack struct {
	Qiepian1, Qiepian2 []int
}

func Constructor1() MinStack {
	return MinStack{}
}

// 入栈
func (this *MinStack) Push(x int) {
	//数据存入切片1中
	this.Qiepian1 = append(this.Qiepian1, x)
}

// 出栈
func (this *MinStack) Pop() {
	//先看切片2中有没有数据
	if len(this.Qiepian1) == 0 {
		for len(this.Qiepian2) != 0 {
			this.Qiepian1 = append(this.Qiepian2, this.Qiepian2[0])
			this.Qiepian2 = this.Qiepian2[1:]
		}
	}
	this.Qiepian1 = this.Qiepian1[:len(this.Qiepian1)-1]
}

// 取栈顶元素
func (this *MinStack) Top() int {
	return this.Qiepian1[len(this.Qiepian1)-1]
}

// 取最小数据
func (this *MinStack) Min() int {
	k := this.Qiepian1[0]
	for i := 0; i < len(this.Qiepian1); i++ {
		if this.Qiepian1[i] <= k {
			k = this.Qiepian1[i]
		}
	}
	return k
}

func main() {

	//Your MinStack object will be instantiated and called as such:
	obj := Constructor1()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.Min())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.Min())
}
