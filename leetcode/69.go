package leetcode

import "math"

func mySqrt(x int) int {
	start := x / 2

	for start*start < x {
		start = start + 1
	}

	return start - 1
}

//不用二分还是慢

func mySqrt(x int) int {
	l, r := 0, math.MaxInt32 //两个边界  0至正无穷
	for l < r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}
