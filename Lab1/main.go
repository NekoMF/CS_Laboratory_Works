package main

import (
	"log"
	"strings"

	"github.com/NekoMF/CS_Laboratory_Works/playfair"
	"github.com/NekoMF/CS_Laboratory_Works/useful"
	"github.com/NekoMF/CS_Laboratory_Works/vigenere"
)

func main() {

	str := useful.CutBlankSpaces("I n   st r u m e      nt     s")

	key := "monarchy"

	log.Println(" This is the string before encryption : ", str)

	playfairEncrypted := playfair.Encrypt(key, str)
	log.Println("\n This is the string after playfair encryption : ", strings.ToLower(playfairEncrypted))

	playfairEncrypted = vigenere.Encrypt(key, str)
	log.Println(" This is the string after vigenere encryption : ", strings.ToLower(playfairEncrypted))

}
