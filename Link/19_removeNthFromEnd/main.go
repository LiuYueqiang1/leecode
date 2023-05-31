package main

import (
	"fmt"
)

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
	rr := removeNthFromEnd(v1, 2)
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
	if n > num {
		fmt.Println("您删除的值不存在")
		return head
	} else if n == 1 && num == 1 {
		head = nil
		return head
	}
	//temp = head
	pre := dummyHead
	for i := 0; i < num-n; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	//最后返回的一定要是dummyHead.Next
	//不然如果是 head 的话就失去了 dummyHead 的意义
	return dummyHead.Next

}
