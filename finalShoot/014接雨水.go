package finalShoot

/*
	单调栈 全文背诵
*/
func trap(height []int) int {
	ans := 0
	singleStack := []int{} //里面存的是index

	for i := 0; i < len(height); i++ {

		for len(singleStack) > 0 && height[i] > height[singleStack[len(singleStack)-1]] {
			curr := singleStack[len(singleStack)-1] //取出栈顶元素
			singleStack = singleStack[:len(singleStack)-1]

			if len(singleStack) == 0 { //栈空了就不用循环了
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
