package utils

import (
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

// RandomInt returns a random integer between min and max.
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomString returns a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner returns a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random amount of money between 0 and 1000.
func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

// RandomCurrency returns a random currency from a list of currencies.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
