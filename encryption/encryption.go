package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/he-man-david/simple-encryption/helper"
)
type EncryptionSrvc struct {
    iv string
    secret string
    message string
}

func NewEncryptionSrvc() *EncryptionSrvc {
    return &EncryptionSrvc{
        iv: "",
        secret: "",
        message: "",
    }
}

func (e *EncryptionSrvc) RunEncryptor() {
    for {
        e.message = helper.StringPrompt("Please enter message that you want to encrypt > ")
        e.secret  = helper.StringPrompt("Please enter your secret password > ")

        encryptedMsg, err := e.encrypt(e.message)
        if err != nil {
            panic(err)
        }
        fmt.Println("Encrypted message > ", encryptedMsg)
    }
}

func (e *EncryptionSrvc) RunDecryptor() {
    for {
        e.message = helper.StringPrompt("Please enter message that you want to decrypt > ")
        e.secret  = helper.StringPrompt("Please enter your secret password > ")

        decryptedMsg, err := e.decrypt(e.message)
        if err != nil {
            panic(err)
        }
        fmt.Println("Decrypted message > ", decryptedMsg)
    }
}

// Encrypt method is to encrypt or hide any classified text
func (e *EncryptionSrvc) encrypt(text string) (string, error) {
    e.iv = helper.FillPadding(e.secret, 16)
    e.secret = helper.FillPadding(e.secret, 32)
	// create new AES cipher block, with our e.secret
	block, err := aes.NewCipher([]byte(e.secret))
	if err != nil {
		return "", err
	}
	// create byte array format of string text
	plainText := []byte(text)
	// create encrypter interface using block, and e.iv
	cfb := cipher.NewCFBEncrypter(block, []byte(e.iv))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	// encode to string from byte array
	return helper.Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func (e *EncryptionSrvc) decrypt(text string) (string, error) {
    e.iv = helper.FillPadding(e.secret, 16)
    e.secret = helper.FillPadding(e.secret, 32)
	// create new AES cipher block, with our e.secret
	block, err := aes.NewCipher([]byte(e.secret))
	if err != nil {
		return "", err
	}
	// convert string to byte []
	cipherText := helper.Decode(text)
	// create decrypter interface from block and e.iv
	cfb := cipher.NewCFBDecrypter(block, []byte(e.iv))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	// return decrypted message as string
	return string(plainText), nil
}