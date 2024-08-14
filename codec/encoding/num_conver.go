package encoding

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertBase converts a number from one base to another
//func ConvertBase(number string, fromBase, toBase int) (string, error) {
//	// 将输入的数字转换为十进制
//	n, err := strconv.ParseInt(number, fromBase, 64)
//	if err != nil {
//		return "", err
//	}
//
//	// 将十进制转换为目标进制
//	return strconv.FormatInt(n, toBase), nil
//}

func ConvertBase(input string) (string, error) {
	// 假设输入格式为 "{{fromBase,toBase}}+number"
	parts := strings.SplitN(input, "$+++$", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("输入格式不正确")
	}

	baseInfo := parts[0]
	number := parts[1]

	// 去掉前后的大括号
	baseInfo = strings.Trim(baseInfo, "{}")

	// 分割 fromBase 和 toBase
	bases := strings.Split(baseInfo, ",")
	if len(bases) != 2 {
		return "", fmt.Errorf("基数信息格式不正确")
	}

	// 解析 fromBase 和 toBase
	fromBase, err := strconv.Atoi(bases[0])
	if err != nil {
		return "", fmt.Errorf("解析 fromBase 时出错: %v", err)
	}

	toBase, err := strconv.Atoi(bases[1])
	if err != nil {
		return "", fmt.Errorf("解析 toBase 时出错: %v", err)
	}

	// 将 number 从 fromBase 转换到 toBase
	n, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		return "", fmt.Errorf("解析 number 时出错: %v", err)
	}

	return strconv.FormatInt(n, toBase), nil
}
