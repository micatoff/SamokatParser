package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func RandomNumber(n int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	digits := []rune("0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = digits[r.Intn(len(digits))]
	}
	return string(s)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
