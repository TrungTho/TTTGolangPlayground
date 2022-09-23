package utils

import (
	"fmt"
	"math/rand"
	"playground/constants"
	"time"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789~!@#$%^&*()_+{}|[]:;'<>?,./")
var noDigits = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+{}|[]:;'<>?,./")

func RandKeyStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func RandSaltStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(noDigits))]
	}
	return string(b)
}

func FakeKey() {
	rand.Seed(time.Now().Unix())
	firstNumber := ""
	numberOfSalt := rand.Intn(5) + 1
	numberOfFlipChars := rand.Intn(9) + 1

	if rand.Intn(2) == 1 {
		firstNumber = fmt.Sprintf("%d", rand.Intn(9)+1)
	}

	salt := RandSaltStr(numberOfSalt)
	prefixKey := RandKeyStr(numberOfFlipChars)
	suffixKey := RandKeyStr(constants.CryptoKeyLength - numberOfFlipChars)

	//real key to store redis
	key := prefixKey + suffixKey

	//fakeKey send to FE
	fakeKey := fmt.Sprintf("%s%s%d%s%s", firstNumber, salt, numberOfFlipChars, suffixKey, prefixKey)

	fmt.Printf("firstNum: \t%s\nsalt:   \t%s\nprefix: \t%s\nsuffix: \t%s\n", firstNumber, salt, prefixKey, suffixKey)
	fmt.Println("key: \n", key)
	fmt.Println("fake key: \n", fakeKey)
}
