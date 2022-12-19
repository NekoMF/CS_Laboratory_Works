package des

import (
	"errors"
	"strconv"
)

// The DES block size in bytes.
const BlockSize = 8

type Cipher interface {
	Encrypt(dst, src []byte)
	Decrypt(dst, src []byte)
}

// desCipher is an instance of DES encryption.
type desCipher struct {
	subkeys [16]uint64
}

// NewCipher creates and returns a new cipher.Block.
func NewCipher(key []byte) (Cipher, error) {
	if len(key) != 8 {
		return nil, errors.New("invalid key size " + strconv.Itoa(len(key)))
	}

	c := new(desCipher)
	c.generateSubkeys(key)
	return c, nil
}

func (c *desCipher) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("input not full block")
	}
	if len(dst) < BlockSize {
		panic("output not full block")
	}
	encryptBlock(c.subkeys[:], dst, src)
}

func (c *desCipher) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("input not full block")
	}
	if len(dst) < BlockSize {
		panic("output not full block")
	}
	decryptBlock(c.subkeys[:], dst, src)
}
