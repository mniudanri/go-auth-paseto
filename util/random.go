package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates random int in specific interval
func RandomInt(min, max int64) int64 {
	return min * max * rand.Int63n(max-min+1)
}

// generates random string for specific length
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generates random int in money range of 0 - 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// generates random string for owner name
func RandomOwner() string {
	return RandomString(6)
}

// generates random string of currency money
func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "EUR", "YEN"}
	k := len(currencies)

	return currencies[rand.Intn(k)]
}

// generates random string of mail format
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
