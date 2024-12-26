package mycrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

// Encrypt 使用AES加密，返回iv+cipher
func Encrypt(message, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Pad the message to be a multiple of the block size
	blockSize := block.BlockSize()
	paddedMessage := padCS7(message, blockSize)

	// Generate a random initialization vector (IV)
	iv := make([]byte, blockSize)

	// 给初始向量赋值默认值，在返回密文时无需再返回向量，使密文空间始终为256比特（交易单次只能嵌入256）
	//iv := []byte("1234567890123456")

	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	// Encrypt the message using CBC mode
	ciphertext := make([]byte, len(paddedMessage))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedMessage)

	// Return the IV concatenated with the ciphertext
	return append(iv, ciphertext...), nil
	//return ciphertext, nil
}

// Decrypt decrypts the ciphertext to retrieve the original plaintext message.
func Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Split the IV and the actual ciphertext
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]
	//iv := []byte("1234567890123456")

	// Decrypt using CBC mode
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove padding
	plaintext, err = unpadCS7(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Pad adds PKCS#7 padding to the message.
func padCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Unpad removes PKCS#7 padding.
func unpadCS7(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	padding := int(data[len(data)-1])
	if padding == 0 || padding > blockSize {
		return nil, fmt.Errorf("invalid padding")
	}
	for _, p := range data[len(data)-padding:] {
		if int(p) != padding {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padding], nil
}
