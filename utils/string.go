package utils

import (
	"strings"
)

// @Author: 张佳伟
// @Function:StringToIntArr
// @Description:字符串转  int 数组
// @Date: 2022/07/26 12:06:35

func StringToIntArr(str string) []int {
	strArr := strings.Split(str, ",")
	return StringToInt(strArr)
}
