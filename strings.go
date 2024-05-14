package swiss

import "strings"

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
