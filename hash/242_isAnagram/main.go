package main

import "fmt"

/*
242.有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
示例 2: 输入: s = "rat", t = "car" 输出: false
说明:你可以假设字符串只包含小写字母
字母异词：字母相同、排列顺序不同
*/

func isAnagram(s string, t string) bool {
	//定义一个数组包含26个字母
	recode := [26]int{}

	// range 一个字符串得到的就是 int32 类型
	//‘a’也是一个 int32 类型
	for _, v := range s {
		recode[v-'a']++
	}
	for _, v := range t {
		recode[v-'a']--
	}
	if recode == [26]int{} {
		return true
	}
	return false
}

func main() {
	s := "apple"
	t := "lepps"
	ret := isAnagram(s, t)
	fmt.Println(ret)
}
