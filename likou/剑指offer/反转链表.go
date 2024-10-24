package main

import "fmt"

//	type ListNode1 struct {
//		Val  int
//		Next *ListNode1
//	}
//
//	func reverseList(head *ListNode1) *ListNode1 {
//		var s1 *ListNode1
//		head1 := head
//		for head1 != nil {
//			//记录后一个节点
//			l := head1.Next
//			//将下一个设置为新节点
//			head1.Next = s1
//			//把头节点放在下一个节点
//			s1 = head1
//			//把下一个放在l位置
//			head1 = l
//		}
//		return s1
//	}
//
// func main() {
//
// }
type Node struct {
	Val  int
	Next *Node
}

func reverseList(head *Node) *Node {
	var s1 *Node
	for head != nil {
		k1 := head.Next //保存下一个节点
		head.Next = s1  //指向是s1
		s1 = head       //
		head = k1
	}
	return s1
}
func main() {
	s1 := &Node{1, nil}
	s2 := &Node{2, s1}
	s3 := &Node{3, s2}
	node := reverseList(s3)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
