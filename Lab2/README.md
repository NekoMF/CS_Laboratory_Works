# Topic: Symmetric Ciphers. Stream Ciphers. Block Ciphers.

### Course: Cryptography & Security

### Author: Savva Nicu

---

## Overview

&ensp;&ensp;&ensp; Symmetric Cryptography deals with the encryption of plain text when having only one encryption key which needs to remain private. Based on the way the plain text is processed/encrypted there are 2 types of ciphers:

- Stream ciphers:
  - The encryption is done one byte at a time.
  - Stream ciphers use confusion to hide the plain text.
  - Make use of substitution techniques to modify the plain text.
  - The implementation is fairly complex.
  - The execution is fast.
- Block ciphers:
  - The encryption is done one block of plain text at a time.
  - Block ciphers use confusion and diffusion to hide the plain text.
  - Make use of transposition techniques to modify the plain text.
  - The implementation is simpler relative to the stream ciphers.
  - The execution is slow compared to the stream ciphers.

## Objectives

1. Get familiar with the symmetric cryptography, stream and block ciphers.

2. Implement an example of a stream cipher RC4.

3. Implement an example of a block cipher DES.

## Implementation description

### RC4 stream cipher

RC4 is a stream cipher and variable-length key algorithm. This algorithm encrypts one byte at a time. A key input is pseudorandom bit generator that produces a stream 8-bit number that is unpredictable without knowledge of input key, The output of the generator is called key-stream, is combined one byte at a time with the plaintext stream cipher using X-OR operation.

**Algorithm:**

1. Initialize the permutation in the array S:

```go
	for i := 0; i < 256; i++ {
		s[i] = uint32(i)
	}

	for i := 0; i < 256; i++ {
		j += uint8(s[i]) + key[i%k]
		s[i], s[j] = s[j], s[i]
	}
```

2. While generating keystream:

   1. Update `i = (i + 1) mod 256`.

   2. Update `j = (j + S[i]) mod 256`.

   3. Swap `S[i]` and `S[j]`.

   4. Generate keystream byte `K = S[(S[i] + S[j]) mod 256]`.

   5. Encrypt the plaintext byte with the keystream byte `K` using X-OR operation.

```go
	for k, v := range src {
		i += 1
		x := s[i]
		j += uint8(x)
		y := s[j]
		s[i], s[j] = y, x
		dst[k] = v ^ uint8(s[uint8(x+y)])
	}
```

### DES block cipher

DES is a block cipher and fixed-length key algorithm. This algorithm encrypts one block of 64 bits at a time. The input is a 64-bit block of plaintext and a 64-bit key. The output is a 64-bit block of ciphertext. The key is divided into two 28-bit halves, C0 and D0. The key is then rotated 1 bit to the left 16 times. The first 28 bits of the key are used to generate 16 subkeys, C1, C2, ..., C16. The last 28 bits of the key are used to generate 16 subkeys, D1, D2, ..., D16. The 16 subkeys are then combined to form 16 48-bit subkeys, K1, K2, ..., K16. The 16 subkeys are used to encrypt the plaintext block.

**Algorithm:**

1. Initial permutation:

   1. The 64-bit plaintext block is permuted using the initial permutation table.

   2. The 64-bit ciphertext block is permuted using the inverse initial permutation table.

```go
  // general purpose function to perform DES block permutations
  func permuteBlock(src uint64, permutation []uint8) (block uint64) {
    for position, n := range permutation {
      bit := (src >> n) & 1
      block |= bit << uint((len(permutation)-1)-position)
    }
    return
  }
```

2. 16 rounds:

   1. The 32-bit right half of the plaintext block is expanded to 48 bits using the expansion table.

   2. The 48-bit expanded right half is XORed with the 48-bit subkey.

   3. The 48-bit result is divided into 8 6-bit blocks.

   4. Each 6-bit block is used to index into the S-boxes to produce a 4-bit number.

   5. The 32-bit result is permuted using the permutation table.

   6. The 32-bit result is XORed with the 32-bit left half of the plaintext block.

   7. The 32-bit left half of the plaintext block is replaced with the 32-bit right half of the plaintext block.

   8. The 32-bit right half of the plaintext block is replaced with the 32-bit result of the XOR operation.

3. Final permutation:

   1. The 32-bit left half of the plaintext block is swapped with the 32-bit right half of the plaintext block.

   2. The 64-bit plaintext block is permuted using the final permutation table.

## Run tests

```sh
$ go test ./...
```

## Conclusions

After performing this laboratory work I have learned the complexity of the symmetric cryptography, stream and block ciphers. I have implemented an example of a stream cipher RC4 and an example of a block cipher DES. The difference between the stream and block ciphers is that the stream ciphers encrypt one byte at a time and the block ciphers encrypt one block of plain text at a time. The stream ciphers use confusion to hide the plain text and make use of substitution techniques to modify the plain text. The block ciphers use confusion and diffusion to hide the plain text and make use of transposition techniques to modify the plain text. The implementation of the stream ciphers is fairly complex and the execution is fast. The implementation of the block ciphers is simpler relative to the stream ciphers and the execution is slow compared to the stream ciphers.
