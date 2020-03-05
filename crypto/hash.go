package crypto

import "crypto/sha256"

func SHA256(data []byte) string {
	checksum := sha256.New().Sum(data)

	return string(checksum)
}

func IsSame(old, new []byte) bool {
	return SHA256(old) == SHA256(new)
}
