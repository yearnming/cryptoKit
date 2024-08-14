package codec

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/cryptoKit/codec/encoding"
	"html"
	"testing"
)

func TestURLEncodeDecode(t *testing.T) {
	// 定义一个包含ASCII 32-126所有字符的字符串
	original := ""
	for i := 32; i <= 126; i++ {
		original += string(rune(i))
	}

	//original := "www.gouguoyin.cn?sex=男&age=18 <div>Hello, 世界!</div"
	encoded, _ := encoding.URLEncodeAll(original)
	decoded, err := encoding.URLDecode(encoded)
	gologger.Info().Msgf("URL: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)

	// 使用 URL 编码，按行分割
	original12, err := EncodeAndSplit(original, "", "", "UnicodeEncode", false, false)
	if err != nil {
		fmt.Println("URL 编码出错:", err)
	} else {
		gologger.Info().Msgf("URL: Encoded: [%s]", original12)
	}
	original123, err := EncodeAndSplit(original, "", "", "UnicodeEncode", false, true)
	if err != nil {
		fmt.Println("URL 编码出错:", err)
	} else {
		gologger.Info().Msgf("URL: EncodedAll: [%s]", original123)
	}
	if err != nil || decoded != original {
		t.Errorf("Failed URL encode/decode: got %s, want %s", decoded, original)
	}

}

//	func ReplaceCustomUnicodePrefix(input, customPrefix string) (string, error) {
//		// 构建正则表达式，匹配自定义前缀后跟4个十六进制数字
//		pattern := regexp.QuoteMeta(customPrefix) + `([0-9a-fA-F]{4})`
//
//		re := regexp.MustCompile(pattern)
//
//		// 使用正则表达式替换所有匹配的Unicode前缀
//		input = re.ReplaceAllString(input, `\\u$1`)
//
//		return input, nil
//	}
func TestUnicodeEncodeDecode(t *testing.T) {
	//input := "123 asd \\u0048 asd%u0065 123U+006cU+006co u+002c\\u0020世界！"
	//customPrefix := "U+" // 自定义Unicode前缀

	//processed, err := encoding.ReplaceCustomUnicodePrefix(input, customPrefix)
	//if err != nil {
	//	fmt.Println("Error replacing custom Unicode prefix:", err)
	//	return
	//}
	//gologger.Info().Msgf("Encoded: [%s]", processed)
	original := "Hello, 世界!"
	//original1 := "a \\u0048\\u0065\\u006c\\u006c\\u006f\\u002c\\u0020\\u4e16\\u754c\\u0021"
	//encoded, _ := UnicodeEncode(original)
	//decoded, _ := UnicodeDecode(encoded)
	//gologger.Info().Msgf("Unicode: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)
	//original11 := "U+0048\\u0065\\u006c\\u006c\\u006f\\u002c\\u0020\\u4e16\\u754cU+0021"
	customPrefix := "U+" // 自定义Unicode前缀
	original1, err := EncodeAndSplit1(original, "", customPrefix, "UnicodeEncode", false, false)
	original123, err := EncodeAndSplit1(original1, customPrefix, "", "UnicodeDecode", false, false)
	if err != nil {
		fmt.Println("ConvertOrHexEncode 编码出错:", err)
	} else {
		gologger.Info().Msgf("ConvertOrHexEncode: [%s] [%s]", original1, original123)
	}

}

func TestHTMLEncodeDecode(t *testing.T) {
	//original := "<div>Hello, 世界!`\\</div>"
	original := ""
	for i := 32; i <= 126; i++ {
		original += string(rune(i))
	}
	encoded, _ := encoding.HTMLrune16Encode(original)
	decoded, _ := HTMLDecode(encoded)
	gologger.Info().Msgf("HTML: Encoded: [%s]", encoded)
	gologger.Info().Msgf("HTML: Decoded: [%s]", decoded)
	escapedString := html.EscapeString(original)
	unescapedString := html.UnescapeString(escapedString)
	gologger.Info().Msgf("HTML: Encoded: [%s] --- Decoded: [%s]\n", escapedString, unescapedString)
	if decoded != original {
		t.Errorf("Failed HTML encode/decode: got %s, want %s", decoded, original)
	}
	original1233, err := EncodeAndSplit1("{&#34;code&#34;:&#34;AccessDenied&#34;,&#34;message&#34;:&#34;Access Denied.&#34;,&#34;requestId&#34", "", "", "HTMLDecode", false, false)
	if err != nil {
		fmt.Println("ConvertOrHexEncode 编码出错:", err)
	} else {
		gologger.Info().Msgf("HTMLDecode: [%s]", original1233)
	}
	//gologger.Info().Msgf("HTML: Decoded: [%s]\n", html.UnescapeString("{&#34;code&#34;:&#34;AccessDenied&#34;,&#34;message&#34;:&#34;Access Denied.&#34;,&#34;requestId&#34"))
}

//
//func TestASCIIEncodeDecode(t *testing.T) {
//	original := "Hello, 橙子!"
//	encoded2, _ := ConvertToBinary(original)
//	decoded2, err := ConvertFromBinary(encoded2)
//	encoded8, _ := ConvertToOctal(original)
//	decoded8, err := ConvertFromOctal(encoded8)
//	encoded10, _ := ConvertToDecimal(original)
//	decoded10, err := ConvertFromDecimal(encoded10)
//	encoded16, _ := ConvertToHex(original)
//	decoded16, err := ConvertFromHex(encoded16)
//	gologger.Info().Msgf("ASCII二进制: Encoded: [%s] --- Decoded: [%s]\n", encoded2, decoded2)
//	gologger.Info().Msgf("ASCII八进制: Encoded: [%s] --- Decoded: [%s]\n", encoded8, decoded8)
//	gologger.Info().Msgf("ASCII十进制: Encoded: [%s] --- Decoded: [%s]\n", encoded10, decoded10)
//	gologger.Info().Msgf("ASCII十六进制: Encoded: [%s] --- Decoded: [%s]\n", encoded16, decoded16)
//	if err != nil || decoded2 != original {
//		t.Errorf("Failed ASCII encode/decode: got %s, want %s", decoded2, original)
//	}
//}

// 进制转换
func TestConvertBase(t *testing.T) {
	//original2 := "42"
	//original123, err := EncodeAndSplit(original2, "10", "12", "ConvertBase", false, false)
	//if err != nil {
	//	fmt.Println("ConvertOrHexEncode 编码出错:", err)
	//} else {
	//	gologger.Info().Msgf("ConvertOrHexEncode: [%s]", original123)
	//}
	// 示例输入
	input := "{{10,16}}$+++$42"
	result, err := encoding.ConvertBase(input)
	if err != nil {
		fmt.Println("转换出错:", err)
		return
	}
	fmt.Println("转换结果:", result)
	original2 := "{{16,10}}$+++$69 64 3a\n69 6e 66\n69 64 3a\n69 6e 66"
	original123, err := EncodeAndSplit1(original2, " ", " ", "ConvertBase", true, false)
	if err != nil {
		fmt.Println("ConvertOrHexEncode 编码出错:", err)
	} else {
		gologger.Info().Msgf("ConvertOrHexEncode: [%s]", original123)
	}
	// 进制转换示例
	//binaryToDecimal, _ := encoding.ConvertBase("101010", 2, 10)
	//gologger.Info().Msgf("101010 二进制到十进制: %s", binaryToDecimal)
	//
	//binaryToHex, _ := encoding.ConvertBase("101010", 2, 16)
	//gologger.Info().Msgf("101010 二进制到十六进制: %s", binaryToHex)
	//
	//decimalToBinary, _ := encoding.ConvertBase("42", 10, 2)
	//gologger.Info().Msgf("42 十进制到二进制: %s", decimalToBinary)
	//
	//decimalToHex, _ := encoding.ConvertBase("42", 10, 16)
	//gologger.Info().Msgf("42 十进制到十六进制: %s", decimalToHex)
	//
	//octalToBinary, _ := encoding.ConvertBase("52", 8, 2)
	//gologger.Info().Msgf("52 八进制到二进制: %s", octalToBinary)
	//
	//hexToDecimal, _ := encoding.ConvertBase("2a", 16, 10)
	//gologger.Info().Msgf("2a 十六进制到十进制: %s", hexToDecimal)
}

func TestBase64EncodeDecode(t *testing.T) {
	original := "www.gouguoa yin.cn\nwww.gouguoa yin.cn"

	encoded16, _ := Base16Encode(original)
	decoded16, _ := Base16Decode(encoded16)
	encoded32, _ := Base32Encode(original)
	decoded32, _ := Base32Decode(encoded32)
	encoded45, _ := Base45Encode(original)
	decoded45, _ := Base45Decode(encoded45)
	encoded58, _ := Base58Encode(original)
	decoded58, _ := Base58Decode(encoded58)
	encoded62, _ := Base62Encode(original)
	decoded62, _ := Base62Decode(encoded62)
	encoded64, _ := Base64Encode(original)
	decoded64, _ := Base64Decode(encoded64)
	encoded64URL, _ := Base64URLEncode(original)
	decoded64URL, _ := Base64URLDecode(encoded64URL)
	encoded85, _ := Base85Encode(original)
	decoded85, _ := Base85Decode(encoded85)
	encoded91, _ := Base91Encode(original)
	decoded91, _ := Base91Decode(encoded91)
	encoded100, _ := Base100Encode(original)
	decoded100, _ := Base100Decode(encoded100)
	gologger.Info().Msgf("Base16: Encoded: [%s] --- Decoded: [%s]\n", encoded16, decoded16)
	gologger.Info().Msgf("Base32: Encoded: [%s] --- Decoded: [%s]\n", encoded32, decoded32)
	gologger.Info().Msgf("Base45: Encoded: [%s] --- Decoded: [%s]\n", encoded45, decoded45)
	gologger.Info().Msgf("Base58: Encoded: [%s] --- Decoded: [%s]\n", encoded58, decoded58)
	gologger.Info().Msgf("Base62: Encoded: [%s] --- Decoded: [%s]\n", encoded62, decoded62)
	gologger.Info().Msgf("Base64: Encoded: [%s] --- Decoded: [%s]\n", encoded64, decoded64)
	gologger.Info().Msgf("Base64URL: Encoded: [%s] --- Decoded: [%s]\n", encoded64URL, decoded64URL)
	gologger.Info().Msgf("Base85: Encoded: [%s] --- Decoded: [%s]\n", encoded85, decoded85)
	gologger.Info().Msgf("Base91: Encoded: [%s] --- Decoded: [%s]\n", encoded91, decoded91)
	gologger.Info().Msgf("Base100: Encoded: [%s] --- Decoded: [%s]\n", encoded100, decoded100)
	original123, err := EncodeAndSplit1(original, "", "", "Base64Encode", true, false)
	original1233, err := EncodeAndSplit1(original123, "", "", "Base64Decode", true, false)
	if err != nil {
		fmt.Println("ConvertOrHexEncode 编码出错:", err)
	} else {
		gologger.Info().Msgf("ConvertOrHexEncode: [%s] --- [%s]", original1233, original123)
	}
}

func TestConvertOrHex(t *testing.T) {
	original := "H世界"
	//original2 := "48 e4 b8 96 e7 95 8c\n48 e4 b8 96 e7 95 8c"
	original22 := "48e4b896e7958c\n48e4b896e7958c"
	//data := []byte(original)
	//databyte := make([]byte, len(data))
	//datastr := make([]string, len(data))
	//for _, b := range data {
	//	// 将每个字节转换为十六进制表示，并写入到缓冲区
	//	databyte1, _ := encoding.ConvertOrHexByteEncode2(b)
	//	databyte = append(databyte, databyte1...)
	//	datastr = append(datastr, "%"+string(databyte1))
	//	// 调用函数示例3
	//	result, err := EncodeDirectly2(b, encoding.ConvertOrHexByteEncode2)
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		return
	//	}
	//	fmt.Println("Encoded result 3:", result, "%"+string(result))
	//}
	//gologger.Info().Msgf("databyte[byte]: %v\n", databyte)
	//gologger.Info().Msgf("databyte[string]: %v\n", datastr)

	//original2 := "%%%%%%73%70%6c%69%74%31%3a%20%70%61%6e%65%6c%2e%73%70%6c%69%74%31%2c%%%"
	//parts := TrimSplit1(original2, "%")
	//data := []byte(original)
	//gologger.Info().Msgf("Parts: %v", parts)
	//encoded, _ := encoding.ConvertOrHexByteEncode(data)
	//decoded, err := encoding.ConvertOrHexByteDecode(encoded)
	//gologger.Info().Msgf("Hex: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)
	original1233, err := EncodeAndSplit1(original, "", "%", "ConvertOrDecimalByteEncode", false, false)

	original123, err := EncodeAndSplit1(original22, " ", "", "ConvertOrHexDecode", true, false)
	if err != nil {
		fmt.Println("ConvertOrHexEncode 编码出错:", err)
	} else {
		gologger.Info().Msgf("ConvertOrHexEncode: [%s] --- [%s]", original1233, original123)
	}
}

func TestShellcodeEncodeDecode(t *testing.T) {
	original := "Hello, 世界!"
	encoded, _ := ShellcodeEncode(original)
	decoded, err := ShellcodeDecode(encoded)
	gologger.Info().Msgf("Shellcode: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)

	if err != nil || decoded != original {
		t.Errorf("Failed Shellcode encode/decode: got %s, want %s", decoded, original)
	}
}

func TestQWERTYEncodeDecode(t *testing.T) {
	original := "qwerty"
	encoded, _ := QWERTYEncode(original)
	decoded, _ := QWERTYDecode(encoded)
	gologger.Info().Msgf("QWERTY: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)

	if decoded != original {
		t.Errorf("Failed QWERTY encode/decode: got %s, want %s", decoded, original)
	}
}

func TestEscapeEncodeDecode(t *testing.T) {
	original := "{&#34;code&#34;:&#34;AccessDenied&#34;,&#34;message&#34;:&#34;Access Denied.&#34;,&#34;requestId&#34"
	encoded, _ := EscapeEncode(original)
	decoded, err := EscapeDecode(original)
	gologger.Info().Msgf("Escape: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)

	if err != nil || decoded != original {
		t.Errorf("Failed Escape encode/decode: got %s, want %s", decoded, original)
	}
}

//func TestConvertToBinaryEncodeDecode(t *testing.T) {
//	original := "Hello, 世界!"
//	encoded, _ := ConvertToBinary(original)
//	decoded, err := ConvertFromBinary(encoded)
//	gologger.Info().Msgf("Convert: Encoded: [%s] --- Decoded: [%s]\n", encoded, decoded)
//
//	if err != nil || decoded != original {
//		t.Errorf("Failed Convert encode/decode: got %s, want %s", decoded, original)
//	}
//}

// 测试所有编码和解码函数
// 测试编码和解码函数
//func TestEncodingAndDecoding(t *testing.T) {
//	var testData = "Hello, chengzi! " +
//		"Hello, 橙子！"
//	// 提取map的键并排序
//	keys := make([]string, 0, len(encodeDecodeMap))
//	for name := range encodeDecodeMap {
//		keys = append(keys, name)
//	}
//	sort.Strings(keys) // 根据字典顺序排序键
//
//	for _, name := range keys {
//		encodeFunc := encodeDecodeMap[name]
//		if strings.HasSuffix(name, "Encode") {
//			decodeName := strings.TrimSuffix(name, "Encode") + "Decode"
//			decodeFunc, exists := encodeDecodeMap[decodeName]
//
//			if !exists {
//				gologger.Error().Msgf("未找到解码函数 %s", name)
//				continue
//			}
//
//			encoded, err := encodeFunc(testData)
//			if err != nil {
//				gologger.Error().Msgf("Encoding %s 编码失败: %v", name, err)
//				return
//			}
//
//			decoded, err := decodeFunc(encoded)
//			if err != nil {
//				gologger.Error().Msgf("Decoding %s 解码失败: %v", decodeName, err)
//				return
//			}
//
//			if decoded != testData {
//				gologger.Error().Msgf("解码数据与原始数据不匹配 %s . Got: %s, Want: %s", strings.TrimSuffix(name, "Encode"), decoded, testData)
//			}
//			gologger.Info().Msgf("Testing %s : Encoded: [%s] --- Decoded: [%s]", strings.TrimSuffix(name, "Encode"), encoded, testData)
//		}
//	}
//}

func TestBase64Encoding(t *testing.T) {
	input := "<Hello, 世界!>\n<1+1=2!>"
	split1 := ","
	split2 := ";"

	// 使用 Base64 编码，不按行分割
	encodedResult, err := EncodeAndSplit(input, split1, split2, "Base64Encode", false, false)
	if err != nil {
		fmt.Println("Base64 编码出错:", err)
	} else {
		fmt.Println("Base64 编码结果:", encodedResult)
	}

	// 使用 URL 编码，不按行分割
	encodedResult, err = EncodeAndSplit(input, split1, split2, "URLEncode", false, false)
	if err != nil {
		fmt.Println("URL 编码出错:", err)
	} else {
		fmt.Println("URL 编码结果:", encodedResult)
	}

	// 按行处理，分别对每行进行编码
	input = "<Hello, 世界!>\n<1+1=2!>"

	// 使用 Base64 编码，按行分割
	encodedResult, err = EncodeAndSplit(input, split1, split2, "Base64Encode", true, false)
	if err != nil {
		fmt.Println("Base64 编码出错:", err)
	} else {
		fmt.Println("按行分割 Base64 编码结果:", encodedResult)
	}

	// 使用 URL 编码，按行分割
	encodedResult, err = EncodeAndSplit(input, split1, split2, "URLEncode", true, false)
	if err != nil {
		fmt.Println("URL 编码出错:", err)
	} else {
		fmt.Println("按行分割 URL 编码结果:", encodedResult)
	}
}
