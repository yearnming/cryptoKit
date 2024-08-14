package codec

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/cryptoKit/codec/encoding"
	"strings"
)

// Functions exposed by the codec package
var (
	URLDecode       = encoding.URLDecode
	URLEncode       = encoding.URLEncode
	UnicodeDecode   = encoding.UnicodeDecode
	UnicodeEncode   = encoding.UnicodeEncode
	HTMLEncode      = encoding.HTML16Encode
	HTMLDecode      = encoding.HTMLDecode
	ShellcodeDecode = encoding.ShellcodeDecode
	ShellcodeEncode = encoding.ShellcodeEncode
	QWERTYDecode    = encoding.QWERTYDecode
	QWERTYEncode    = encoding.QWERTYEncode
	Base16Encode    = encoding.Base16Encode
	Base16Decode    = encoding.Base16Decode
	Base32Encode    = encoding.Base32Encode
	Base32Decode    = encoding.Base32Decode
	Base45Encode    = encoding.Base45Encode
	Base45Decode    = encoding.Base45Decode
	Base58Encode    = encoding.Base58Encode
	Base58Decode    = encoding.Base58Decode
	Base62Encode    = encoding.Base62Encode
	Base62Decode    = encoding.Base62Decode
	Base64Encode    = encoding.Base64Encode
	Base64Decode    = encoding.Base64Decode
	Base64URLEncode = encoding.Base64URLEncode
	Base64URLDecode = encoding.Base64URLDecode
	Base85Encode    = encoding.Base85Encode
	Base85Decode    = encoding.Base85Decode
	Base91Encode    = encoding.Base91Encode
	Base91Decode    = encoding.Base91Decode
	Base100Encode   = encoding.Base100Encode
	Base100Decode   = encoding.Base100Decode
	EscapeEncode    = encoding.EscapeEncode
	EscapeDecode    = encoding.EscapeDecode
	//ASCIIEncodeToBinary    = encoding.ASCIIEncodeToBinaryEncode
	//ASCIIDecodeFromBinary  = encoding.ASCIIDecodeFromBinaryDecode
	//ASCIIEncodeToOctal     = encoding.ASCIIEncodeToOctalEncode
	//ASCIIDecodeFromOctal   = encoding.ASCIIDecodeFromOctalDecode
	//ASCIIEncodeToDecimal   = encoding.ASCIIEncodeToDecimalEncode
	//ASCIIDecodeFromDecimal = encoding.ASCIIDecodeFromDecimalDecode
	//ASCIIEncodeToHex       = encoding.ASCIIEncodeToHexEncode
	//ASCIIDecodeFromHex     = encoding.ASCIIDecodeFromHexDecode
	//ConvertToBinary        = encoding.ConvertToBinaryEncode
	//ConvertFromBinary      = encoding.ConvertFromBinaryDecode
	//ConvertToOctal         = encoding.ConvertToOctalEncode
	//ConvertFromOctal       = encoding.ConvertFromOctalDecode
	//ConvertToDecimal       = encoding.ConvertToDecimalEncode
	//ConvertFromDecimal     = encoding.ConvertFromDecimalDecode
	//ConvertToHex           = encoding.ConvertToHexEncode
	//ConvertFromHex         = encoding.ConvertFromHexDecode
)

var encodeDecodeMap = map[string]interface{}{
	"URLEncode":        encoding.URLEncode,
	"URLEncodeAll":     encoding.URLEncodeAll,
	"URLDecode":        encoding.URLDecode,
	"UnicodeEncode":    encoding.UnicodeEncode, //unicode (\uXXXX)
	"UnicodeDecode":    encoding.UnicodeDecode,
	"EscapeEncode":     encoding.EscapeEncode, //unicode  %uXXXX
	"EscapeDecode":     encoding.EscapeDecode,
	"HTML10Encode":     encoding.HTML10Encode,
	"HTML16Encode":     encoding.HTML16Encode,
	"HTMLruneEncode":   encoding.HTMLruneEncode,
	"HTMLrune10Encode": encoding.HTMLrune10Encode,
	"HTMLrune16Encode": encoding.HTMLrune16Encode,
	"HTMLDecode":       encoding.HTMLDecode,
	//"ShellcodeEncode":   encoding.ShellcodeEncode,
	//"ShellcodeDecode":   encoding.ShellcodeDecode,
	"QWERTYEncode":      encoding.QWERTYEncode,
	"QWERTYDecode":      encoding.QWERTYDecode,
	"StringOrHexEncode": encoding.StringOrHexEncode,
	"StringOrHexDecode": encoding.StringOrHexDecode,
	"Base16Encode":      encoding.Base16Encode,
	"Base16Decode":      encoding.Base16Decode,
	"Base32Encode":      encoding.Base32Encode,
	"Base32Decode":      encoding.Base32Decode,
	"Base45Encode":      encoding.Base45Encode,
	"Base45Decode":      encoding.Base45Decode,
	"Base58Encode":      encoding.Base58Encode,
	"Base58Decode":      encoding.Base58Decode,
	"Base62Encode":      encoding.Base62Encode,
	"Base62Decode":      encoding.Base62Decode,
	"Base64Encode":      encoding.Base64Encode,
	"Base64Decode":      encoding.Base64Decode,
	"Base64URLEncode":   encoding.Base64URLEncode,
	"Base64URLDecode":   encoding.Base64URLDecode,
	"Base85Encode":      encoding.Base85Encode,
	"Base85Decode":      encoding.Base85Decode,
	"Base91Encode":      encoding.Base91Encode,
	"Base91Decode":      encoding.Base91Decode,
	"Base100Encode":     encoding.Base100Encode,
	"Base100Decode":     encoding.Base100Decode,
	//"ASCIIOrBinaryEncode":    encoding.ASCIIOrBinaryEncode,
	//"ASCIIOrBinaryDecode":    encoding.ASCIIOrBinaryDecode,
	//"ASCIIOrOctalEncode":     encoding.ASCIIOrOctalEncode,
	//"ASCIIOrOctalDecode":     encoding.ASCIIOrOctalDecode,
	//"ASCIIOrDecimalEncode":   encoding.ASCIIOrDecimalEncode,
	//"ASCIIOrDecimalDecode":   encoding.ASCIIOrDecimalDecode,
	//"ASCIIOrHexEncode":       encoding.ASCIIOrHexEncode,
	//"ASCIIOrHexDecode":       encoding.ASCIIOrHexDecode,
	"ConvertOrBinaryByteEncode":  encoding.ConvertOrBinaryByteEncode,
	"ConvertOrBinaryDecode":      encoding.ConvertOrBinaryDecode,
	"ConvertOrOctalByteEncode":   encoding.ConvertOrOctalByteEncode,
	"ConvertOrOctalDecode":       encoding.ConvertOrOctalDecode,
	"ConvertOrDecimalByteEncode": encoding.ConvertOrDecimalByteEncode,
	"ConvertOrDecimalDecode":     encoding.ConvertOrDecimalDecode,
	"ConvertOrHexByteEncode":     encoding.ConvertOrHexByteEncode,
	"ConvertOrHexDecode":         encoding.ConvertOrHexDecode,

	"ConvertBase": encoding.ConvertBase,
}

// EncodeOrDecode 定义一个函数，根据编码方式和操作类型进行编解码
//func EncodeOrDecode(encodingType, input string) (string, error) {
//	// 从映射中获取函数
//	fn, exists := encodeDecodeMap[encodingType]
//	if !exists {
//		return "", fmt.Errorf("不支持的编码类型: %s", encodingType)
//	}
//	// 调用函数并返回结果
//	return fn(input)
//}

// 获取编码函数
//func getEncodeFunc(name string) (func(string) (string, error), error) {
//	if fn, ok := encodeDecodeMap[name]; ok {
//		return fn.(func(string) (string, error)), nil
//	}
//	gologger.Error().Msgf("编码函数不存在: %s", name)
//	return nil, fmt.Errorf("编码函数不存在: " + name)
//}

// getEncodeFunc 获取编码函数，支持不同的返回类型
func getEncodeFunc(name string) (interface{}, error) {
	if fn, ok := encodeDecodeMap[name]; ok {
		switch fn := fn.(type) {
		case func(string) (string, error):
			return fn, nil
		case func([]byte) ([]byte, error):
			return fn, nil
		default:
			gologger.Error().Msgf("编码函数签名不匹配: %s", name)
			return nil, fmt.Errorf("编码函数签名不匹配: " + name)
		}
	}
	gologger.Error().Msgf("编码函数不存在: %s", name)
	return nil, fmt.Errorf("编码函数不存在: " + name)
}

// EncodeAndSplit 先分割字符串，再进行编码，最后再分割编码后的结果
func EncodeAndSplit(input, split1, split2, encodingType string, splitByLine bool, all bool) (string, error) {

	var encodingType1 string
	if all {
		encodingType1 = fmt.Sprintf("%sAll", encodingType)
		_, err := getEncodeFunc(encodingType1)
		if err != nil {
			gologger.Error().Msgf("此编码函数不存在all: %s", encodingType1)
		} else {
			encodingType = encodingType1
		}
	}
	// 获取编码函数
	encodeFunc, err := getEncodeFunc(encodingType)
	if err != nil {
		return "", err
	}

	// 如果按行处理
	if splitByLine {
		lines := strings.Split(input, "\n")
		encodedLines := make([]string, len(lines))
		for i, line := range lines {
			if encodingType == "ConvertBase" {
				partss := strings.SplitN(input, "$+++$", 2)
				if len(partss) != 2 {
					return "", fmt.Errorf("输入格式不正确")
				}
				// 要拼接的字符串
				partss[0] = partss[0] + "$+++$"

				// 第一个元素不拼接
				// 从第二个元素开始使用 + 操作符拼接字符串
				for i := 1; i < len(lines); i++ {
					lines[i] = partss[0] + lines[i]
				}

			}
			encodedLine, err := EncodeAndSplit(line, split1, split2, encodingType, false, all) // 对每行使用默认的分隔符处理
			if err != nil {
				return "", err
			}
			encodedLines[i] = encodedLine
		}
		return strings.Join(encodedLines, "\n"), nil
	}

	// 如果没有分隔符，不做分割处理
	if split1 == "" && split2 == "" {
		// 假设我们有一个需要编码的字符串
		//inputStr := "Hello, World!"
		switch fn := encodeFunc.(type) {
		case func(string) (string, error):
			encoded, err := fn(input)
			if err != nil {
				fmt.Println("编码出错:", err)
				return "", err
			}
			return encoded, nil
		case func([]byte) ([]byte, error):
			result, err := fn([]byte(input))
			if err != nil {
				fmt.Println("编码出错:", err)
				return "", err
			} else {
				//fmt.Println("结果:", string(result))
				return string(result), nil
			}
		default:
			fmt.Println("未知的编码函数类型")
		}
		//return encodeFunc(input)
	}
	// 按第一个分隔符分割输入字符串
	parts := TrimSplit(input, split1)

	if encodingType == "ConvertBase" && splitByLine == false {
		partss := strings.SplitN(input, "$+++$", 2)
		if len(partss) != 2 {
			return "", fmt.Errorf("输入格式不正确")
		}
		//parts := []string{"first", "apple", "banana", "cherry"}

		// 要拼接的字符串
		partss[0] = partss[0] + "$+++$"

		// 第一个元素不拼接
		// 从第二个元素开始使用 + 操作符拼接字符串
		for i := 1; i < len(parts); i++ {
			parts[i] = partss[0] + parts[i]
		}

	}

	// 对每个部分进行编码
	encodedParts := make([]string, len(parts))
	for i, part := range parts {
		//encoded, err := encodeFunc(part)
		var encoded string
		switch fn := encodeFunc.(type) {
		case func(string) (string, error):
			encoded, err = fn(part)
			if err != nil {
				fmt.Println("编码出错:", err)
				return "", err
			}
			return encoded, nil
		case func([]byte) ([]byte, error):
			result, err := fn([]byte(part))
			if err != nil {
				fmt.Println("编码出错:", err)
				//return "", err
			} else {
				//fmt.Println("结果:", string(result))
				encoded = string(result)
			}
		default:
			fmt.Println("未知的编码函数类型")
		}
		if err != nil {
			gologger.Error().Msgf("无法编码部分: %s", err)
			encodedParts[i] = fmt.Sprintf("编码失败: %s", err)
			//return "", err
		} else {
			encodedParts[i] = encoded
		}

	}

	// 如果没有指定第二个分隔符，则按行连接
	if split2 == "" {
		return AddPrefixAndJoin(encodedParts, ""), nil
	}

	// 按第二个分隔符连接编码后的部分
	return AddPrefixAndJoin(encodedParts, split2), nil
}

// TrimSplit 函数分割字符串并去除结果切片首尾的空字符串
func TrimSplit(input, split string) []string {
	// 去除输入字符串首尾的空白
	trimmedInput := strings.TrimSpace(input)
	// 使用 strings.Split 分割字符串
	parts := strings.Split(trimmedInput, split)

	// 去除首尾的空字符串
	// 初始化索引 i
	i := 0
	// 从左到右找到第一个非空字符串的索引
	for ; i < len(parts) && parts[i] == ""; i++ {
	}

	// 初始化索引 j
	j := len(parts) - 1
	// 从右到左找到第一个非空字符串的索引
	for ; j >= 0 && parts[j] == ""; j-- {
	}

	// 如果 i 和 j 相遇或经过了非空字符串，截取并返回结果
	if i <= j {
		return parts[i : j+1]
	}
	// 如果没有非空字符串，返回空切片
	return []string{}
}

func AddPrefixAndJoin(encodedParts []string, prefix string) string {
	for i, part := range encodedParts {
		encodedParts[i] = prefix + part
	}
	return strings.Join(encodedParts, "")
}
