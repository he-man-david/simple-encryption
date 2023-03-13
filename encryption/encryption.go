package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
    iv = ""
    secret = ""
    message = ""
)

func RunEncryptor() {
    for {
        
    }
}

func RunDecryptor() {
    for {

    }
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
    MySecret = fillPadding(MySecret, 32)
    iv := fillPadding(MySecret, 16)
	// create new AES cipher block, with our secret
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	// create byte array format of string text
	plainText := []byte(text)
	// create encrypter interface using block, and iv
	cfb := cipher.NewCFBEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	// encode to string from byte array
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
    MySecret = fillPadding(MySecret, 32)
    iv := fillPadding(MySecret, 16)
	// create new AES cipher block, with our secret
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	// convert string to byte []
	cipherText := Decode(text)
	// create decrypter interface from block and iv
	cfb := cipher.NewCFBDecrypter(block, []byte(iv))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	// return decrypted message as string
	return string(plainText), nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func fillPadding(text string, size int) string {
    pad := size - len(text)
    for i := 0; i < pad; i++ {
        text += fmt.Sprint(i)
    }
    return text
} 