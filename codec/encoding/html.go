package encoding

import (
	"fmt"
	"html"
	"strings"
	"unicode"
)

// HTML16Encode HTMLEncode encodes a string into HTML-encoded format
func HTML16Encode(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			encoded.WriteRune(r)
		} else {
			encoded.WriteString(fmt.Sprintf("&#x%X;", r))
		}
	}
	return encoded.String(), nil
}

func HTML16EncodeAll(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		encoded.WriteString(fmt.Sprintf("&#x%X;", r))
	}
	return encoded.String(), nil
}

func HTML10Encode(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			encoded.WriteRune(r)
		} else {
			encoded.WriteString(fmt.Sprintf("&#%d;", r))
		}
	}
	return encoded.String(), nil
}

func HTML10EncodeAll(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		encoded.WriteString(fmt.Sprintf("&#%d;", r))
	}
	return encoded.String(), nil
}

func customEncode(r rune) string {
	switch r {
	case '!':
		return "&excl;"
	case '"':
		return "&quot;"
	case '#':
		return "&num;"
	case '$':
		return "&dollar;"
	case '%':
		return "&percnt;"
	case '&':
		return "&amp;"
	case '\'':
		return "&apos;"
	case '(':
		return "&lpar;"
	case ')':
		return "&rpar;"
	case '*':
		return "&ast;"
	case '+':
		return "&plus;"
	case ',':
		return "&comma;"
	//case '-':
	//	return "&minus;"
	case '.':
		return "&period;"
	case '/':
		return "&sol;"
	case ':':
		return "&colon;"
	case ';':
		return "&semi;"
	case '<':
		return "&lt;"
	case '=':
		return "&equals;"
	case '>':
		return "&gt;"
	case '?':
		return "&quest;"
	case '@':
		return "&commat;"
	case '[':
		return "&lsqb;"
	case '\\':
		return "&bsol;"
	case ']':
		return "&rsqb;"
	case '^':
		return "&Hat;"
	case '_':
		return "&lowbar;"
	case '`':
		return "&grave;"
	case '{':
		return "&lcub;"
	case '|':
		return "&verbar;"
	case '}':
		return "&rcub;"
	default:
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return string(r)
		}
		return fmt.Sprintf("&#x%X;", r)
	}
}

func HTMLruneEncode(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		if entity, exists := htmlEntityMap[r]; exists {
			encoded.WriteString(entity)
		} else {
			encoded.WriteRune(r)
		}
	}
	return encoded.String(), nil
}

func HTMLrune10Encode(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		if entity, exists := htmlEntityMap[r]; exists {
			encoded.WriteString(entity)
		} else {
			encoded.WriteString(fmt.Sprintf("&#%d;", r))
		}

	}
	return encoded.String(), nil
}

func HTMLrune16Encode(input string) (string, error) {
	var encoded strings.Builder
	for _, r := range input {
		if entity, exists := htmlEntityMap[r]; exists {
			encoded.WriteString(entity)
		} else {
			encoded.WriteString(fmt.Sprintf("&#x%X;", r))
		}

	}
	return encoded.String(), nil
}

// HTMLDecode decodes an HTML-encoded string
func HTMLDecode(input string) (string, error) {
	return html.UnescapeString(input), nil
}

// htmlEntityMap 存储字符到HTML实体编码的映射
var htmlEntityMap = map[rune]string{
	'&':  "&amp;",
	'<':  "&lt;",
	'>':  "&gt;",
	'"':  "&quot;",
	'\'': "&apos;",
	'©':  "&copy;",
	'!':  "&excl;",
	'#':  "&num;",
	'$':  "&dollar;",
	'%':  "&percnt;",
	'(':  "&lpar;",
	')':  "&rpar;",
	'*':  "&ast;",
	'+':  "&plus;",
	',':  "&comma;",
	'.':  "&period;",
	'/':  "&sol;",
	':':  "&colon;",
	';':  "&semi;",
	'=':  "&equals;",
	'?':  "&quest;",
	'@':  "&commat;",
	'[':  "&lsqb;",
	'\\': "&bsol;",
	']':  "&rsqb;",
	'^':  "&Hat;",
	'_':  "&lowbar;",
	'`':  "&grave;",
	'{':  "&lcub;",
	'|':  "&verbar;",
	'}':  "&rcub;",
}
