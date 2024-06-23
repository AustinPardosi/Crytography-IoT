package Backend___Mobile__ECC___Digital_Signature_

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"ecc"
	"fmt"
)

func main() {
	// Contoh data glukosa yang akan dikirim
	data := []byte("95.4")

	// Kunci publik dan privat untuk enkripsi dan tanda tangan
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey

	// Enkripsi data
	encryptedData, err := ecc.EncryptDataECC(data, publicKey)
	if err != nil {
		fmt.Println("Enkripsi gagal:", err)
	} else {
		fmt.Println("Data terenkripsi:", encryptedData)
	}

	// Tanda tangan data
	signature, err := ecc.SignDataECC(data, privateKey)
	if err != nil {
		fmt.Println("Tanda tangan gagal:", err)
	} else {
		fmt.Println("Tanda tangan:", signature)
	}

	// Verifikasi tanda tangan
	err = ecc.VerifySignatureECC(data, signature, publicKey)
	if err != nil {
		fmt.Println("Verifikasi tanda tangan gagal:", err)
	} else {
		fmt.Println("Verifikasi tanda tangan berhasil")
	}
}
