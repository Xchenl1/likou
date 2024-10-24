package main

//
//import (
//	"math"
//	"strconv"
//)
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func FanZhuan(num int) int {
//	if num == 0 {
//		return 0
//	}
//	var n float64
//	nums := make([]int, 0)
//	for num != 0 {
//		num1 := num % 10
//		nums = append(nums, num1)
//		num /= 10
//	}
//
//	for i, v := range nums {
//		n += float64(v) * math.Pow10(len(nums)-i-1)
//	}
//
//	return int(n)
//}
//
//func ListToNum(list *ListNode) int {
//	num := ""
//	for list != nil {
//		v := strconv.Itoa(list.Val)
//		num = num + v
//		list = list.Next
//	}
//	v, _ := strconv.Atoi(num)
//	return v
//}
//
//func NumToList(num int) *ListNode {
//	if num == 0 {
//		return &ListNode{Val: num}
//	}
//	dummy := &ListNode{}
//	current := dummy
//	for num != 0 {
//		n := num % 10
//		newNode := &ListNode{Val: n}
//		current.Next = newNode
//		current = current.Next
//		num /= 10
//	}
//	return dummy.Next
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	num1 := ListToNum(l1)
//	n1 := FanZhuan(num1)
//	num2 := ListToNum(l2)
//	n2 := FanZhuan(num2)
//	n3 := n1 + n2
//	return NumToList(n3)
//}
//
//func main() {
//	l1 := &ListNode{
//		Val: 0,
//	}
//	l2 := &ListNode{
//		Val:  0,
//		Next: l1,
//	}
//	addTwoNumbers(l1, l2)
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
	//tail.Next = &ListNode{Val: 2}
	//tail = tail.Next
}
