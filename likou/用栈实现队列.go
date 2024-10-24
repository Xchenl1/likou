package main

// 总的思想就是先存数据到zhan1 之后再存入到zhan2
type MyQueue struct {
	zhan1, zhan2 []int
}

func Constructor() MyQueue {
	a := MyQueue{}
	return a
}

// Push 入队
func (this *MyQueue) Push(x int) {
	this.zhan1 = append(this.zhan1, x)
}

// Pop 出队
func (this *MyQueue) Pop() int {
	if len(this.zhan2) == 0 {
		for len(this.zhan1) != 0 {
			//每次将栈1的最后一个元素存到zhan2中 该步会将切片1的长度减一
			this.zhan2 = append(this.zhan2, this.zhan1[len(this.zhan1)-1])
			this.zhan1 = this.zhan1[:len(this.zhan1)-1]
		}
	}
	x := this.zhan2[len(this.zhan2)-1]
	this.zhan2 = this.zhan2[:len(this.zhan2)-1]
	return x
}

// Peek 取出队头元素
func (this *MyQueue) Peek() int {
	if len(this.zhan2) == 0 {
		for len(this.zhan1) != 0 {
			//每次将栈1的最后一个元素存到zhan2中 该步会将切片1的长度减一
			this.zhan2 = append(this.zhan2, this.zhan1[len(this.zhan1)-1])
			this.zhan1 = this.zhan1[:len(this.zhan1)-1]
		}
	}
	return this.zhan2[len(this.zhan2)-1]
}

// Empty 判断是否为空
func (this *MyQueue) Empty() bool {
	return len(this.zhan1)+len(this.zhan2) == 0
}

func main() {

}
