package finalShoot

/*
	一般要求空间复杂度O(1)的解法，不要求直接排序或者啥的
	核心思想： 把数字 n 交换到 arr[n] 位置
*/
func findRepeatNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i] //假设取出来的num是2

		if num == i { //如果2就在arr[2]，那没事了
			continue
		}

		//不然判断现在 arr[2] 里面放的是不是2，如果是，那不重了么，找到了就
		if num == nums[num] {
			ans = num
			break
		}

		//不是的话，把2放到arr[2]
		nums[i], nums[num] = nums[num], nums[i]
	}

	return ans
}
