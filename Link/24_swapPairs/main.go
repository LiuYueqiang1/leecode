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
	sp := swapPairs(v1)
	fmt.Println()
	viewList(sp)
}

// 反转链表
func swapPairs(head *ListNode) *ListNode {
	//定义虚拟头节点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead
	temp := head

	for temp != nil && temp.Next != nil {
		pre.Next = temp.Next
		newcur := temp.Next.Next
		temp.Next.Next = temp
		temp.Next = newcur
		pre = temp
		temp = newcur
	}

	return dummyHead.Next

}
