package leetcode

func canJump(nums []int) bool {
	longestIndex := 0 //表示上一步最远能到达的index

	for i := 0; i < len(nums); i++ {
		if i > longestIndex { // 如果我现在处于的index就已经超过了上一步能到达的最远距离，那说明到不了我这index，也就到不了最后了啊
			return false
		}

		max(i+nums[i], longestIndex) //最远能到达的距离更新一下，当前位置加上能走的距离
	}

	return true
}
