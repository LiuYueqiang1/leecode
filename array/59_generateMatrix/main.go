package main

import "fmt"

func generateMatrix(n int) [][]int {
	// 初始化矩阵
	matrix := make([][]int, n)
	// 初始化行中的元素
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 里面的数字
	num := 1
	left, right := 0, n-1
	top, bottom := 0, n-1

	for left <= right && top <= bottom { //左闭右闭
		//从左到右
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		//从上到下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		//从右往左
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		//从下往上
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

func main() {
	res := generateMatrix(3)
	for _, v := range res {
		fmt.Printf("%d\n", v)
	}
}
