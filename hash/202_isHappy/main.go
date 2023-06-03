package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true //因为我们给一个循环中所有的数都赋值为true了
		// 而bool值默认为false，所以进入循环的一直是 ture ，则进入循环
		// 将m[4]=true后，那么！m[4]=false，则为退出循环的条件，此时退出循环
		n = getSum(n)
		//n,m[n]=getSum(n),true 多重赋值，好处是可以两个参数一起赋值
		//比如是这样
		//n = getSum(n)
		//m[n] = true
		//那么 m[n]用到就不是上一个值，而是新的n，我们要保证两个都是同一个n
	}
	fmt.Println(m)
	return n == 1

	/*
	  for n != 1 {
	        if m[n] {
	            return false
	        }
	        m[n] = true
	        n = getSum(n)
	    }
	    return true
	*/
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10) //只保留个位数
		n = n / 10
	}
	return sum
}
func main() {
	bo := isHappy(4)
	fmt.Println(bo)
}
