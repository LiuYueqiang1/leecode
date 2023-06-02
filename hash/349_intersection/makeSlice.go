package main

import "fmt"

// make([]T, size, cap)
func main() {
	s := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		s = append(s, fmt.Sprintf("%v", i))
	}
	fmt.Println(s) //[     0 1 2 3 4 5 6 7 8 9] append前五个为空
}

//切片无 DELETE 函数
//map有 delete 函数  delete(map,key)
