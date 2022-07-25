package utils

import "strconv"

// @Author: 张佳伟
// @Function:Intersection
// @Description:比较两个数组大小
// @Date:2022/07/25 22:36:14

func Intersection(s1 []int, s2 []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)

	for _, v1 := range s1 {
		set[v1] = true
	}

	for _, v2 := range s2 {
		if true_or_false, ok := set[v2]; ok && true_or_false {
			res = append(res, v2)
			set[v2] = false //防止重复输出
		}
	}
	return res
}

// 字符数组 转换 int 数组
func StringToInt(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
