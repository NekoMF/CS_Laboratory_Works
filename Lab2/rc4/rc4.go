package rc4

import (
	"errors"
	"strconv"
)

// XORKeyStream sets dst to the result of XORing src with the key stream.
// The key argument should be the RC4 key, at least 1 byte and at most 256 bytes.
func XORKeyStream(dst, src, key []byte) error {

	var s [256]uint32 // stream
	var i, j uint8

	k := len(key)
	if k < 1 || k > 256 {
		return errors.New("invalid key size " + strconv.Itoa(k))
	}

	for i := 0; i < 256; i++ {
		s[i] = uint32(i)
	}

	for i := 0; i < 256; i++ {
		j += uint8(s[i]) + key[i%k]
		s[i], s[j] = s[j], s[i]
	}

	if len(src) == 0 {
		return nil
	}

	_ = dst[len(src)-1]
	dst = dst[:len(src)] // eliminate bounds check from loop

	for k, v := range src {
		i += 1
		x := s[i]
		j += uint8(x)
		y := s[j]
		s[i], s[j] = y, x
		dst[k] = v ^ uint8(s[uint8(x+y)])
	}

	return nil
}
