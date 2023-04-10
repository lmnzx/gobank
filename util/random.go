package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandString generates a integer between min and max
func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandString generates a random string of length n
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// RandOwner generates a random owner name
func RandOwner() string {
	return RandString(6)
}

// RandMoney generates a random amount of money
func RandMoney() int64 {
	return RandInt(0, 1000)
}

// RandCurrency generates a random currency
func RandCurrency() string {
	currencies := []string{"USD", "EUR", "GBP"}
	return currencies[rand.Intn(len(currencies))]
}
