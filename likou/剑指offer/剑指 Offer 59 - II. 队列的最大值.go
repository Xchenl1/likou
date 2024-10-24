package main

import "fmt"

// MaxQueue 我的想法是可以采用双指针
type MaxQueue struct {
	Dui   []int
	Paixu []int
}

func Constructor2() MaxQueue {
	return MaxQueue{make([]int, 0), make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Dui) == 0 {
		return -1
	}
	return this.Paixu[0]
}

// Push_back 入队
func (this *MaxQueue) Push_back(value int) {
	this.Dui = append(this.Dui, value)
	if len(this.Paixu) == 0 {
		this.Paixu = append(this.Paixu, value)
		return
	}
	for len(this.Paixu) > 0 && value > this.Paixu[len(this.Paixu)-1] {
		this.Paixu = this.Paixu[:len(this.Paixu)-1]
	}
	this.Paixu = append(this.Paixu, value)
}

// Pop_front 出队
func (this *MaxQueue) Pop_front() int {
	if len(this.Dui) == 0 {
		return -1
	}
	x := this.Dui[0]
	this.Dui = this.Dui[1:]
	if x == this.Paixu[0] {
		this.Paixu = this.Paixu[1:]
	}
	return x
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
func main() {
	obj := Constructor2()
	param_1 := obj.Max_value()
	fmt.Println(param_1)
	obj.Push_back(3)
	param_3 := obj.Pop_front()
	fmt.Println(param_3)
}
