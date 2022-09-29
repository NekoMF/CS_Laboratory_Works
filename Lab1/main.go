package main

import (
	"log"

	"github.com/NekoMF/CS_Laboratory_Works/playfair"
)

func main() {

	str := "I n   st r u m e      nt     s"
	key := "monarchy"

	log.Println(" This is the string before encryption : ", str)

	playfairEncrypted := playfair.Encrypt(key, str)
	log.Println(" This is the string after playfair encryption : ", playfairEncrypted)

}
