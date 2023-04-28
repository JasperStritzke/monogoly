package util

import (
	crypto "crypto/rand"
	"math/big"
	math "math/rand"
)

var runes = []rune("0123456789")

func RandomGameId() string {
	var result []rune

	result = append(result, runes[math.Intn(len(runes)-1)+1])
	for i := 0; i < 3; i++ {
		result = append(result, runes[math.Intn(len(runes))])
	}

	return string(result)
}

func RandomSessionId() string {
	const n = 32
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	result := make([]byte, n)

	for i := 0; i < n; i++ {
		num, err := crypto.Int(crypto.Reader, big.NewInt(int64(len(letters))))

		if err != nil {
			return ""
		}

		result[i] = letters[num.Int64()]
	}

	return string(result)
}
