package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mergermarket/go-pkcs7"
)

type RewardServiceError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Resp struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

func main() {

	data := Resp{UserId: 1, UserName: "User2"}
	text, _ := json.Marshal(data)

	fmt.Println("data to encrypt: ", string(text))
	e, _ := Encrypt(string(text), secretKey)
	fmt.Println("encrypted:", e)
	d, _ := Decrypt("f7aa599aebc8939c092b8ea6392e2e184baf2245bdd77b800c7d0b179d72955a62323d72d3a1228c45caa932a2f2d254e44b8a16f7572a7221506e27b53edf56", secretKey)
	fmt.Println("decrypted: ", d)
}

const secretKey = "Sqrb[1R.1#.a~Kl5sdTM|6Z'65zhBi}~"

// Encrypt encrypts plain text string into cipher text string
func Encrypt(data string, secretKey string) (string, error) {
	key := []byte(secretKey)
	plainText := []byte(data)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string, secretKey string) (string, error) {
	key := []byte(secretKey)
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return fmt.Sprintf("%s", cipherText), nil
}
