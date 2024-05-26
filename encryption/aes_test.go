package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"testing"
)
// AES 对称加密
func TestAes(t *testing.T) {
	// AES密钥，长度可以是16（AES-128）、24（AES-192）或32（AES-256）
	key := []byte("myverystrongpasswordo32bitlength")

	// 要加密的数据
	plaintext := []byte("Hello, world!")

	// 创建cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 对数据进行填充，使其长度满足AES的要求
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)

	// 创建加密器
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// 创建解密器
	mode = cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(ciphertext, ciphertext)

	// 去掉填充
	padding = int(ciphertext[len(ciphertext)-1])
	plaintext = ciphertext[:len(ciphertext)-padding]

	fmt.Printf("Plaintext: %s\n", plaintext)
}
