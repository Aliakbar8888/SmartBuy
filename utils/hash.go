package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(i string) string {
	hash := sha256.New()
	hash.Write([]byte(i))
	hashByte := hash.Sum(nil)
	return hex.EncodeToString(hashByte)
}
