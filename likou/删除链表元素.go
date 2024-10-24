package main

import "fmt"

//type ListNode1 struct {
//	Val  int
//	Next *ListNode1
//}
//
//func removeElements(head *ListNode1, val int) *ListNode1 {
//	if head == nil {
//		return head
//	}
//	//由于链表的头节点有可能需要被删除 搞一个空节点 指向head
//	head2 := head1
//	for head1.Next != nil {
//		if head1.Next.Val == val {
//			head1.Next = head1.Next.Next
//			continue
//		}
//		head1 = head1.Next
//		if head1 == nil {
//			break
//		}
//	}
//	return head2.Next
//}

type Node struct {
	x    int
	next *Node
}

func DeleteNode(root *Node, x1 int) *Node {
	if root == nil {
		return nil
	}
	node1 := &Node{}
	//保存刚开始的头节点
	node := root
	for node != nil {
		if node == root && node.x == x1 {
			root = node.next
			break
		} else if node.x == x1 {
			node1.next = node.next
			break
		}
		node1 = node //保存前一个节点
		node = node.next
	}
	return root
}

//func DeleteNode(root *Node, x int) *Node {
//	if root == nil {
//		return nil
//	}
//	node := root
//	i := 0
//	for node != nil {
//		i++
//		node = node.next
//	}
//	if x > i {
//		return nil
//	}
//	node = root
//	for i := 1; i < x-1; i++ {
//		node = node.next
//	}
//	node.next = node.next.next
//	return root
//}

func main() {
	//var s1 ListNode1 = ListNode1{7, nil}
	//s4 := &ListNode1{7, &s1}
	//s3 := &ListNode1{7, s4}
	//s2 := &ListNode1{7, s3}
	//m := removeElements(s2, 7)
	//for m != nil {
	//	fmt.Println(m.Val)
	//	m = m.Next
	//}
	s1 := &Node{1, nil}
	s2 := &Node{2, s1}
	s3 := &Node{3, s2}
	s4 := &Node{4, s3}
	node1 := DeleteNode(s4, 3)
	fmt.Println(node1)
	for node1 != nil {
		fmt.Println(node1.x)
		node1 = node1.next
	}
}
