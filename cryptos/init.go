package cryptos

import (
	"encoding/json"
	"fmt"

	"playground/settings"

	"github.com/Luzifer/go-openssl/v4"
)

type Resp struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

func Demo() {
	data := Resp{UserId: 1, UserName: "golang"}
	text, _ := json.Marshal(data)

	fmt.Println("data to encrypt: ", string(text))
	e, _ := EncryptAES(string(text), settings.AES_KEY)
	fmt.Println("encrypted:", e)
	d, _ := DecryptAES("U2FsdGVkX18UFhbCo6OQBy6mOvJpd9q1YFKTsFWWrIO1YXbmjAv4s1e98voS/9Kfy796/4+NGERc54QDIjd6jA==", settings.AES_KEY)
	fmt.Println("decrypted: ", d)
}

// Encrypt encrypts plain text string into cipher text string
func EncryptAES(data string, secretKey string) (string, error) {
	o := openssl.New()

	enc, err := o.EncryptBytes(secretKey, []byte(data), openssl.BytesToKeyMD5)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}

// Decrypt decrypts cipher text string into plain text string
func DecryptAES(encrypted string, secretKey string) (string, error) {
	o := openssl.New()

	dec, err := o.DecryptBytes(secretKey, []byte(encrypted), openssl.BytesToKeyMD5)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}
