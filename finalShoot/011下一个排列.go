package finalShoot

import "sort"

/*
	https://www.bilibili.com/video/BV1Nh411k7gw/
	看我自己的讲解吧
*/
func nextPermutation(nums []int) {
	l := len(nums)

	for i := l - 1; l > 0; l-- { //从后往前找
		if nums[i] > nums[i-1] { //找到前一位数比后一位大的那位
			sort.Ints(nums[i:]) //给后面那些排序，不能简单地找比nums[i]大的，要找大，但是又是后面最小的那个
			for j := i; j < l; j++ {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					return
				}
			}
		}
	}

	sort.Ints(nums)
	return
}
