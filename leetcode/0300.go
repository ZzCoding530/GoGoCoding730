package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	ans := dp[0]

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}

		ans = max(ans, dp[i])
	}

	return ans
}
