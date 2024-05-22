package swiss

import (
	"strings"
)

// IsUpper checks if a string is all uppercase.
func IsUpper(s string) bool {
	return len(s) > 0 && s == strings.ToUpper(s)
}

// IsLower checks if a string is all lowercase.
func IsLower(s string) bool {
	return len(s) > 0 && s == strings.ToLower(s)
}

// IsAlpha checks if a string is all alphabetic characters.
func IsAlpha(s string) bool {
	return len(s) > 0 && s == strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return r
		}
		return -1
	}, s)
}

// IsNumeric checks if a string is all numeric characters.
func IsNumeric(s string) bool {
	return len(s) > 0 && s == strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// IsAlphaNumeric checks if a string is all alphanumeric characters.
func IsAlphaNumeric(s string) bool {
	return len(s) > 0 && s == strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// IsHexChar checks if a character is a valid hexadecimal character.
func IsHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// IsUUID will take a string and determine if it is a valid UUID,
// it will return true if it is and false if it is not.
func IsUUID(s string) bool {
	if len(s) != 36 {
		return false
	}

	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			if s[i] != '-' {
				return false
			}
		default:
			if !IsHexChar(byte(s[i])) {
				return false
			}
		}
	}
	return true
}

// SnakeCase converts a string to snake case.
// Inputs can be space separated, camel case, or pascal case.
func SnakeCase(s string) string {
	if len(s) < 1 {
		return s
	}

	hasSpaces := strings.Contains(s, " ")

	ss := strings.Split(s, "")
	res := []string{}
	for i := 0; i < len(ss); i++ {
		if hasSpaces {
			if ss[i] == " " {
				res = append(res, "_")
				continue
			}
		} else {
			if i != 0 && IsUpper(ss[i]) {
				res = append(res, "_", strings.ToLower(ss[i]))
				continue
			}
		}
		res = append(res, strings.ToLower(ss[i]))
	}
	return strings.Join(res, "")
}

// TitleCase converts a string to title case.
func TitleCase(s string) string {
	if len(s) < 1 {
		return s
	}
	ss := strings.Split(s, "")
	ss[0] = strings.ToUpper(ss[0])
	for i := 1; i < len(ss); i++ {
		if ss[i-1] == " " {
			ss[i] = strings.ToUpper(ss[i])
		} else {
			ss[i] = strings.ToLower(ss[i])
		}
	}
	return strings.Join(ss, "")
}

// CamelCase converts a string to camel case.
func CamelCase(s string) string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	for i, word := range words {
		if i == 0 {
			continue
		}

		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, "")
}

// PascalCase converts a string to pascal case, it can take in a string that is both
// title case and camel case and convert it to camel case.
func PascalCase(s string) string {
	c := CamelCase(s)
	return strings.ToUpper(c[:1]) + c[1:]
}
