package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomID() uuid.UUID { return uuid.New() }

func RandomInt(min, max int64) int64 { return min + rand.Int63n(max-min+1) }

//
func RandomAmount() int64 { return RandomInt(0, 1000) }

const alphabets = "abcdefghijklmnopqrstuvwxyz"
const digits = "0123456789"

func RandomString(n int) string {
	var sb strings.Builder
	l := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPhoneNumber(n int) string {
	var sb strings.Builder
	l := len(digits)

	for i := 0; i < n; i++ {
		c := digits[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string { return RandomString(15) }

func RandomEmail() string { return fmt.Sprintf("%s@email.com", RandomString(6)) }
