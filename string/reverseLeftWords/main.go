package main

import "fmt"

//字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
//请定义一个函数实现字符串左旋转操作的功能。
//比如，输入字符串"abcdefg"和数字2
//该函数将返回左旋转两位得到的结果"cdefgab"。

// 思路 ：局部反转+整体反转
// 1、先反转前n个字符
// 2、再反转n+1到末尾
// 3、再全部反转
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(&b, 0, n-1)
	reverse(&b, n, len(b)-1)
	reverse(&b, 0, len(b)-1)
	return string(b)
}
func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
func main() {
	s := "abcdefg"
	res := reverseLeftWords(s, 2)
	fmt.Println(res)
}
