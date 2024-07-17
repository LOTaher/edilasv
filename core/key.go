package core

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateKey() string {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(key)
}
