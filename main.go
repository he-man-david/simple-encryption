package main

import (
	"github.com/he-man-david/simple-encryption/encryption"
	"github.com/he-man-david/simple-encryption/helper"
)

func main() {
	e := encryption.NewEncryptionSrvc()

	cmd := helper.StringPrompt("What do you want to do? Enter 'encrypt' or 'decrypt' > ")

	switch {
	case cmd == "encrypt":
		e.RunEncryptor()
	case cmd == "decrypt":
		e.RunDecryptor()
	default:
		panic("Invalid entry, enter 'encrypt' or 'decrypt' only > ")
	}
}

// Example encrypted message _E_
// bH2vsI4IHWVjSiDeRHX7W9Cp+09jmgLlsiq8Rn8sq/X3EbPWFJB28hkYCsAH8+YKdSOH92brXyn3pO0gY6A/+nqd2VTPIySktxYidePxdJQX73FPnH9wuwinNaYrg0ICtxSiHxN86ylVsflkHbOOh8DEIepImN5ttNMxBob7AvU3Ovz4g/xc2RRSHMwW19AqqBntmANjVghUxW1XoqVYgp/+mtJMFKzI7xcPF0+utE+6K42CTsEN3kpLK5dtoqAdFjyCIqqjuu1Yqr7UR6o6mbqp9NDDIiCyjhzC9H/J1tt6esQSAD4xRwlk8hMaksv+lch1ivMKrST90D4wfoZxaog5e2RM3EQtLinOdmx6bDEHqsMYDF0lykS3yUNRi8QKYYk0bXrbRquxbldBbV+n91jH1+vD9vvHqjSB6udiZVCJllgRwSHlcAirD6yocgwUb5DQdrCSnULjhjUwzvy6K4wBeerp31GMu6/ykvSQYi52Mntk2AlatAX8z0+VX7xciK6GOg5OvVB3ZY7FfslnBGNkn4Ej2BP+tUSOopwuSkOgbb8VA+srQ9ik9nrynOiKMbZ7ln79DpDmNCdAT7ak+qqB5lURPmeIhNhNhNvJ6Jc3bE8e05ebiHmsjFUYCQuHtbRQDwtFcY8z3WzMEgCTC98H7x3Eh50DSu2nMHenXY8=