package encryption

import (
	"testing"
)

var (
	messageA = "h0iuI+XUnhil1CKdzX0nR8KdP+3cVzD7P4N7wR7wbgLBH/aUNcT0cHGZO+ULj4ZDHmCNFg8I9URdy/Y0qHN7KhxBclk4XB94FekqzD9PN9289DfwoZEvhh1Y04xny8QZRJIX+2S1w17VV1UMWC1LWBMHD3fpBNN83j7tWI3KtOJw9v0EofyKv8uBTUQvtuGgogr3Y/VL+R0F1xXFg5f2tj9mo4TlEOHsiOZu786177emCuRX77qkt8YoKgR9lCB1AD7//dw6Jpj/CYf/fG9g51EZ2ytFDu9RZoWzMireOKcGbUWXnH3f/Hj9UbcT+5baJX6/KtO330fDSlFcBT4tz3twr79p0PbOi1I6DxkG5YHhcWzaEiBlBP9gAucmMdfUOTLIPundNiYiMrWLwSPK9iqbRHQeiHooO9A94vQ5CmhrKsItGOb0XlsCiF1UdOXsJiY7f39ZakOxjuMtHQLpPBiDt4ljPwSH4Twwb+eICQQZ2KL0SY97ffYhLCOeIGk1r4sYrRXm0x8Mvw5VOor4560+zuszlYFp1zv1Vf1NlKASVdQgHv11Hmk373toi3TlLznBHp2S34hgs6CPJH4+FlOL2lo3x0OmqFO888TokyCybPo8tknnBYSBvbIXNferOhkaV6njPan+zKl8L2ipHt8SLHfpnEGCWZYUaJMwaRgfRLK6dZs/eiNmtO5BUxyMQq1/+9BW6ZFMS4JAyYp+qlHfp6UXRqop6xyCz65v21qT6dSiKXAnL6e5ubZh+oebFLbiH+TjcYRqAGN4LZ3R+R5tLUFaR8ygT3NAnR308vQmtKLl7xok5g=="
)

func TestEncryption(t *testing.T) {
	e := NewEncryptionSrvcTester()

	encryptedMsg, err := e.Encrypt(messageA)
	if err != nil {
		t.Errorf("Failed to encrypt message! %s", messageA)
	}

	decryptedMsg, err := e.Decrypt(encryptedMsg)
	if err != nil {
		t.Errorf("Failed to decrypt message! %s", encryptedMsg)
	}

	if decryptedMsg != messageA {
		t.Errorf("Decrypted msg not the same as original msg. \n Decrypted: %s \n Original: %s", decryptedMsg, messageA)
	}
}