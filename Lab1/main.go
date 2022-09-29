package main

import (
	"log"

	"github.com/NekoMF/CS_Laboratory_Works/useful"
)

func main() {

	str := "I n   st r u m e      nt     s"
	log.Println(" This is the string before encryption : ", str)
	//key := "Monarchy"

	//playfair.Encrypt(str, key)

	str = useful.CutBlankSpaces(str)

	log.Println(" This is the string after encryption : ", str)
}
