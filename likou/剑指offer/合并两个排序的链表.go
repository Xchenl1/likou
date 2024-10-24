package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	} else if list2 == nil {
//		return list1
//	}
//	lode3 := &ListNode{}
//	node := lode3
//	for list1 != nil && list2 != nil {
//		if list1.Val <= list2.Val {
//			lode3.Next = list1
//			list1 = list1.Next
//		} else {
//			lode3.Next = list2
//			list2 = list2.Next
//		}
//		lode3 = lode3.Next
//		//下次循环
//	}
//	if list1 != nil {
//		for list1 != nil {
//			lode3.Next = list1
//			list1 = list1.Next
//			lode3 = lode3.Next
//		}
//		return node.Next
//	}
//	if list2 != nil {
//		for list2 != nil {
//			lode3.Next = list2
//			list2 = list2.Next
//			lode3 = lode3.Next
//		}
//	}
//	return node.Next
//}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	//可以循环向后撤查找、
//	k3 := &ListNode{}
//	k4 := k3
//	for l1 != nil && l2 != nil {
//		if l1.Val <= l2.Val {
//			k3.Next = l1
//			l1 = l1.Next
//		} else {
//			k3.Next = l2
//			l2 = l2.Next
//		}
//		k3 = k3.Next
//	}
//	if l2 != nil {
//		k3.Next = l2
//		l2 = l2.Next
//	}
//	if l1 != nil {
//		k3.Next = l1
//		l1 = l1.Next
//	}
//	return k4.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	head := l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	for l1 != nil {
		l3.Next = l1
		l1 = l1.Next
		l3 = l3.Next
	}
	for l2 != nil {
		l3.Next = l2
		l2 = l2.Next
		l3 = l3.Next
	}
	return head.Next
}

func main() {
	var s1 = ListNode{8, nil}
	s4 := &ListNode{7, &s1}
	s3 := &ListNode{6, s4}
	s2 := &ListNode{5, s3}
	var s11 = ListNode{8, nil}
	s44 := &ListNode{7, &s11}
	s33 := &ListNode{6, s44}
	s22 := &ListNode{5, s33}
	m := mergeTwoLists(s2, s22)
	for m != nil {
		fmt.Println(m.Val)
		m = m.Next
	}
}
