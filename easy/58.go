package main

import (
	"fmt"
)

// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
// 输入：s = "Hello World"
// 输出：5
// 解释：最后一个单词是“World”，长度为5。
func lengthOfLastWord(s string) {
	//1、遍历所有的字符串
	//2、如果遇到空格，返回此时的index，将 [:index] 的放到一个切片中
	//3、继续遍历后面的字符串，执行第二步，判断这个切片长度==0，则丢弃这个切片
	//4、没有空格后，返回上一个空格的索引，[index:],放入一个切片中
	s0 := make([]int32, 0, 10)
	for i, v := range s {
		s0 = append(s0, v)
		//	fmt.Printf("%T %v\n", s0, s0)
		if v == 32 {
			s1 := make([]int32, 0, 10)
			s1 = append(s1, s0[:i]...)
			fmt.Println(s1)
		}
	}
}
func main() {
	str := "Hello  W orld"
	lengthOfLastWord(str)

}
