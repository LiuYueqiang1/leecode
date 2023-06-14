package main

import "fmt"

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//输入：s = "We are happy."

//输出："We%20are%20happy."

// " " ASCII码中为 32
func replaceSpace(s string) string {
	b := make([]byte, 0)
	for _, v := range s {
		if v != 32 {
			b = append(b, byte(v))
		} else {
			b = append(b, []byte("%20")...)
		}
	}
	return string(b)
}
func main() {
	s := "We are happy"
	res := replaceSpace(s)
	fmt.Println(res)
}
