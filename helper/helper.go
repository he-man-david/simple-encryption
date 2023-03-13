package helper

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

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

func FillPadding(text string, size int) string {
    pad := size - len(text)
    for i := 0; i < pad; i++ {
        text += "_"
    }
    return text
} 

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+"")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}