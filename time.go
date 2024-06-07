package swiss

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ErrTimeFormat = errors.New("unrecognized time format")

var (
	datePatterns = map[string]string{
		"yesterday":    `yesterday`,
		"today":        `today`,
		"tomorrow":     `tomorrow`,
		"minutes_ago":  `(\d+)\s*min(?:ute)?s?\s*ago`,
		"minutes_from": `in\s*(\d+)\s*min(?:ute)?s?`,
		"hours_ago":    `(\d+)\s*hour(?:s)?\s*ago`,
		"hours_from":   `in\s*(\d+)\s*hour(?:s)?`,
		"days_ago":     `(\d+)\s*day(?:s)?\s*ago`,
		"days_from":    `in\s*(\d+)\s*day(?:s)?`,
		"weeks_ago":    `(\d+)\s*week(?:s)?\s*ago`,
		"weeks_from":   `in\s*(\d+)\s*week(?:s)?`,
		"months_ago":   `(\d+)\s*month(?:s)?\s*ago`,
		"months_from":  `in\s*(\d+)\s*month(?:s)?`,
		"years_ago":    `(\d+)\s*year(?:s)?\s*ago`,
		"years_from":   `in\s*(\d+)\s*year(?:s)?`,
	}
	timeReg = regexp.MustCompile(`at\s*(\d{1,2})(:(\d{2}))?\s*(am|pm)?`)
)

// TimeFromHumanReadable will take in a string and try to parse out a time.
func TimeFromHumanReadable(s string) (time.Time, error) {
	s = strings.ToLower(s) // to lower so we don't need complicated regex
	now := time.Now().UTC()
	builtTime := time.Time{}

	for unit, pattern := range datePatterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(s)

		if len(matches) > 0 {
			switch unit {
			case "yesterday":
				builtTime = now.AddDate(0, 0, -1)
			case "tomorrow":
				builtTime = now.AddDate(0, 0, 1)
			default:
				amount, err := strconv.Atoi(matches[1])
				if err != nil {
					return time.Time{}, err
				}

				switch unit {
				case "minutes_ago":
					builtTime = now.Add(time.Duration(-amount) * time.Minute)
				case "minutes_from":
					builtTime = now.Add(time.Duration(amount) * time.Minute)
				case "hours_ago":
					builtTime = now.Add(time.Duration(-amount) * time.Hour)
				case "hours_from":
					builtTime = now.Add(time.Duration(amount) * time.Hour)
				case "days_ago":
					builtTime = now.AddDate(0, 0, -amount)
				case "days_from":
					builtTime = now.AddDate(0, 0, amount)
				case "weeks_ago":
					builtTime = now.AddDate(0, 0, -amount*7)
				case "weeks_from":
					builtTime = now.AddDate(0, 0, amount*7)
				case "months_ago":
					builtTime = now.AddDate(0, -amount, 0)
				case "months_from":
					builtTime = now.AddDate(0, amount, 0)
				case "years_ago":
					builtTime = now.AddDate(-amount, 0, 0)
				case "years_from":
					builtTime = now.AddDate(amount, 0, 0)
				}
			}
		}
	}

	if matches := timeReg.FindStringSubmatch(s); len(matches) > 0 {
		hour, err := strconv.Atoi(matches[1])
		if err != nil {
			return time.Time{}, err
		}

		minute := 0
		if matches[3] != "" {
			minute, err = strconv.Atoi(matches[3])
			if err != nil {
				return time.Time{}, err
			}
		}

		meridian := matches[4]
		if meridian == "pm" && hour != 12 {
			hour += 12
		}
		if meridian == "am" && hour == 12 {
			hour = 0
		}

		builtTime = time.Date(builtTime.Year(), builtTime.Month(), builtTime.Day(), hour, minute, 0, 0, builtTime.Location())
	}

	if builtTime.IsZero() {
		return time.Time{}, ErrTimeFormat
	}

	return builtTime, nil
}
