package encoding

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// ASCIIEncodeToBinaryEncode ASCIIEncodeToBinary converts a string to its ASCII binary representation
func ASCIIOrBinaryEncode(s string) (string, error) {
	var binaryStr strings.Builder
	for _, r := range s {
		b := []byte(string(r))
		for _, byteVal := range b {
			binaryStr.WriteString(fmt.Sprintf("%08b ", byteVal))
		}
	}
	return strings.TrimSpace(binaryStr.String()), nil
}

// ASCIIOrBinaryDecode ASCIIDecodeFromBinary converts an ASCII binary representation to a string
func ASCIIOrBinaryDecode(binaryStr string) (string, error) {
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

// ASCIIEncodeToOctalEncode ASCIIEncodeToOctal 将字符串转换为其 ASCII 八进制表示形式
func ASCIIOrOctalEncode(s string) (string, error) {
	var octalStr strings.Builder
	for _, r := range s {
		b := []byte(string(r))
		for _, byteVal := range b {
			octalStr.WriteString(fmt.Sprintf("%03o ", byteVal))
		}
	}
	return strings.TrimSpace(octalStr.String()), nil
}

// ASCIIDecodeFromOctalDecode ASCIIDecodeFromOctal converts an ASCII octal representation to a string
func ASCIIOrOctalDecode(octalStr string) (string, error) {
	octalCodes := strings.Split(octalStr, " ")
	var buf bytes.Buffer
	for _, octalCode := range octalCodes {
		var byteVal byte
		_, err := fmt.Sscanf(octalCode, "%o", &byteVal)
		if err != nil {
			return "", err
		}
		buf.WriteByte(byteVal)
	}
	return buf.String(), nil
}

// ASCIIEncodeToDecimalEncode ASCIIEncodeToDecimal converts a string to its ASCII decimal representation
func ASCIIOrDecimalEncode(input string) (string, error) {
	var output strings.Builder
	for _, r := range input {
		output.WriteString(fmt.Sprintf("%d ", r))
	}
	return strings.TrimSpace(output.String()), nil
}

// ASCIIDecodeFromDecimal converts an ASCII decimal representation to a string
func ASCIIOrDecimalDecode(input string) (string, error) {
	var output strings.Builder
	decimals := strings.Fields(input)
	for _, dec := range decimals {
		val, err := strconv.ParseUint(dec, 10, 8)
		if err != nil {
			return "", err
		}
		output.WriteByte(byte(val))
	}
	return output.String(), nil
}

// ASCIIEncodeToHex converts a string to its ASCII hexadecimal representation
func ASCIIOrHexEncode(input string) (string, error) {
	var output strings.Builder
	for _, r := range input {
		output.WriteString(fmt.Sprintf("%02x ", r))
	}
	return strings.TrimSpace(output.String()), nil
}

// ASCIIDecodeFromHex converts an ASCII hexadecimal representation to a string
func ASCIIOrHexDecode(input string) (string, error) {
	var output strings.Builder
	hexes := strings.Fields(input)
	for _, hex := range hexes {
		val, err := strconv.ParseUint(hex, 16, 8)
		if err != nil {
			return "", err
		}
		output.WriteByte(byte(val))
	}
	return output.String(), nil
}
