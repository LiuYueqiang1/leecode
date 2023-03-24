package main

import "math"

func mySqrt(x int) int {
	x2 := float64(x)
	ret := math.Sqrt(x2)
	reti := int(ret)
	return reti
}
