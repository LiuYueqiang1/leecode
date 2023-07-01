package main

import "fmt"

// 给出由小写字母组成的字符串S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
//
// 在 S 上反复执行重复项删除操作，直到无法继续删除。
//
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一

//输入："abbaca"
//输出："ca"

// 使用栈操作。
// a->  入栈 a
// b跟a比较 不相等 b-> 入栈 ab
// b跟b比较 ，相等  b出栈  a
// a跟a比较 ，相等  a出栈  “"
// c入栈  c
// a入栈  ca
type stack struct {
	topStack    int
	bottomStack int
	arr         []rune
}

func removeDuplicates(s string) string {
	st := &stack{
		topStack:    -1,
		bottomStack: 0,
	}
	for _, v := range s {
		if st.topStack == -1 {
			st.arr = append(st.arr, v)
			st.topStack++
		} else if v != st.arr[st.topStack] {
			st.arr = append(st.arr, v)
			st.topStack++
		} else {
			st.arr = st.arr[:st.topStack]
			st.topStack--
		}
	}
	res := st.arr[:st.topStack+1]
	rest := string(res)
	return rest
}

func main() {
	res := removeDuplicates("abbaca")
	fmt.Println(res)
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[:3]) //1,2,3
	fmt.Println(s[:0]) //[]
}
