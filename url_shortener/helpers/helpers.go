package helpers

import (
	"crypto/rand"
	"math/big"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const shortURLLength = 6

func GenerateShortURL(longURL string) string {
	var shortURL string
	for i := 0; i < shortURLLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		shortURL += string(alphabet[num.Int64()])
	}
	return shortURL
}
