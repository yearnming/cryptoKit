package encoding

import (
	"fmt"
)

// EscapeEncode encodes a string into %uXXXX format
func EscapeEncode(input string) (string, error) {
	output := ""
	for _, r := range input {
		output += fmt.Sprintf("%%u%04X", r)
	}
	return output, nil
}

// EscapeDecode decodes a %uXXXX encoded string into a normal string
func EscapeDecode(input string) (string, error) {
	var output string
	for len(input) > 0 {
		if input[0:2] == "%u" {
			var r rune
			_, err := fmt.Sscanf(input[2:6], "%04X", &r)
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
