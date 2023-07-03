package main

import (
	"fmt"
	"strconv"
)

//逆波兰表达式
//输入：tokens = ["2","1","+","3","*"]
//输出：9
//解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

// 思想：
// 栈
// 遇到符号则取出两个数做运算
// 否则加入到栈中
// 将 string 转化为 int ，运算后后在转化为 string 放入栈中
type stack struct {
	topStack     int
	botttomStack int
	arr          []string
}

func evalRPN(tokens []string) int {
	st := &stack{
		topStack:     -1,
		botttomStack: 0,
	}
	for _, v := range tokens {
		switch v {
		case "+":
			a, _ := strconv.Atoi(st.arr[st.topStack-1])
			b, _ := strconv.Atoi(st.arr[st.topStack])
			sum := strconv.Itoa(a + b) // 将其转化为 int 类型的相加后 转化为 string类型
			st.topStack -= 2
			st.arr = st.arr[:st.topStack+1]
			st.arr = append(st.arr, sum)
			st.topStack++
		case "-":
			a, _ := strconv.Atoi(st.arr[st.topStack-1])
			b, _ := strconv.Atoi(st.arr[st.topStack])
			sum := strconv.Itoa(a - b) // 将其转化为 int 类型的相加后 转化为 string类型
			st.topStack -= 2
			st.arr = st.arr[:st.topStack+1]
			st.arr = append(st.arr, sum)
			st.topStack++
		case "*":
			a, _ := strconv.Atoi(st.arr[st.topStack-1])
			b, _ := strconv.Atoi(st.arr[st.topStack])
			sum := strconv.Itoa(a * b) // 将其转化为 int 类型的相加后 转化为 string类型
			st.topStack -= 2
			st.arr = st.arr[:st.topStack+1]
			st.arr = append(st.arr, sum)
			st.topStack++
		case "/":
			a, _ := strconv.Atoi(st.arr[st.topStack-1])
			b, _ := strconv.Atoi(st.arr[st.topStack])
			sum := strconv.Itoa(a / b) // 将其转化为 int 类型的相加后 转化为 string类型
			st.topStack -= 2
			st.arr = st.arr[:st.topStack+1]
			st.arr = append(st.arr, sum)
			st.topStack++
		default:
			st.arr = append(st.arr, v)
			st.topStack++
		}

		//if v == "+" || v == "-" || v == "*" || v == "/" {
		//	a, _ := strconv.Atoi(st.arr[st.topStack])
		//	b, _ := strconv.Atoi(st.arr[st.topStack-1])
		//	sum := strconv.Itoa(a + b) // 将其转化为 int 类型的相加后 转化为 string类型
		//	st.topStack -= 2
		//	st.arr = st.arr[:st.topStack+1]
		//	st.arr = append(st.arr, sum)
		//	st.topStack++
		//} else {
		//	st.arr = append(st.arr, v)
		//	st.topStack++
		//}
	}
	ret, _ := strconv.Atoi(st.arr[st.topStack])
	return ret
}
func main() {
	ret := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	fmt.Println(ret)
}
