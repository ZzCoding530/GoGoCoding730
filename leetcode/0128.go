package leetcode

/*
	看官方题解
*/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mySet := make(map[int]bool)

	for _, num := range nums {
		mySet[num] = true
	}

	curr := 1
	longest := 1
	for numSet := range mySet {
		if !mySet[numSet-1] {
			curr = 1
			for mySet[numSet+1] {
				curr++
				numSet++
			}

			if longest < curr {
				longest = curr
			}

		}

	}

	return longest
}
