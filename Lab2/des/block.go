package des

import (
	"encoding/binary"
	"sync"
)

// Encrypt one block from src into dst, using the subkeys.
func encryptBlock(subkeys []uint64, dst, src []byte) {
	cryptBlock(subkeys, dst, src, false)
}

// Decrypt one block from src into dst, using the subkeys.
func decryptBlock(subkeys []uint64, dst, src []byte) {
	cryptBlock(subkeys, dst, src, true)
}

func cryptBlock(subkeys []uint64, dst, src []byte, decrypt bool) {
	b := binary.BigEndian.Uint64(src)
	b = permuteBlock(b, initialPermutation[:])
	left, right := uint32(b>>32), uint32(b)

	left = (left << 1) | (left >> 31)
	right = (right << 1) | (right >> 31)

	if decrypt {
		for i := 0; i < 8; i++ {
			left, right = feistel(left, right, subkeys[15-2*i], subkeys[15-(2*i+1)])
		}
	} else {
		for i := 0; i < 8; i++ {
			left, right = feistel(left, right, subkeys[2*i], subkeys[2*i+1])
		}
	}

	left = (left << 31) | (left >> 1)
	right = (right << 31) | (right >> 1)

	// switch left & right and perform final permutation
	preOutput := (uint64(right) << 32) | uint64(left)
	binary.BigEndian.PutUint64(dst, permuteBlock(preOutput, finalPermutation[:]))
}

// general purpose function to perform DES block permutations
func permuteBlock(src uint64, permutation []uint8) (block uint64) {
	for position, n := range permutation {
		bit := (src >> n) & 1
		block |= bit << uint((len(permutation)-1)-position)
	}
	return
}

// feistelBox[s][16*i+j] contains the output of permutationFunction
// for sBoxes[s][i][j] << 4*(7-s)
var feistelBox [8][64]uint32

// DES Feistel function.
func feistel(l, r uint32, k0, k1 uint64) (lout, rout uint32) {
	var t uint32

	t = r ^ uint32(k0>>32)
	l ^= feistelBox[7][t&0x3f] ^
		feistelBox[5][(t>>8)&0x3f] ^
		feistelBox[3][(t>>16)&0x3f] ^
		feistelBox[1][(t>>24)&0x3f]

	t = ((r << 28) | (r >> 4)) ^ uint32(k0)
	l ^= feistelBox[6][(t)&0x3f] ^
		feistelBox[4][(t>>8)&0x3f] ^
		feistelBox[2][(t>>16)&0x3f] ^
		feistelBox[0][(t>>24)&0x3f]

	t = l ^ uint32(k1>>32)
	r ^= feistelBox[7][t&0x3f] ^
		feistelBox[5][(t>>8)&0x3f] ^
		feistelBox[3][(t>>16)&0x3f] ^
		feistelBox[1][(t>>24)&0x3f]

	t = ((l << 28) | (l >> 4)) ^ uint32(k1)
	r ^= feistelBox[6][(t)&0x3f] ^
		feistelBox[4][(t>>8)&0x3f] ^
		feistelBox[2][(t>>16)&0x3f] ^
		feistelBox[0][(t>>24)&0x3f]

	return l, r
}

var feistelBoxOnce sync.Once

// creates 16 56-bit subkeys from the original key
func (c *desCipher) generateSubkeys(keyBytes []byte) {
	feistelBoxOnce.Do(initFeistelBox)

	// apply PC1 permutation to key
	key := binary.BigEndian.Uint64(keyBytes)
	permutedKey := permuteBlock(key, permutedChoice1[:])

	// rotate halves of permuted key according to the rotation schedule
	leftRotations := ksRotate(uint32(permutedKey >> 28))
	rightRotations := ksRotate(uint32(permutedKey<<4) >> 4)

	// generate subkeys
	for i := 0; i < 16; i++ {
		// combine halves to form 56-bit input to PC2
		pc2Input := uint64(leftRotations[i])<<28 | uint64(rightRotations[i])
		// apply PC2 permutation to 7 byte input
		c.subkeys[i] = unpack(permuteBlock(pc2Input, permutedChoice2[:]))
	}
}

func initFeistelBox() {
	for s := range sBoxes {
		for i := 0; i < 4; i++ {
			for j := 0; j < 16; j++ {
				f := uint64(sBoxes[s][i][j]) << (4 * (7 - uint(s)))
				f = permuteBlock(f, permutationFunction[:])

				// Row is determined by the 1st and 6th bit.
				// Column is the middle four bits.
				row := uint8(((i & 2) << 4) | i&1)
				col := uint8(j << 1)
				t := row | col

				// The rotation was performed in the feistel rounds, being factored out and now mixed into the feistelBox.
				f = (f << 1) | (f >> 31)

				feistelBox[s][t] = uint32(f)
			}
		}
	}
}

// creates 16 28-bit blocks rotated according
// to the rotation schedule
func ksRotate(in uint32) (out []uint32) {
	out = make([]uint32, 16)
	last := in
	for i := 0; i < 16; i++ {
		// 28-bit circular left shift
		left := (last << (4 + ksRotations[i])) >> 4
		right := (last << 4) >> (32 - ksRotations[i])
		out[i] = left | right
		last = out[i]
	}
	return
}

// Expand 48-bit input to 64-bit, with each 6-bit block padded by extra two bits at the top.
// By doing so, we can have the input blocks (four bits each), and the key blocks (six bits each) well-aligned without
// extra shifts/rotations for alignments.
func unpack(x uint64) uint64 {
	return ((x>>(6*1))&0xff)<<(8*0) |
		((x>>(6*3))&0xff)<<(8*1) |
		((x>>(6*5))&0xff)<<(8*2) |
		((x>>(6*7))&0xff)<<(8*3) |
		((x>>(6*0))&0xff)<<(8*4) |
		((x>>(6*2))&0xff)<<(8*5) |
		((x>>(6*4))&0xff)<<(8*6) |
		((x>>(6*6))&0xff)<<(8*7)
}
