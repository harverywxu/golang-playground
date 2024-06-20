package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestHmacSHA256Test(t *testing.T) {
	// 密钥
	key := []byte("secret")

	// 客户的环境信息
	info := "MachineID, InstallationDate"
	info = ""

	// 生成激活码
	activationCode := GenerateActivationCodeV2(key, info)

	fmt.Println("Activation Code:", activationCode)

	// 验证激活码
	isValid := ValidateActivationCodeV2(key, info, activationCode)

	fmt.Println("Is Valid:", isValid)
}

// GenerateActivationCode 生成激活码
func GenerateActivationCodeV2(key []byte, info string) string {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(info))

	// 为了使激活码不超过12位，我们只取哈希值的前12位
	// 注意：这将降低激活码的安全性
	hs := h.Sum(nil)
	str := base64.StdEncoding.EncodeToString(hs)
	return str[:12]
}

// ValidateActivationCode 验证激活码
func ValidateActivationCodeV2(key []byte, info, activationCode string) bool {
	expectedActivationCode := GenerateActivationCodeV2(key, info)
	return hmac.Equal([]byte(expectedActivationCode), []byte(activationCode))
}

