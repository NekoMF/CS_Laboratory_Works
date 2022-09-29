package caesar

import (
	"math/rand"
	"strings"
)

func Encrypt(alphabet string, shift int, text string) string {
	var result strings.Builder
	for _, c := range text {
		if c == ' ' {
			result.WriteByte(' ')
			continue
		}

		index := strings.Index(alphabet, string(c))
		if index == -1 {
			panic("invalid character")
		}

		result.WriteByte(alphabet[(index+shift)%len(alphabet)])
	}
	return result.String()
}

func Decrypt(alphabet string, shift int, text string) string {
	return Encrypt(alphabet, len(alphabet)-shift, text)
}

func PermEncrypt(alphabet string, seed int64, shift int, text string) string {

	newAlphabet := permute(alphabet, seed)

	return Encrypt(string(newAlphabet), shift, text)
}

func PermDecrypt(alphabet string, seed int64, shift int, text string) string {

	newAlphabet := permute(alphabet, seed)

	return Decrypt(string(newAlphabet), shift, text)
}

func permute(alphabet string, seed int64) string {
	a := []rune(alphabet)

	rand.Seed(seed)
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return string(a)
}
