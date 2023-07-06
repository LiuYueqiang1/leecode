package main

import "fmt"

// 滑动窗口最大值
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 封装单调队列的方式解题
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}
func maxSlidingWindow2(nums []int, k int) []int {
	length := len(nums)
	result := make([]int, 0)
	deque := make([]int, 0) //存放当前元素的索引

	for i := 0; i < length; i++ {
		// 如果队列不为空且队列的第一个元素已经不在滑动窗口中，移除它
		if len(deque) > 0 && deque[0] < i-k+1 { //移除3  索引为1<4-3+1
			deque = deque[1:]
		}

		// 移除队列中所有小于当前元素的元素，它们不可能成为最大值     // 将当前deque中的元素挨个与新元素比较 < 则删掉
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 将当前元素的索引添加到队列中
		deque = append(deque, i)

		// 当滑动窗口的大小达到k时，记录当前窗口的最大值
		if i >= k-1 { //每当 i > 2 开始则记录一次
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//ret := make([]int, 0)
	//ret = append(ret, maxSlidingWindow(nums, 3)...)
	fmt.Println(maxSlidingWindow2(nums, 3))
}
