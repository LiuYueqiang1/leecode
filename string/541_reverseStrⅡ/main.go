package main

import "fmt"

//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

// 输入：s = "abcdefg", k = 2
// 输出："bacd efgjj"  --> "bacd fegj j" --> "bacd fegj j"

// 1、每次移动 4 个
// 2、每次交换前两个 0-1 ，4-5，8-9， 交换 [i:i+k-1]  i+=4
// 3、剩下的：
func reverseStr(s string, k int) string {
	lengh := len(s)
	ss := []byte(s)

	for i := 0; i < lengh; i += 2 * k {
		if i+k <= lengh { //只要i+k在 字符串里面，就交换
			reverse(ss[i : i+k]) //交换的就是 [i:i+k-1]   因为下面的for循环是left取不到right的
		} else {
			reverse(ss[i:lengh]) //否则直接交换剩下的，中间的那部分直接不管的
		}
	}
	return string(ss)
}

// 字符串前后反转的函数 344
func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	s := "abcdefgjj"
	res := reverseStr(s, 2)
	fmt.Println(res)
}
