package leetcode

/*
	单调栈 建议全文背诵
*/
func trap(height []int) int {
	ans := 0
	var singleStack []int

	for i := 0; i < len(height); i++ {

		for len(singleStack) > 0 && height[i] > height[singleStack[len(singleStack)-1]] {
			curr := singleStack[len(singleStack)-1]
			singleStack = singleStack[:len(singleStack)-1]

			if len(singleStack) == 0 {
				break
			}

			l := singleStack[len(singleStack)-1]
			r := i

			ans = ans + (r-l-1)*(min(height[l], height[r])-height[curr])

		}

		singleStack = append(singleStack, i)

	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
