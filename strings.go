package swiss

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/bidi"
)

var Language = language.English

var validEmailRE = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsEmail checks if a string is a valid standard email address.
// RFC 5322 is too permissive for general use such as quoted strings and local hosts.
func IsEmail(s string) bool {
	return validEmailRE.MatchString(s)
}

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

// IsSnakeCase checks to see if supplied string is in snake case.
func IsSnakeCase(s string) bool {
	return strings.Contains(s, "_") && IsLower(s) && !strings.Contains(s, " ")
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

// Reverse returns the string in reverse order.
// This function is an alias for [bidi.ReverseString].
func Reverse(s string) string {
	return bidi.ReverseString(s)
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
	return cases.Title(Language).String(s)
}

// CamelCase converts a string to camel case.
func CamelCase(s string) string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	if IsSnakeCase(s) {
		words = strings.Split(s, "_")
	}
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

// SwapCase swaps the case of a string.
func SwapCase(str string) string {
	var r string
	for _, c := range str {
		switch {
		case c >= 'a' && c <= 'z':
			r += string(c - 32)
		case c >= 'A' && c <= 'Z':
			r += string(c + 32)
		default:
			r += string(c)
		}
	}
	return r
}

// Slugify will take a string and create a dash separated string for use in URLs.
func Slugify(s string) string {
	s = strings.ToLower(s)
	reg := regexp.MustCompile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")

	return s
}
