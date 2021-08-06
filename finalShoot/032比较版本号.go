package finalShoot

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	arr1, arr2 := strings.Split(version1, "."), strings.Split(version2, ".") //分割一下得到俩数组

	for i := 0; i < max(len(arr1), len(arr2)); i++ {
		n1, n2 := getValue(arr1, i), getValue(arr2, i) //n1, n2 是俩数字
		if n1 < n2 {                                   //小就-1
			return -1
		} else if n1 > n2 { //不然 1
			return 1
		} //都不是就比较下一段
	}
	return 0
}

/*
	输入一个string的字符串，输出转化成的数字int类型
*/
func getValue(arr []string, n int) int {
	if n < len(arr) {
		ans, _ := strconv.Atoi(arr[n])
		return ans
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
