package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func generateRSAKeys() (*rsa.PrivateKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, err
    }
    return privateKey, nil
}

func rsaEncrypt(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
    label := []byte("")
    hash := sha256.New()
    ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, plaintext, label)
    if err != nil {
        return nil, err
    }
    return ciphertext, nil
}

func rsaDecrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
    label := []byte("")
    hash := sha256.New()
    plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
    if err != nil {
        return nil, err
    }
    return plaintext, nil
}

func main() {
    privateKey, _ := generateRSAKeys()
    publicKey := &privateKey.PublicKey

    message := "Hello, RSA!"
    fmt.Println("Original message:", message)

    encrypted, _ := rsaEncrypt(publicKey, []byte(message))
    fmt.Println("Encrypted message:", encrypted)

    decrypted, _ := rsaDecrypt(privateKey, encrypted)
    fmt.Println("Decrypted message:", string(decrypted))
}
