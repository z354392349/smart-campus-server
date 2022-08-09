package utils

import (
	"fmt"
	"strconv"
)

// @Author: 张佳伟
// @Function:NumToFixed
// @Description: 浮点数保存多少位小数
// @Date:2022/08/09 21:03:03

func NumToFixed(a float64, b float64, num string) (res float64) {
	result, _ := strconv.ParseFloat(fmt.Sprintf("%."+num+"f", a/b), 64) // 保留2位小数
	return result
}
