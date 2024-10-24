package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func hasCycle(head *ListNode) bool {
	//很明显 用双指针法 超出时间限制
	//快慢指针
	var p, q *ListNode
	//非常牛逼的写法
	//实际存入一个地址 p指针用来循环
	q, p = head, head
	//每次让q比p夺走一次，如果q追上p那么一定存在环 如果没有追上 那么退出 不存在环
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
		if q == p {
			return true
		}
	}
	return false
	//for q != nil {
	//	for p != nil {
	//		if p.Next == q {
	//			return true
	//		}
	//		p = p.Next
	//	}
	//	p = q
	//	q = q.Next
	//}
}

func main() {
	var s1 ListNode = ListNode{1, nil}
	//var s5 ListNode = ListNode{1, &s1}
	//var s6 ListNode = ListNode{2, &s5}
	s4 := &ListNode{4, &s1}
	s3 := &ListNode{3, s4}
	s2 := &ListNode{2, s3}
	s1.Next = s2
	//s6.Next=&s1

	fmt.Println(hasCycle(&s1))
}
