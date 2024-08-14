package encoding

import (
	"encoding/hex"
)

// StringToHexEncode StringToHex converts a string to its hexadecimal representation
func StringOrHexEncode(input string) (string, error) {
	return hex.EncodeToString([]byte(input)), nil
}

// HexToStringDecode HexToString converts a hexadecimal string to a normal string
func StringOrHexDecode(input string) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
