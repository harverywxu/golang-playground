package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaActiveCode(t *testing.T) {
	// 生成一个RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 客户的环境信息
	info := "MachineID, InstallationDate"

	// 使用私钥生成激活码
	activationCode, err := GenerateActivationCode(privateKey, info)
	if err != nil {
		panic(err)
	}

	fmt.Println("Activation Code:", activationCode)

	// 使用公钥验证激活码
	isValid, err := ValidateActivationCode(publicKey, info, activationCode)
	if err != nil {
		panic(err)
	}

	fmt.Println("Is Valid:", isValid)
}

// GenerateActivationCode 生成激活码
func GenerateActivationCode(privateKey *rsa.PrivateKey, info string) (string, error) {
	hashedInfo := sha256.Sum256([]byte(info))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedInfo[:])
	if err != nil {
		return "", err
	}

	// 为了使激活码不超过12位，我们只取签名的前12位
	// 注意：这将降低激活码的安全性
	return base64.StdEncoding.EncodeToString(signature[:12]), nil
}

// ValidateActivationCode 验证激活码
func ValidateActivationCode(publicKey *rsa.PublicKey, info, activationCode string) (bool, error) {
	signature, err := base64.StdEncoding.DecodeString(activationCode)
	if err != nil {
		return false, err
	}

	// 由于我们在生成激活码时只取了签名的前12位，所以在验证时也只能验证这12位
	// 注意：这将降低激活码的安全性
	hashedInfo := sha256.Sum256([]byte(info))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedInfo[:], signature)
	if err != nil {
		return false, nil
	}

	return true, nil
}