package main

import (
	"fmt"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 例如，121 是回文，而 123 不是
//
// 想法：
// 将这个每个都拿出来，放到一个切片中，用切片中的第一个数与len(n-1)做比较

func isPalindrome(x int) bool {
	// 1、将这个数每个都拿出来
	s1 := make([]int, 0)
	for x > 0 {
		a := x % 10
		x = x / 10
		s1 = append(s1, a)
	}
	fmt.Println(len(s1))
	//s2 := make([]int, 0)
	//2、判断第一个和最后一个是否相等
	for i := 0; i < len(s1)/2; {
		//s1[0]==s1[len(s1)-1]
		//s1[1]==s1[len(s1)-2]
		//s1[2]==s1[len(s1)-3]
		//...
		//s1[i]==s1[len(s1)-i-1]
	tag:
		i++
		if s1[i] == s1[len(s1)-i-1] {
			//循环结果如果全部相等，s2切片就记录1
			//s3 := fmt.Sprint(1)
			//s4, _ := strconv.Atoi(s3)
			//s2 = append(s2, s4)
			if i == len(s1)/2-1 {
				//fmt.Println(s2)
				return true
				break
			}
			goto tag

		} else {
			return false
		}
	}

	return false
}
func main() {
	pan := isPalindrome(3125655213)
	fmt.Println(pan)
}
