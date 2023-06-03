package algo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func GetChecksum(data []byte, hashname string) string {
	if hashname == "md5" {
		return fmt.Sprintf("%x", md5.Sum(data))
	} else if hashname == "sha1" {
		return fmt.Sprintf("%x", sha1.Sum(data))
	} else {
		return fmt.Sprintf("%x", sha256.Sum256(data))
	}
}
