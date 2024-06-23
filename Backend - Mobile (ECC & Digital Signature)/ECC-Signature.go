package Backend___Mobile__ECC___Digital_Signature_

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

// Fungsi untuk menandatangani data menggunakan kunci privat
func SignDataECC(data []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Fungsi untuk memverifikasi tanda tangan menggunakan kunci publik
func VerifySignatureECC(data []byte, signature string, publicKey *ecdsa.PublicKey) error {
	hash := sha256.Sum256(data)
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	r := new(big.Int).SetBytes(decodedSignature[:len(decodedSignature)/2])
	s := new(big.Int).SetBytes(decodedSignature[len(decodedSignature)/2:])
	valid := ecdsa.Verify(publicKey, hash[:], r, s)
	if !valid {
		return fmt.Errorf("signature verification failed")
	}
	return nil
}
