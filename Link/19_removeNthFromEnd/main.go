package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addList(head, newlist *ListNode) {
	//定义虚拟头节点
	dummyHead := &ListNode{Next: head}
	temp := dummyHead
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newlist
}
func viewList(head *ListNode) {
	//定义虚拟头节点
	dummyHead := &ListNode{Next: head}
	temp := dummyHead
	for temp != nil && temp.Next != nil { //dummyHead 也不允许为空
		fmt.Printf("%d =>", temp.Next.Val)
		temp = temp.Next
	}
}

func main() {
	v1 := &ListNode{
		Val: 1,
	}
	v2 := &ListNode{
		Val: 2,
	}
	v3 := &ListNode{
		Val: 3,
	}
	v4 := &ListNode{
		Val: 4,
	}

	addList(v1, v2)
	addList(v1, v3)
	addList(v1, v4)
	viewList(v1)
	fmt.Println()
	rr := removeNthFromEnd(v1, 1)
	viewList(rr)
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	temp := head
	num := 1
	for temp.Next != nil {
		num++
		temp = temp.Next
	}
	temp = head
	for i := 1; i < num-n; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return head
}
