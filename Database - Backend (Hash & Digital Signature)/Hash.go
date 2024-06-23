package Database___Backend__Hash___Digital_Signature_

import (
	"crypto/sha256"
)

// Fungsi untuk membuat hash dari data
func createHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
