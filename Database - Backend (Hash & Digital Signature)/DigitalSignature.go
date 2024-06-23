package Database___Backend__Hash___Digital_Signature_

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

// Fungsi untuk menandatangani data menggunakan kunci privat
func signData(data []byte, privateKey *rsa.PrivateKey) (string, error) {
	hash := createHash(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Fungsi untuk memverifikasi tanda tangan menggunakan kunci publik
func verifySignature(data []byte, signature string, publicKey *rsa.PublicKey) error {
	hash := createHash(data)
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, decodedSignature)
}
