package swiss

import "strings"

// IsUpper checks if a string is all uppercase.
func IsUpper(s string) bool {
	return len(s) > 0 && s == strings.ToUpper(s)
}

// IsLower checks if a string is all lowercase.
func IsLower(s string) bool {
	return len(s) > 0 && s == strings.ToLower(s)
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

// CamelCase converts a string to camel case
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
