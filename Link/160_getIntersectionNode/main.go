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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//定义两个链表的头节点
	tempA := headA
	tempB := headB
	//两个链表的长度
	numA := 1
	numB := 1
	for tempA != nil && tempA.Next != nil {
		numA++
		tempA = tempA.Next
	}
	for tempB != nil && tempB.Next != nil {
		numB++
		tempB = tempB.Next
	}
	//将头节点置于初始位置
	tempA = headA
	tempB = headB
	//比较两个链表的长度,并将长链表的头节点向后移动
	if numA >= numB {
		for i := 0; i < numA-numB; i++ {
			tempA = tempA.Next
		}
	} else {
		for i := 0; i < numB-numA; i++ {
			tempB = tempB.Next
		}
	}
	//考虑到其中某个链表为空时的情况
	if tempA == nil || tempB == nil {
		return nil
	}
	for tempB.Val != tempA.Val {
		tempA = tempA.Next
		tempB = tempB.Next
	}
	return tempB
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
	v5 := &ListNode{
		Val: 6,
	}
	h1 := &ListNode{
		Val: 5,
	}
	h2 := &ListNode{
		Val: 4,
	}
	h3 := &ListNode{
		Val: 6,
	}
	addList(v1, v2)
	addList(v1, v3)
	addList(v1, v4)
	addList(v1, v5)
	addList(h1, h2)
	addList(h1, h3)
	fmt.Println("链表1：")
	viewList(v1)
	fmt.Println()
	fmt.Println("链表2：")
	viewList(h1)
	fmt.Println()
	ret := getIntersectionNode(v1, h1)
	fmt.Println(ret.Val)
}
