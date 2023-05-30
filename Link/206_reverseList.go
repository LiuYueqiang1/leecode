package main

import (
	"fmt"
)

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode2) *ListNode2 {
	pre := head
	cur := head
	temp := head
	tep := head.Next
	for temp != nil {
		temp = temp.Next
	}
	pre.Next = temp

	for tep != nil {
		num := tep
		tep = tep.Next
		num.Next = cur
		cur = num
	}
	return pre
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func addElements2(head, newVal *ListNode2) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = newVal
}
func listElements2(head *ListNode2) {
	temp := head
	//空链表
	if temp.Next == nil {
		fmt.Println("链表为空！")
		return
	}
	for {
		fmt.Print(temp.Next.Val, "==>")
		//退出条件
		temp = temp.Next
		if temp.Next == nil {
			break
		}

	}
}
func main() {
	dummyHead := &ListNode2{}
	v1 := &ListNode2{
		Val: 1,
	}
	v2 := &ListNode2{
		Val: 2,
	}
	v3 := &ListNode2{
		Val: 3,
	}
	v4 := &ListNode2{
		Val: 4,
	}
	v5 := &ListNode2{
		Val: 5,
	}
	v6 := &ListNode2{
		Val: 6,
	}

	addElements2(dummyHead, v1)
	addElements2(dummyHead, v2)
	addElements2(dummyHead, v3)
	addElements2(dummyHead, v4)
	addElements2(dummyHead, v5)
	addElements2(dummyHead, v6)

	listElements2(dummyHead)
	fmt.Println()
	w := reverseList(dummyHead)
	listElements2(w)
}
