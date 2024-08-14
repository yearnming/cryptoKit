package encoding

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// ConvertOrBinaryEncode 将字符串转换为其二进制表示形式
func ConvertOrBinaryEncode(s string) (string, error) {
	var binaryStr strings.Builder
	for _, r := range s {
		b := []byte(string(r))
		for _, byteVal := range b {
			binaryStr.WriteString(fmt.Sprintf("%08b ", byteVal))
		}
	}
	return strings.TrimSpace(binaryStr.String()), nil
}

func ConvertOrBinaryByteEncode(s byte) ([]byte, error) {
	var binaryStr bytes.Buffer

	binaryStr.WriteString(fmt.Sprintf("%08b ", s))

	return binaryStr.Bytes(), nil
}

// ConvertFromBinaryDecode ConvertFromBinary 将二进制表示转换为字符串
func ConvertOrBinaryDecode(binaryStr string) (string, error) {
	binaryCodes := strings.Split(binaryStr, " ")
	var buf bytes.Buffer
	for _, binaryCode := range binaryCodes {
		var byteVal byte
		_, err := fmt.Sscanf(binaryCode, "%b", &byteVal)
		if err != nil {
			return "", err
		}
		buf.WriteByte(byteVal)
	}
	return buf.String(), nil
}

// ConvertToOctalEncode ConvertToOctal 将字符串转换为其八进制表示形式
func ConvertOrOctalEncode(s string) (string, error) {
	var octalStr strings.Builder
	for _, r := range s {
		b := []byte(string(r))
		for _, byteVal := range b {
			octalStr.WriteString(fmt.Sprintf("%03o ", byteVal))
		}
	}
	return strings.TrimSpace(octalStr.String()), nil
}

func ConvertOrOctalByteEncode(s byte) ([]byte, error) {
	var octalStr bytes.Buffer

	octalStr.WriteString(fmt.Sprintf("%03o ", s))

	return octalStr.Bytes(), nil
}

// ConvertFromOctalDecode ConvertFromOctal 将八进制表示转换为字符串
func ConvertOrOctalDecode(octalStr string) (string, error) {
	octalCodes := strings.Split(octalStr, " ")
	var buf bytes.Buffer
	for _, octalCode := range octalCodes {
		var byteVal byte
		_, err := fmt.Sscanf(octalCode, "%03o", &byteVal)
		if err != nil {
			return "", err
		}
		buf.WriteByte(byteVal)
	}
	return buf.String(), nil
}

// ConvertToDecimalEncode ConvertToDecimal 将字符串转换为其十进制表示形式
func ConvertOrDecimalEncode(s string) (string, error) {
	var decimalStr strings.Builder
	for _, r := range s {
		b := []byte(string(r))
		for _, byteVal := range b {
			decimalStr.WriteString(fmt.Sprintf("%03d ", byteVal))
		}
	}
	return strings.TrimSpace(decimalStr.String()), nil
}

func ConvertOrDecimalByteEncode(s byte) ([]byte, error) {
	var decimalStr bytes.Buffer

	decimalStr.WriteString(fmt.Sprintf("%03d", s))

	return decimalStr.Bytes(), nil
}

// ConvertFromDecimalDecode ConvertFromDecimal 将十进制表示形式转换为字符串
func ConvertOrDecimalDecode(decimalStr string) (string, error) {
	decimalCodes := strings.Split(decimalStr, " ")
	var buf bytes.Buffer
	for _, decimalCode := range decimalCodes {
		var byteVal byte
		_, err := fmt.Sscanf(decimalCode, "%d", &byteVal)
		if err != nil {
			return "", err
		}
		buf.WriteByte(byteVal)
	}
	return buf.String(), nil
}

// ConvertOrHexEncode  将字符串转换为其十六进制表示形式
func ConvertOrHexEncode(s string) (string, error) {
	var hexStr strings.Builder
	for _, r := range s {
		b := []byte(string(r))
		for _, byteVal := range b {
			hexStr.WriteString(fmt.Sprintf("%02x ", byteVal))
		}
	}
	return strings.TrimSpace(hexStr.String()), nil
}

func ConvertOrHexDecode(hexStr string) (string, error) {
	if len(hexStr)%2 != 0 {
		return "", fmt.Errorf("无效的十六进制字符串长度")
	}

	var buf bytes.Buffer
	for i := 0; i < len(hexStr); i += 2 {
		var byteVal byte
		_, err := fmt.Sscanf(hexStr[i:i+2], "%x", &byteVal)
		if err != nil {
			return "", err
		}
		buf.WriteByte(byteVal)
	}
	return buf.String(), nil
}

// ConvertOrHexByteDecode 将十六进制字符串解码为字节切片
func ConvertOrHexByteDecode(hexStr []byte) ([]byte, error) {
	if len(hexStr)%2 != 0 {
		return nil, fmt.Errorf("无效的十六进制字符串长度")
	}

	var buf bytes.Buffer
	for i := 0; i < len(hexStr); i += 2 {
		// 使用strconv.ParseUint替代fmt.Sscanf来处理十六进制到字节的转换
		byteVal, err := strconv.ParseUint(string(hexStr[i:i+2]), 16, 8)
		if err != nil {
			return nil, err
		}
		buf.WriteByte(byte(byteVal))
	}
	return buf.Bytes(), nil
}

func ConvertOrHexByteEncode(data byte) ([]byte, error) {
	// 使用bytes.Buffer来构建十六进制字符串
	var hexBuf bytes.Buffer
	// 将接收到的单个字节转换为十六进制表示，并写入缓冲区
	hexBuf.WriteString(fmt.Sprintf("%02x", data))
	return hexBuf.Bytes(), nil
}
