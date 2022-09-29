# Intro to Cryptography. Classical ciphers. Caesar cipher.

### Course: Cryptography & Security

### Author: Marcel Vlasenco

---

## Theory

&ensp;&ensp;&ensp; Caesar cipher has a key which is used to substitute the characters with the next ones, by the order number in a pre-established alphabet. Mathematically it would be expressed as follows:

$em = enc_{k}(x) = x + k (mod \; n),$

$dm = dec_{k}(x) = x + k (mod \; n),$

where:

- em: the encrypted message,
- dm: the decrypted message (i.e. the original one),
- x: input,
- k: key,
- n: size of the alphabet.

&ensp;&ensp;&ensp; Judging by the encryption mechanism one can conclude that this cipher is pretty easy to break. In fact, a brute force attack would have **_O(nm)_** complexity, where **_n_** would be the size of the alphabet and **_m_** the size of the message. This is why there were other variations of this cipher, which are supposed to make the cryptanalysis more complex.

## Objectives:

Implement 4 types of the classical ciphers:

- Caesar cipher with one key used for substitution (as explained above)
- Caesar cipher with one key used for substitution, and a permutation of the alphabet
- Vigenere cipher
- Playfair cipher

