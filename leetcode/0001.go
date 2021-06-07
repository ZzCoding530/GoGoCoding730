package leetcode

/*
	hashMap	上面存数值本身，下面存index
*/
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, num := range nums {
		if val, ok := hashTable[target-num]; ok { //这里取出来的其实也是索引
			return []int{val, index}
		}
		// key - 值 	value - 索引
		hashTable[num] = index
	}
	return nil
}
