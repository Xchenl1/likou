package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var k *ListNode
	l := head
	for l != nil {
		//先保存该节点
		a := l.Next
		//让l指向前一个结点
		l.Next = k
		//再将现在的节点保存起来
		k = l
		//然后将保存的a再赋值给l
		l = a
	}
	return k
}

func main() {
	var s1 ListNode = ListNode{1, nil}
	//var s5 ListNode = ListNode{1, &s1}
	//var s6 ListNode = ListNode{2, &s5}
	s4 := &ListNode{4, &s1}
	s3 := &ListNode{3, s4}
	s2 := &ListNode{2, s3}
	l1 := reverseList(s2)
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}
