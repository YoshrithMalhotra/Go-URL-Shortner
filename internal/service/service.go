package service

import (
	"crypto/rand"
	"math/big"
)

const shortCodeLength = 6

var shortCodeAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateShortCode() string {
	code := make([]rune, shortCodeLength)
	for index := range code {
		number, err := rand.Int(rand.Reader, big.NewInt(int64(len(shortCodeAlphabet))))
		if err != nil {
			code[index] = shortCodeAlphabet[0]
			continue
		}
		code[index] = shortCodeAlphabet[number.Int64()]
	}
	return string(code)
}
