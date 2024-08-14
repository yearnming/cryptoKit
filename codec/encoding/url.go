package encoding

import (
	"net/url"
	"strings"
)

// URLEncodeAll 对所有特殊字符进行编码
func URLEncodeAll(input string) (string, error) {
	encoded := url.QueryEscape(input)
	replacedString := strings.ReplaceAll(encoded, "-", "%2D")
	replacedString = strings.ReplaceAll(replacedString, ".", "%2E")
	replacedString = strings.ReplaceAll(replacedString, "_", "%5F")
	replacedString = strings.ReplaceAll(replacedString, "~", "%7E")
	replacedString = strings.ReplaceAll(replacedString, "+", "%20")
	// 对 url 字符进行转义编码，输出字符串
	//encoded := dongle.Encode.FromString(input).BySafeURL().ToString() // www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18
	//return encoded, nil
	return replacedString, nil
}

// URLEncode 进行url编码 不对"- . _ ~  "进行编码 默认将空格" "变成"+"
func URLEncode(input string) (string, error) {
	encoded := url.QueryEscape(input)
	return encoded, nil
}

// URLDecode decodes a URL-encoded string
func URLDecode(input string) (string, error) {

	// 对 url 字符进行转义解码，输出字符串
	//decoded := dongle.Decode.FromString(input).BySafeURL().ToString() // www.gouguoyin.cn?sex=男&age=18
	//return decoded, nil
	return url.QueryUnescape(input)
}
