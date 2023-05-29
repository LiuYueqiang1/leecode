package main

import (
	"fmt"
)

//在链表类中实现这些功能：
//
//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
//如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

type MyLinkedList struct {
	// Id   int
	Val  int
	Next *MyLinkedList
}

func main() {
	dummyHead := MyLinkedList{}
	dummyHead.AddAtTail(2)
	dummyHead.AddAtTail(3)
	dummyHead.AddAtHead(1)
	dummyHead.listElement()
	dummyHead.AddAtIndex(3, 4)
	fmt.Println()
	dummyHead.listElement()
	fmt.Println()
	var choice int
	fmt.Println("请输入您要查看的值的位置：")
	fmt.Scanln(&choice)
	val := dummyHead.Get(choice)
	fmt.Printf("您要查看 %d 位的值为：%d\n", choice, val)
	var del int
	fmt.Println("请输入您要删除的值的位置：")
	fmt.Scanln(&del)
	dummyHead.DeleteAtIndex(del)
	fmt.Printf("您要删除第 %d 位的值\n", del)
	fmt.Println("删除后的链表为：")
	dummyHead.listElement()
	Constructor()
}

// 打印
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) listElement() {
	temp := this
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

// 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	temp := this
	//var newLinkList = new(MyLinkedList)
	lens := 0
	for temp.Next != nil {
		lens++
		temp = temp.Next
	}
	temp = this
	if index > lens {
		return -1
	}
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.Val
}

// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	temp := this
	var newLinkList = new(MyLinkedList)
	newLinkList.Next = temp.Next
	temp.Next = newLinkList
	temp.Next.Val = val

}

// 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {

	temp := this
	//把 temp 置于最后一位
	for temp.Next != nil {
		temp = temp.Next
	}
	var newLinkList = new(MyLinkedList)
	temp.Next = newLinkList
	temp.Next.Val = val
}

//在链表中的第 index 个节点之前添加值为 val  的节点。
//如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	temp := this
	var newLinkList = new(MyLinkedList)
	lens := 0
	for temp.Next != nil {
		lens++
		temp = temp.Next
	}
	temp = this
	if index > lens {
		return
	} else if index < 0 {
		newLinkList.Next = temp.Next
		temp.Next = newLinkList
		temp.Next.Val = val
	} else if index == lens {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newLinkList
		temp.Next.Val = val
	}
}

// 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	temp := this
	//var newLinkList = new(MyLinkedList)
	lens := 0
	for temp.Next != nil {
		lens++
		temp = temp.Next
	}
	temp = this
	if index > lens {
		fmt.Println("您输入的索引无效")
		return
	}
	//保证要删除的是前一位
	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}
	if temp.Next != nil {
		temp.Next = temp.Next.Next
	} else {
		temp.Next = nil
	}

}
