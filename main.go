package main

import (
	"fmt"

	e "github.com/he-man-david/simple-encryption/encryption"
)

func main() {
	var cmd int

	fmt.Println("Welcome to my encryption project.....\n", 
	"What do you want to do? Enter 1 for encrpt, 2 for decrypt.")
	fmt.Scanln(&cmd)

	switch {
	case cmd == 1:
		e.RunEncryptor()
	case cmd == 2:
		e.RunDecryptor()
	default:
		panic("Invalid entry, please enter 1 or 2 only!")
	}
}
