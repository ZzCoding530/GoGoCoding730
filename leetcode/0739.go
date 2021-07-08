package leetcode

/*
	单调栈即可
*/
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length) //返回值
	stack := []int{}           //单调栈

	for i := 0; i < length; i++ {
		temperature := temperatures[i]                                          //当前温度
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] { //单调栈模板
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
