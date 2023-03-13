package main

import (
	"github.com/he-man-david/simple-encryption/encryption"
	"github.com/he-man-david/simple-encryption/helper"
)

func main() {
	e := encryption.NewEncryptionSrvc()

	cmd := helper.StringPrompt("What do you want to do? Enter 'encrypt' or 'decrypt' -> ")

	switch {
	case cmd == "encrypt":
		e.RunEncryptor()
	case cmd == "decrypt":
		e.RunDecryptor()
	default:
		panic("Invalid entry, enter 'encrypt' or 'decrypt' only")
	}
}
