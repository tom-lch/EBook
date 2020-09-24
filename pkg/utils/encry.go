package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func SHA1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func SHA256(str string) string {
	s := sha256.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
