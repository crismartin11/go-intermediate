package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateState() string {
	// Generate a random byte array for state paramter
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
