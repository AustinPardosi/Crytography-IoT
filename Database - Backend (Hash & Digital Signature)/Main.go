package Database___Backend__Hash___Digital_Signature_

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// Fungsi utama
func main() {
	// Contoh data yang diterima dari database
	data := []byte("95.4")
	signature := "base64_encoded_signature"

	// Kunci publik untuk verifikasi
	publicKeyPEM := `-----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...
    -----END PUBLIC KEY-----`
	block, _ := pem.Decode([]byte(publicKeyPEM))
	publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPublicKey := publicKey.(*rsa.PublicKey)

	// Verifikasi tanda tangan
	err := verifySignature(data, signature, rsaPublicKey)
	if err != nil {
		fmt.Println("Verifikasi gagal:", err)
	} else {
		fmt.Println("Verifikasi berhasil")
	}
}
