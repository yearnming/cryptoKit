package encoding

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// UnicodeEncode 将字符串编码为 Unicode 格式 (\uXXXX)
func UnicodeEncode(input string) (string, error) {
	output := ""
	for _, r := range input {
		output += fmt.Sprintf("\\u%04x", r)
	}
	return output, nil
}

// UnicodeDecode 将 Unicode 字符串 (\uXXXX) 解码为普通字符串
func UnicodeDecode1(input string) (string, error) {
	var output string
	for len(input) > 0 {
		if input[0:2] == "\\u" {
			var r rune
			_, err := fmt.Sscanf(input[2:6], "%04x", &r)
			if err != nil {
				return "", err
			}
			output += string(r)
			input = input[6:]
		} else {
			output += string(input[0])
			input = input[1:]
		}
	}
	return output, nil
}

func UnicodeDecode(input string) (string, error) {
	var outputBuilder strings.Builder

	for len(input) > 0 {
		if strings.HasPrefix(input, "\\u") {
			// 处理以\u开头的Unicode转义序列
			if len(input) < 6 {
				return "", fmt.Errorf("invalid unicode escape sequence at the end of the string")
			}
			r, err := parseUnicodeEscape(input[2:6])
			if err != nil {
				return "", err
			}
			outputBuilder.WriteRune(r)
			input = input[6:] // 跳过已处理的转义序列
		} else {
			// 处理普通字符或中文字符
			r, size := utf8.DecodeRuneInString(input)
			outputBuilder.WriteRune(r)
			input = input[size:] // 跳过已处理的字符
		}
	}

	return outputBuilder.String(), nil
}

// parseUnicodeEscape 解析Unicode转义序列
func parseUnicodeEscape(hexStr string) (rune, error) {
	var r rune
	n, err := fmt.Sscanf(hexStr, "%04x", &r)
	if err != nil || n != 1 {
		return 0, fmt.Errorf("failed to parse unicode escape sequence: %s", hexStr)
	}
	return r, nil
}

func ReplaceCustomUnicodePrefix(input, split1 string, split2 string) (string, error) {
	if split1 == "" {
		split1 = "\\u"
	}
	if split2 == "" {
		split2 = "\\u"
	}
	// 构建正则表达式，匹配自定义前缀后跟4个十六进制数字
	pattern := regexp.QuoteMeta(split1) + `([0-9a-fA-F]{4})`

	re := regexp.MustCompile(pattern)

	// 使用正则表达式替换所有匹配的Unicode前缀
	input = re.ReplaceAllString(input, fmt.Sprintf("%s$1", split2))

	return input, nil
}
