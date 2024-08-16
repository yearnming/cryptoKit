package codec

import (
	"errors"
	"fmt"
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/cryptoKit/codec/encoding"
	"reflect"
	"strings"
)

// GetEncodeFunc1 获取编码函数
func GetEncodeFunc1(name string) (interface{}, error) {
	if fn, ok := encodeDecodeMap[name]; ok {
		switch fn := fn.(type) {
		case func(string) (string, error), func([]byte) ([]byte, error), func(byte) ([]byte, error):
			return fn, nil
		default:
			errMsg := fmt.Sprintf("编码函数签名不匹配: %s", name)
			gologger.Error().Msg(errMsg)
			return nil, fmt.Errorf(errMsg)
		}
	}
	errMsg := fmt.Sprintf("编码函数不存在: %s", name)
	gologger.Error().Msg(errMsg)
	return nil, fmt.Errorf(errMsg)
}

// EncodeAndSplit1 编解码入口函数
func EncodeAndSplit1(input, split1, split2, encodingType string, splitByLine bool, all bool) (string, error) {
	// 确定是否使用 "All" 变体的编码函数
	if all {
		encodingType1 := fmt.Sprintf("%sAll", encodingType)
		if _, err := GetEncodeFunc1(encodingType1); err == nil {
			encodingType = encodingType1
		} else {
			gologger.Error().Msgf("此编码函数不存在all: %s", encodingType1)
		}
	}

	// 获取编码函数
	encodeFunc, err := GetEncodeFunc1(encodingType)
	if err != nil {
		return "", err
	}

	// 按行编码处理
	if splitByLine {
		return EncodeByLine(input, split1, split2, encodingType, all, encodeFunc)
	}

	//// 处理无分隔符的情况
	//if split1 == "" && split2 == "" {
	//	//return encodeDirectly(input, encodeFunc)
	//	directly, err := EncodeDirectly1(input, encodeFunc)
	//	if err != nil {
	//		return "", err
	//	}
	//	return string(directly), nil
	//}
	if strings.HasPrefix(encodingType, "ConvertOr") && strings.HasSuffix(encodingType, "Decode") && split1 == "" {
		split1 = " "
	}
	if strings.HasPrefix(encodingType, "Base") && (strings.HasSuffix(encodingType, "Encode") || strings.HasSuffix(encodingType, "Decode")) && split1 == "" {
		split1 = "\n"
	}
	if strings.HasPrefix(encodingType, "ConvertBase") && split1 == "" {
		split1 = " "
	}
	// 按第一个分隔符分割输入字符串，并处理特定编码情况
	parts := TrimSplit1(input, split1)
	if encodingType == "ConvertBase" && !splitByLine {
		parts, _ = handleConvertBase(parts)
	}
	if encodingType == "HTMLDecode" && split1 == "" {
		parts = []string{input}
	}
	//if encodingType == "ConvertOrHexByteEncode2" || encodingType == "ConvertOrDecimalByteEncode" || encodingType == "ConvertOrOctalByteEncode" || encodingType == "ConvertOrBinaryByteEncode" {
	//	return handleConvertOrHexByte(input, split2, encodeFunc)
	//}
	if strings.HasSuffix(encodingType, "ByteEncode") {
		return handleConvertOrHexByte(input, split2, encodeFunc)
	}
	if encodingType == "UnicodeDecode" {
		return handleUnicodeDecode(input, split1, split2, encodeFunc)
	}
	if encodingType == "UnicodeEncode" {
		return handleUnicodeEncode(input, split1, split2, encodeFunc)
	}
	// 编码每个部分并用第二个分隔符连接
	return EncodeAndJoin(parts, split2, encodeFunc)
}

// EncodeByLine 按行处理编解码的逻辑函数
func EncodeByLine(input, split1, split2, encodingType string, all bool, encodeFunc interface{}) (string, error) {
	lines := strings.Split(input, "\n")
	encodedLines := make([]string, len(lines))
	if encodingType == "ConvertBase" {
		partss := strings.SplitN(lines[0], "$+++$", 2)
		if len(partss) != 2 {
			return "", fmt.Errorf("输入格式不正确")
		}
		partss[0] = partss[0] + "$+++$"
		for i := 1; i < len(lines); i++ {
			lines[i] = partss[0] + lines[i]
		}
	}
	for i, line := range lines {
		//if encodingType == "ConvertBase" {
		//	partss := strings.SplitN(line, "$+++$", 2)
		//	if len(partss) != 2 {
		//		return "", fmt.Errorf("输入格式不正确")
		//	}
		//	partss[0] = partss[0] + "$+++$"
		//	for i := 1; i < len(lines); i++ {
		//		lines[i] = partss[0] + lines[i]
		//	}
		//}
		encodedLine, err := EncodeAndSplit1(line, split1, split2, encodingType, false, all)
		if err != nil {
			return "", err
		}
		encodedLines[i] = encodedLine
	}
	return strings.Join(encodedLines, "\n"), nil
}

// EncodeDirectly 编解码函数的具体执行 返回值是字符串
func EncodeDirectly(input string, encodeFunc interface{}) (string, error) {
	switch fn := encodeFunc.(type) {
	case func(string) (string, error):
		return fn(input)
	case func([]byte) ([]byte, error):
		result, err := fn([]byte(input))
		if err != nil {
			return "", err
		}
		return string(result), nil
	default:
		return "", fmt.Errorf("未知的编码函数类型")
	}
}

// EncodeDirectly1 编解码函数的具体执行 返回值是字节切片
func EncodeDirectly1(input interface{}, encodeFunc interface{}) ([]byte, error) {
	// 将输入转换为 []byte
	var inputData string
	var inputbyte1 []byte
	var inputbyte byte
	switch v := input.(type) {
	case string:
		inputData = v
	case []byte:
		inputbyte1 = v
	case byte:
		inputbyte = v
	default:
		return nil, fmt.Errorf("不支持的输入类型")
	}

	// 根据 encodeFunc 的类型调用相应的编码函数
	switch f := encodeFunc.(type) {
	case func(string) (string, error):
		encodedStr, err := f(inputData)
		if err != nil {
			return nil, err
		}
		return []byte(encodedStr), nil
	case func([]byte) ([]byte, error):
		// 确保 inputbyte1 不为空
		if len(inputbyte1) == 0 {
			return nil, fmt.Errorf("输入类型为 []byte，但实际输入为空")
		}
		return f(inputbyte1)
	case func(byte) ([]byte, error): // 添加对单个字节处理函数的支持
		return f(inputbyte)
	default:
		return nil, fmt.Errorf("未知的编码函数类型")
	}
}

// 编码函数的选择器
func EncodeDirectly2(input interface{}, encodeFunc interface{}) ([]byte, error) {
	inputValue := reflect.ValueOf(input)
	funcValue := reflect.ValueOf(encodeFunc)
	funcType := funcValue.Type()

	// 确保 encodeFunc 是一个函数
	if funcType.Kind() != reflect.Func {
		return nil, errors.New("encodeFunc 必须是一个函数")
	}

	// 准备函数参数
	var args []reflect.Value

	switch inputValue.Kind() {
	case reflect.String:
		if funcType.NumIn() == 1 && funcType.In(0).Kind() == reflect.String {
			args = append(args, inputValue)
		} else {
			return nil, errors.New("函数参数类型不匹配: 预期 string")
		}
	case reflect.Slice:
		if funcType.NumIn() == 1 && funcType.In(0).Kind() == reflect.Slice && funcType.In(0).Elem().Kind() == reflect.Uint8 {
			args = append(args, inputValue)
		} else {
			return nil, errors.New("函数参数类型不匹配: 预期 []byte")
		}
	case reflect.Uint8:
		if funcType.NumIn() == 1 && funcType.In(0).Kind() == reflect.Uint8 {
			args = append(args, inputValue)
		} else if funcType.NumIn() == 1 && funcType.In(0).Kind() == reflect.Slice && funcType.In(0).Elem().Kind() == reflect.Uint8 {
			args = append(args, reflect.ValueOf([]byte{input.(byte)}))
		} else {
			return nil, errors.New("函数参数类型不匹配: 预期 byte")
		}
	default:
		return nil, errors.New("不支持的输入类型")
	}

	// 调用函数
	results := funcValue.Call(args)

	// 确保返回值数量和类型正确
	if len(results) != 2 {
		return nil, errors.New("函数返回值数量不正确")
	}
	if err, ok := results[1].Interface().(error); ok && err != nil {
		return nil, err
	}
	if result, ok := results[0].Interface().([]byte); ok {
		return result, nil
	}
	return nil, errors.New("函数返回值类型不正确")
}

// 处理特定编码情况 这里是转换进制
func handleConvertBase(parts []string) ([]string, error) {
	partss := strings.SplitN(parts[0], "$+++$", 2)
	if len(partss) != 2 {
		return nil, fmt.Errorf("输入格式不正确")
	}
	partss[0] = partss[0] + "$+++$"
	for i := 1; i < len(parts); i++ {
		parts[i] = partss[0] + parts[i]
	}
	return parts, nil
}

// handleUnicodeDecode 处理特定编码情况 这里是UnicodeDecode
func handleUnicodeDecode(input string, split1 string, split2 string, encodeFunc interface{}) (string, error) {
	if split1 == "" {
		data, _ := EncodeDirectly1(input, encodeFunc)
		return string(data), nil
	} else {
		input1, err := encoding.ReplaceCustomUnicodePrefix(input, split1, split2)
		if err != nil {
			fmt.Println("Error replacing custom Unicode prefix:", err)
		}
		data, _ := EncodeDirectly1(input1, encodeFunc)

		return string(data), nil
	}
}

// handleUnicodeEncode 处理特定编码情况 这里是UnicodeEncode
func handleUnicodeEncode(input string, split1 string, split2 string, encodeFunc interface{}) (string, error) {
	if split2 == "" {
		data, _ := EncodeDirectly1(input, encodeFunc)
		return string(data), nil
	} else {
		data, _ := EncodeDirectly1(input, encodeFunc)
		input1, err := encoding.ReplaceCustomUnicodePrefix(string(data), split1, split2)
		if err != nil {
			fmt.Println("Error replacing custom Unicode prefix:", err)
		}
		return input1, nil
	}
}

// handleConvertOrHexByte 处理特定编码情况 这里是字符串转二进制等
func handleConvertOrHexByte(input string, split2 string, encodeFunc interface{}) (string, error) {
	data := []byte(input)
	//databyte := make([]byte, len(data))
	datastr := make([]string, len(data))
	for _, b := range data {
		// 将每个字节转换为十六进制表示，并写入到缓冲区
		//databyte1, _ := encoding.ConvertOrHexByteEncode2(b)
		databyte1, _ := EncodeDirectly1(b, encodeFunc)
		//databyte = append(databyte, databyte1...)
		datastr = append(datastr, split2+string(databyte1))
	}
	return strings.TrimSpace(strings.Join(datastr, "")), nil
}

// EncodeAndJoin 编解码 并且 每个部分并用第二个分隔符连接
func EncodeAndJoin(parts []string, split2 string, encodeFunc interface{}) (string, error) {
	encodedParts := make([]string, len(parts))
	for i, part := range parts {
		//encoded, err := encodeDirectly(part, encodeFunc)

		encoded, err := EncodeDirectly1(part, encodeFunc)
		if err != nil {
			gologger.Error().Msgf("无法编码部分: %s", err)
			encodedParts[i] = fmt.Sprintf("编码失败: %s", err)
		} else {
			encodedParts[i] = string(encoded)
		}
	}

	if split2 == "" {
		return AddPrefixAndJoin1(encodedParts, ""), nil
	}
	return AddPrefixAndJoin1(encodedParts, split2), nil
}

// TrimSplit1 去除首尾空白换行符 然后按照分隔符进行分割字符串
func TrimSplit1(input, split string) []string {
	trimmedInput := strings.TrimSpace(input)
	parts := strings.Split(trimmedInput, split)
	i, j := 0, len(parts)-1
	for ; i < len(parts) && parts[i] == ""; i++ {
	}
	for ; j >= 0 && parts[j] == ""; j-- {
	}
	if i <= j {
		return parts[i : j+1]
	}
	return []string{}
}

// AddPrefixAndJoin1 添加后缀分隔符
func AddPrefixAndJoin1(encodedParts []string, prefix string) string {
	for i, part := range encodedParts {
		encodedParts[i] = prefix + part
	}
	return strings.Join(encodedParts, "")
}
