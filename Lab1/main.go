package main

import (
	"log"
	"strings"

	"github.com/NekoMF/CS_Laboratory_Works/caesar"
	"github.com/NekoMF/CS_Laboratory_Works/playfair"
	"github.com/NekoMF/CS_Laboratory_Works/useful"
	"github.com/NekoMF/CS_Laboratory_Works/vigenere"
)

func main() {

	alfabet := "abcdefghijklmnopqrstuvwxyz"
	var s int64 = 69420
	permutation := 10
	str := "I   nst  rum     ents"
	key := "monarchy"

	str = strings.ToLower(useful.CutBlankSpaces(str))

	log.Println(" This is the string before encryption : ", str)

	caesarEncrypted := caesar.PermEncrypt(alfabet, s, permutation, str)
	log.Println("\nThis is the string after CAESAR encryption :  ", caesarEncrypted)

	caesarDecrypted := caesar.PermDecrypt(alfabet, s, permutation, caesarEncrypted)
	log.Println("\nDecrypted : ", caesarDecrypted)

	playfairEncrypted := playfair.Encrypt(key, str)
	log.Println("\n This is the string after playfair encryption : ", strings.ToLower(playfairEncrypted))
	playfairDecrypted := playfair.Decrypt(key, playfairEncrypted)
	log.Println("\n Decrypted : ", strings.ToLower(playfairDecrypted))

	vigenereEncrypted := vigenere.Encrypt(key, str)
	log.Println("\n This is the string after vigenere encryption : ", strings.ToLower(vigenereEncrypted))
	vigenereDecrypted := vigenere.Decrypt(key, vigenereEncrypted)
	log.Println("\n Decrypted : ", strings.ToLower(vigenereDecrypted))

	// caesarEncrypted := caesar.PermEncrypt(alfabet, s, permutations, str)
	// log.Println("\n This is the string after caesar permutation encryption : ", caesarEncrypted)

}
