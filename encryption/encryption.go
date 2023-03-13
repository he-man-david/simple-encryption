package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/he-man-david/simple-encryption/helper"
)
type encryptionSrvc struct {
    Message string
}

func NewEncryptionSrvc() *encryptionSrvc {
    return &encryptionSrvc{
        Message: "",
    }
}

func (e *encryptionSrvc) RunEncryptor() {
    e.Message = helper.StringPrompt("Please enter message that you want to encrypt -> ")
    for {
        secret := helper.StringPrompt("Please enter your secret password -> ")

        encryptedMsg, err := e.Encrypt(e.Message, secret)
        if err != nil {
            panic(err)
        }
        fmt.Println("Encrypted message -> ", encryptedMsg)
        e.Message = encryptedMsg
    }
}

func (e *encryptionSrvc) RunDecryptor() {
    e.Message = helper.StringPrompt("Please enter message that you want to decrypt -> ")
    for {
        secret := helper.StringPrompt("Please enter your secret password -> ")

        decryptedMsg, err := e.Decrypt(e.Message, secret)
        if err != nil {
            panic(err)
        }
        fmt.Println("Decrypted message -> ", decryptedMsg)
        e.Message = decryptedMsg
    }
}

// Encrypt method is to encrypt or hide any classified text
func (e *encryptionSrvc) Encrypt(text, secret string) (string, error) {
    // IV is like a nonce, ideally should store securely along with 
    // the secret, but for experimenting we just calc on the fly
    iv := helper.FillPadding(secret, 16)   
    secret = helper.FillPadding(secret, 32)
	// create new AES cipher block, with our secret
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	// create byte array format of string text
	plainText := []byte(text)
	// create encrypter interface using block, and e.iv
	cfb := cipher.NewCFBEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	// encode to string from byte array
	return helper.Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func (e *encryptionSrvc) Decrypt(text, secret string) (string, error) {
    // IV is like a nonce, ideally should store securely along with 
    // the secret, but for experimenting we just calc on the fly
    iv := helper.FillPadding(secret, 16)
    secret = helper.FillPadding(secret, 32)
	// create new AES cipher block, with our secret
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	// convert string to byte []
	cipherText := helper.Decode(text)
	// create decrypter interface from block and e.iv
	cfb := cipher.NewCFBDecrypter(block, []byte(iv))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	// return decrypted message as string
	return string(plainText), nil
}