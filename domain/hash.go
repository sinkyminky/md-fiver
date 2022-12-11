package domain

import (
	"crypto/md5"
	"fmt"
)

func Hash(m []byte) string {
	hash := md5.Sum(m)
	return fmt.Sprintf("%x", hash)
}
