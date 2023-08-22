package utils

import (
	"math/rand"
	"strings"
)

var alphabets = "abcdefghijklmnopqrstuvwxyz"

// generate random integers
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate random strings
func randomString(n int) string {
	var builder strings.Builder

	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		builder.WriteByte(c)
	}
	return builder.String()
}

// generate a random Owner
func RandomOwner() string {
	return randomString(5)
}

func RandomBalance() int64 {
	return randomInt(100 , 1000)
}

func RandomID() int64 { 
	return randomInt(1, 200)
}

func RandomCurrency() string {
	currencies := []string{"NGN", "GBP", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}