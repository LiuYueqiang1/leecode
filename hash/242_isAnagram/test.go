package main

import (
	"fmt"
	"strconv"
)

// 测试字符转换
func main() {
	exp := "3+3*4-3"
	ch := exp[0:1]
	ch1 := exp[3:4]
	//fmt.Println(ch)  3

	//temp1, _ := strconv.Atoi(ch)
	temp3, _ := strconv.Atoi(ch1)
	//fmt.Println(temp1) //3
	fmt.Println("temp3", temp3) // * 但是确是 0
	temp2 := int([]byte(ch)[0]) // 转换为ASCII码
	fmt.Println(temp2)          //51   ASCII码
	fmt.Println(rune(temp2))    // 转换为ASCII码  51

	fmt.Printf("%#v\n", ch)         //String 类型
	temp4 := int32([]byte(ch)[0])   // 转换为ASCII码   将String类型转换为byte类型，再转换为int类型
	fmt.Printf("temp4:%T\n", temp4) //int32 类型

	fmt.Printf("%T\n", 'a')       //int32 类型
	fmt.Printf("%T\n", rune('a')) //int32 类型

	for _, v := range "abd" {
		fmt.Printf("%T:\n", v) //int32
	}
	//res := temp1 + temp3 + temp1
	//
	//fmt.Println("res:", res)
}
