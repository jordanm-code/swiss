package swiss

import (
	"testing"
	"time"
)

func TestTimeFromHumanReadable(t *testing.T) {
	now := time.Now().UTC()
	tests := []struct {
		input string
		want  time.Time
	}{
		{"yesterday", now.AddDate(0, 0, -1)},
		{"5 mins ago", now.Add(-5 * time.Minute)},
		{"5 minutes ago", now.Add(-5 * time.Minute)},
		{"1 week ago", now.AddDate(0, 0, -7)},
		{"tomorrow", now.AddDate(0, 0, 1)},
		{"in 5 minutes", now.Add(5 * time.Minute)},
		{"in 3 weeks", now.AddDate(0, 0, 21)},
		{"tomorrow at 4:00 AM", time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, now.Location()).AddDate(0, 0, 1)},
		{"in 1 year", now.AddDate(1, 0, 0)},
		{"in 12 months", now.AddDate(1, 0, 0)},
		{"6 months ago", now.AddDate(0, -6, 0)},
		{"yesterday at 4AM", time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, now.Location()).AddDate(0, 0, -1)},
	}
	for _, test := range tests {
		got, err := TimeFromHumanReadable(test.input)
		if err != nil {
			t.Error(err)
			continue
		}

		if got.IsZero() {
			t.Errorf("TimeFromHumanReadable(%v) returned zero time", test.input)
			continue
		}

		if !withinDurationJitter(got, test.want, time.Second) {
			t.Errorf("TimeFromHumanReadable(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

// withinDurationJitter checks if two times are withing a given duration of
// eachother, so the nanoseconds don't cause flaky tests.
func withinDurationJitter(t1, t2 time.Time, d time.Duration) bool {
	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}
	return diff <= d
}
