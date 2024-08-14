package encoding

import (
	"fmt"
)

// ShellcodeEncode encodes a string into shellcode format (\xXX)
func ShellcodeEncode(input string) (string, error) {
	output := ""
	for _, r := range input {
		output += fmt.Sprintf("\\x%02x", r)
	}
	return output, nil
}

// ShellcodeDecode decodes a shellcode string (\xXX) into a normal string
func ShellcodeDecode(input string) (string, error) {
	var output string
	for len(input) > 0 {
		if input[0:2] == "\\x" {
			var b byte
			_, err := fmt.Sscanf(input[2:4], "%02x", &b)
			if err != nil {
				return "", err
			}
			output += string(b)
			input = input[4:]
		} else {
			output += string(input[0])
			input = input[1:]
		}
	}
	return output, nil
}
