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
