package leetcode

import "strconv"

func restoreIpAddresses(s string) []string {
	var res, path []string
	backTracking(s, path, 0, &res)
	return res
}
func backTracking(s string, path []string, startIndex int, res *[]string) {
	//终止条件
	//1. 遍历到结尾了
	//2. 已经找到四段了
	if startIndex == len(s) && len(path) == 4 {
		tmpIpString := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
		*res = append(*res, tmpIpString)
	}

	for i := startIndex; i < len(s); i++ {
		//处理
		path := append(path, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			//递归
			backTracking(s, path, i+1, res)
		} else { //如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
			return
		}
		//回溯
		path = path[:len(path)-1]
	}
}
func isNormalIp(s string, startIndex, end int) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1])
	if end-startIndex+1 > 1 && s[startIndex] == '0' { //对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
