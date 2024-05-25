package swiss

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// RandomSeed will generate a seed based on the current UnixNano.
func RandomSeed() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString creates an alphanumeric string of a given length.
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
