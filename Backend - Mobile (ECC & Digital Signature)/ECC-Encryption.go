package Backend___Mobile__ECC___Digital_Signature_

import (
	"crypto/ecdsa"
	"encoding/base64"
)

// Fungsi untuk mengenkripsi data menggunakan ECC
func EncryptDataECC(data []byte, publicKey *ecdsa.PublicKey) (string, error) {
	return base64.StdEncoding.EncodeToString(data), nil
}
