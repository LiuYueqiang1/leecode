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
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { //如果快的那个节点不为空的话 说明有环
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			index1 := fast //记录它们相遇的节点
			index2 := head //将index2初始化为头节点
			//而不是在这判断 fast！=slow 因为这个是一直循环的，单凭一次结果不能说明什么
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}

func main() {
	v1 := &ListNode{
		Val: 3,
	}
	v2 := &ListNode{
		Val: 2,
	}
	v3 := &ListNode{
		Val: 0,
	}
	v4 := &ListNode{
		Val: -4,
	}

	addList(v1, v2)
	addList(v1, v3)
	addList(v1, v4)
	v4.Next = v2 //环从v2开始的
	ret := detectCycle(v1)
	fmt.Println(ret)
}

/*
slow:x+y
fast:x+y+n(z+y)
因为fast=2slow
x+y+n(z+y)=2(x+y)
x+y=n(z+y)
入口节点：
x=n(z+y)-y
 =(n-1)(z+y)+z    n>=1
所以令index1在y处，index2在头节点
index1走过 (n-1) 圈的z后必会在入口节点处相遇

*/
