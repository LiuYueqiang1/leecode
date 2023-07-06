package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// 方法一：小顶堆
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent2(nums []int, k int) []int {
	frequency := make(map[int]int)

	// 统计每个元素的出现频率
	for _, num := range nums {
		frequency[num]++
	}

	// 将频率和元素对应起来，用于后续排序
	pairs := make([][2]int, 0)
	for num, freq := range frequency {
		pairs = append(pairs, [2]int{num, freq})
	}

	// 按照频率进行排序   // 从高到低进行排序   将频率
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	// 取出前K个频率最高的元素
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i][0]
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent3(nums, k)
	fmt.Println(result) // 输出 [1, 2]
}

func topKFrequent3(nums []int, k int) []int {
	// 将 nums 写道一个map中  [值]频率
	frequency := make(map[int]int)
	for _, v := range nums {
		frequency[v]++
	}
	//  将 值 和 频率写到一个二维切片中
	pairs := make([][2]int, 0)
	//遍历整个map
	for i, v := range frequency {
		pairs = append(pairs, [2]int{i, v}) //将这个二维数组加到切片中
		//[数值][频率]
		//[数值][频率]
		//[数值][频率]
	}
	// 将此切片按照频率进行排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1] // 按照频率的从大到小排序
	})
	// 拿到频率最高的值写到一个一维切片中
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i][0] // 排完序后只传前 k 个值即可 ,pairs[][0] 将第 0 列写给result
	}
	return result
}
