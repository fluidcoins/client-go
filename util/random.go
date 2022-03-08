package util

import (
	"encoding/base64"
	"math/rand"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func RandomBool() bool {
	b := []bool{true, false}
	n := len(b)
	return b[rand.Intn(n)]
}

func RandomCurrency() string {
	b := []string{"USD", "NGN"}
	n := len(b)
	return b[rand.Intn(n)]
}

func RandomInt(min, max int64) int64 { return min + rand.Int63n(max-min+1) }

func RandomAmount() int64 { return RandomInt(7777, 598970) }
