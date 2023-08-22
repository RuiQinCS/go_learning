package format_string

import (
	"fmt"
	"strconv"
	"strings"
)

// 获取格式化字符串
func GetFormatString(str string) string {
	idx := strings.Index(str, ".")
	if idx == -1 { // 无小数点
		_, err := strconv.Atoi(str)
		if err != nil {
			return "%v"
		}
		return "%.0f"
	}

	return "%." + fmt.Sprint(len(str[idx+1:])) + "f" //fmt.Sprintf("%.%vf", len(str[idx+1:]))
}
