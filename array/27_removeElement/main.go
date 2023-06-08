package main

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
// 要知道数组的元素在内存地址中是连续的，不能单独删除数组中的某个元素，只能覆盖。

// 示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
// 你不需要考虑数组中超出新长度后面的元素。
func removeElement1(nums []int, val int) int {
	//与双指针法类似
	//虽然这种写法比较简洁，但它会创建一个新的切片对象，并将原来的数组内容复制到新的切片中返回。
	lenght := len(nums)
	ret := 0
	for i := 0; i < lenght; i++ {
		if nums[i] != val {
			nums[ret] = nums[i]
			ret++
		}
	}
	nums = nums[:ret]
	return len(nums)
}

func removeElement2(nums []int, val int) int {
tag1:
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			//在 nums[i+1:]... 中，nums[i+1:] 是一个切片，而 ... 表示将这个切片中的所有元素打散后传递给 append 函数，
			//相当于依次传递了 nums[i+1], nums[i+2], ... 等元素作为可变参数。
			goto tag1
		}
	}
	return len(nums)
}

// 双指针法
func removeElement3(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast <= len(nums)-1 {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
			//fast++
			//下面出去找个 if 语句已经加了
		}
		fast++
	}

	return len(nums[:slow])
}
