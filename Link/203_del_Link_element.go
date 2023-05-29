package main

import "fmt"

// 题意：删除链表中等于给定值 val 的所有节点。
//
// 示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]
//
// 示例 2： 输入：head = [], val = 1 输出：[]
//
// 示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
type ListNode struct {
	Val  int
	Next *ListNode
}

func addElements(head, newVal *ListNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = newVal
}
func listElements(head *ListNode) {
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

func removeElements(head *ListNode, val int) *ListNode {
	temp := head
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.Val == val {
			if temp.Next.Next == nil {
				temp.Next = nil
				break
			}
			temp.Next = temp.Next.Next
		}
		temp = temp.Next
	}
	return temp
}
func main() {
	head := &ListNode{}
	fmt.Println("test:", head == nil)
	v1 := &ListNode{
		Val: 6,
	}
	v2 := &ListNode{
		Val: 6,
	}
	v3 := &ListNode{
		Val: 6,
	}
	v4 := &ListNode{
		Val: 6,
	}
	v5 := &ListNode{
		Val: 6,
	}
	v6 := &ListNode{
		Val: 6,
	}
	v7 := &ListNode{
		Val: 6,
	}
	addElements(head, v1)
	addElements(head, v2)
	addElements(head, v3)
	addElements(head, v4)
	addElements(head, v5)
	addElements(head, v6)
	addElements(head, v7)
	listElements(head)
	fmt.Println()
	removeElements(head, 6)
	removeElements(head, 5)
	removeElements(head, 1)
	listElements(head)
}
