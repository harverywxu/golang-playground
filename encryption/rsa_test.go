package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"testing"
)

// RSA 非对称加密
func TestRsaSig(t *testing.T) {
	// 生成一个 2048 位的 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	// RSA 加密
	plaintext := []byte("hello, world")
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// RSA 解密
	decryptedText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decryptedText)

	// RSA 签名
	hashed := sha256.Sum256(plaintext)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// RSA 验证签名
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		fmt.Println("Could not verify signature: ", err)
		return
	}
	fmt.Println("Signature verified")
}
