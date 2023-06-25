package main

import "fmt"

//实现 MyQueue 类：
//
//void push(int x) 将元素 x 压入栈顶。
//int pop() 移除并返回栈顶元素。
//int top() 返回栈顶元素。
//boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

//

// 解释：
//MyStack myStack = new MyStack();
//myStack.push(1);
//myStack.push(2);
//myStack.top(); // 返回 2
//myStack.pop(); // 返回 2
//myStack.empty(); // 返回 False

type MyQueue struct {
	stackTop    int
	stackBottom int
	arr         []int
}

func main() {
	que := MyQueue{
		stackTop:    -1,
		stackBottom: 0,
	}
	que.Push(1)
	que.Push(2)

	que.view()
	fmt.Println(que.Top())
	que.Pop()

	que.view()
	fmt.Println(que.Empty())
	que.Pop()
	que.view()
	fmt.Println(que.Empty())
}
func (this *MyQueue) view() {
	if len(this.arr) == 0 {
		return
	}
	for i := this.stackBottom; i <= this.stackTop; i++ {
		fmt.Printf("%d-->", this.arr[i])
	}
	fmt.Println()
}
func Constructor() MyQueue {
	return MyQueue{
		arr: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
	this.stackTop++
}

func (this *MyQueue) Pop() int {
	if len(this.arr) == 0 {
		return 0
	}
	val := this.arr[this.stackTop]

	this.arr = this.arr[:this.stackTop]
	this.stackTop--
	return val

}

func (this *MyQueue) Top() int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.arr[this.stackTop]
}

func (this *MyQueue) Empty() bool {
	if len(this.arr) == 0 {
		return true
	}
	//fmt.Println(len(this.arr))
	return false
}
