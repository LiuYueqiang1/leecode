package main

import "fmt"

//实现 MyQueue 类：
//
//void push(int x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//boolean empty() 如果队列为空，返回 true ；否则，返回 false
//

// 解释：
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false
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
	que.Push(3)
	que.Push(4)
	que.view()

	que.Pop()
	que.Push(5)
	que.Pop()
	que.Pop()
	que.Pop()
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
	val := this.arr[this.stackBottom]
	this.stackBottom++
	this.arr = this.arr[this.stackBottom:]
	this.stackBottom--
	return val

	//此处很巧妙
	//为了使得数组中的元素真正的出去，而不是只让bottom+1
	//所以我们先让bottom+1
	//将数组替换，此时数组的底部变为0
	// 而bottom=1  ，所以我们再让bottom -1  ，让bottom 变成底部
}

func (this *MyQueue) Peek() int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.arr[this.stackBottom]
}

func (this *MyQueue) Empty() bool {
	if len(this.arr) == 0 {
		return true
	}
	//fmt.Println(len(this.arr))
	return false
}
