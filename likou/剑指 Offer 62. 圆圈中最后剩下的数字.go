package main

import "fmt"

//	type listnode struct {
//		node int
//		next *listnode
//	}
//
//	func lastRemaining(n int, m int) int {
//		//用链表做 然后之后循环m次 循环链表
//		head := &listnode{}
//		l := &listnode{}
//		for i := 1; i < n; i++ {
//			listnodei := listnode{i, nil}
//			l.next = &listnodei
//			if i == 1 {
//				head = l
//				//fmt.Println(head)
//			}
//			l = &listnodei
//		}
//		l.next = head // 形成循环链表
//		i := 1
//		k := head
//		for {
//			if i == m-1 {
//				k.next = k.next.next
//				k = k.next
//				i = 1
//			}
//			if k.next == k {
//				return k.node
//			}
//			k = k.next
//			i++
//		}
//	}
//
// 参考答案理解的
func lastRemaining(n int, m int) int {
	idx := 0
	for i := 2; i <= n; i++ {
		idx = (idx + m) % i
	}
	return idx
}

func main() {
	fmt.Println(lastRemaining(10, 17))
}
