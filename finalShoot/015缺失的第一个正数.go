package finalShoot

/*
	找出数组中缺少的第一个正数， 正数只会出现在1---N+1这个范围内，所以把数字都按照本身的值放在数组里
	nums = [3,4,-1,1]
*/
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {

		//nums[i] = Number,如果Number在[1,N+1]这个范围内，并且Number没在nums[Number-1]位置上，就换过去
		//例如：数值 3 应该放在索引 2 的位置上
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	//都交换完以后，看现在数组里，第一个nums[i] != i+1的，说明index = i 这个位置没对应上该存在的正数，该存在的是 i+1
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1 // 都正确则返回数组长度 + 1

}
