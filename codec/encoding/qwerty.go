package encoding

import (
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	qwerty   = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
)

// createQWERTYMap 创建字母到QWERTY布局的映射
func createQWERTYMap() map[rune]rune {
	qwertyMap := make(map[rune]rune)
	for i, r := range qwerty {
		var alphabetR rune
		if i < len(alphabet) {
			alphabetR = rune(alphabet[i])
		}
		qwertyMap[alphabetR] = r
	}
	return qwertyMap
}

// 创建QWERTY解码映射
func createQWERTYReverseMap() map[rune]rune {
	qwertyReverseMap := make(map[rune]rune)
	for i, r := range alphabet {
		qwertyReverseMap[rune(qwerty[i])] = r
	}
	return qwertyReverseMap
}

// QWERTYEncode 将输入字符串编码为QWERTY布局
func QWERTYEncode(input string) (string, error) {
	qwertyMap := createQWERTYMap()
	var output strings.Builder
	for _, r := range input {
		if mapped, ok := qwertyMap[r]; ok {
			output.WriteRune(mapped)
		} else {
			output.WriteRune('?') // 如果字符不在映射中，用问号表示
		}
	}
	return output.String(), nil
}

// QWERTYDecode 将QWERTY编码的字符串解码回原始形式
func QWERTYDecode(input string) (string, error) {
	qwertyReverseMap := createQWERTYReverseMap()
	var output strings.Builder
	for _, r := range input {
		if mapped, ok := qwertyReverseMap[r]; ok {
			output.WriteRune(mapped)
		} else {
			output.WriteRune('?') // 如果字符不在映射中，用问号表示
		}
	}
	return output.String(), nil
}
