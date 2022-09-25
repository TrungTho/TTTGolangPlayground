package main

import (
	"encoding/json"
	"fmt"

	openssl "github.com/Luzifer/go-openssl/v4"
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

	data := Resp{UserId: 1, UserName: "golang"}
	text, _ := json.Marshal(data)

	fmt.Println("data to encrypt: ", string(text))
	e, _ := Encrypt(string(text), secretKey)
	fmt.Println("encrypted:", e)
	d, _ := Decrypt("U2FsdGVkX18UFhbCo6OQBy6mOvJpd9q1YFKTsFWWrIO1YXbmjAv4s1e98voS/9Kfy796/4+NGERc54QDIjd6jA==", secretKey)
	fmt.Println("decrypted: ", d)
}

const secretKey = "Sqrb[1R.1#.a~Kl5sdTM|6Z'65zhBi}~"

// Encrypt encrypts plain text string into cipher text string
func Encrypt(data string, secretKey string) (string, error) {
	o := openssl.New()

	enc, err := o.EncryptBytes(secretKey, []byte(data), openssl.BytesToKeyMD5)
	if err != nil {
		return "", err
	}

	return string(enc), nil

}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string, secretKey string) (string, error) {
	o := openssl.New()

	dec, err := o.DecryptBytes(secretKey, []byte(encrypted), openssl.BytesToKeyMD5)
	if err != nil {
		return "", err
	}

	return string(dec), nil

}
