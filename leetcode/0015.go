package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	l := len(nums)

	if nums == nil || l < 3 {
		return ans
	}

	sort.Ints(nums) //先排序

	for i := 0; i < l; i++ {
		if nums[i] > 0 { //如果当前值都大于0，那和后面的数相加不可能等于0，结束循环
			break
		}

		if i > 0 && nums[i] == nums[i-1] { //如果跟刚才前一个数一样，那不用算了，继续往后走
			continue
		}

		left, right := i+1, l-1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]

			if sum == 0 {
				ans = append(ans, []int{nums[left], nums[right], nums[i]}) //等于0就加进去答案

				for left < right && nums[left] == nums[left+1] { //避免重复的答案
					left++
				}

				for left < right && nums[right] == nums[right-1] { //避免重复的答案
					right--
				}

				left++
				right--
			} else if sum > 0 { //大了就右侧指针往小一点指，往前走一位
				right--
			} else { //小了就左侧指针往大一点指，往后走一位
				left++
			}
		}
	}

	return ans

}
