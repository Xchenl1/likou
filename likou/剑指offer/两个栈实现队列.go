package main

import "fmt"

type CQueue struct {
	zhan1, zhan2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

// 栈1 记录数据
func (this *CQueue) AppendTail(value int) {
	this.zhan1 = append(this.zhan1, value)
}

// 删除队头
func (this *CQueue) DeleteHead() int {
	if len(this.zhan1) == 0 && len(this.zhan2) == 0 {
		return -1
	}
	if len(this.zhan2) == 0 {
		for len(this.zhan1) != 0 {
			//将栈1逆序添加到栈2中
			this.zhan2 = append(this.zhan2, this.zhan1[len(this.zhan1)-1])
			this.zhan1 = this.zhan1[:len(this.zhan1)-1]
		}
	}
	//如果栈2中还有数据 那么就直接将栈2的最后一个元素返回
	x := this.zhan2[len(this.zhan2)-1]
	this.zhan2 = this.zhan2[:len(this.zhan2)-1] //左闭右开
	return x
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
}
