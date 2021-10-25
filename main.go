package main

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type RCipher struct {
	IV    string
	Value string
	mac   string
}

var key []byte
var text []byte
var err error

func decode(textString string) {
	var jsonDecode RCipher
	text, err = b64.StdEncoding.DecodeString(textString)
	if err != nil {
		fmt.Println(err)
	}
	err := json.Unmarshal(text, &jsonDecode)
	if err != nil {
		fmt.Println(err)
	}
	ciphertext, _ := b64.StdEncoding.DecodeString(jsonDecode.Value)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	iv, _ := b64.StdEncoding.DecodeString(jsonDecode.IV)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	fmt.Println(string(ciphertext))
}

func setKey(keyString string) {
	key, err = b64.StdEncoding.DecodeString(keyString)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	start := time.Now()
	//key from .env
	//setKey(key_from_env)
	//text to decode
	// decode(text_to_decode)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}
