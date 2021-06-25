package leetcode

/*
	两个栈分别将"("和"*"的序号压入栈中，
	每次遇到右括号，首先检测左括号栈中是否为空，不为空则弹出元素，否则弹出star栈，
	最后考虑left和star栈可能存在元素，判断此时栈中元素的序号大小，如果left栈中的序号大于star中的则表明 存在"*("此种情况，直接false
*/
func checkValidString(s string) bool {
	var stack1 []int
	var stack2 []int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack1 = append(stack1, i)
		} else if s[i] == '*' {
			stack2 = append(stack2, i)
		} else if s[i] == ')' {

			if len(stack1) > 0 {
				stack1 = stack1[:len(stack1)-1]
			} else if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			} else {
				return false
			}
		}

	}

	for len(stack1) > 0 && len(stack2) > 0 {
		if stack1[len(stack1)-1] > stack2[len(stack2)-1] {
			return false
		}

		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
	}

	if len(stack1) > 0 {
		return false
	}

	return true
}
