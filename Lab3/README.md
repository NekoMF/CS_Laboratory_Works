# Topic: Asymmetric Ciphers.

### Course: Cryptography & Security
### Author: Savva Nicu

----

## Overview
&ensp;&ensp;&ensp; Asymmetric Cryptography (a.k.a. Public-Key Cryptography)deals with the encryption of plain text when having 2 keys, one being public and the other one private. The keys form a pair and despite being different they are related.

&ensp;&ensp;&ensp; As the name implies, the public key is available to the public but the private one is available only to the authenticated recipients. 

&ensp;&ensp;&ensp; A popular use case of the asymmetric encryption is in SSL/TLS certificates along side symmetric encryption mechanisms. It is necessary to use both types of encryption because asymmetric ciphers are computationally expensive, so these are usually used for the communication initiation and key exchange, or sometimes called handshake. The messages after that are encrypted with symmetric ciphers.


## Examples
1. RSA
2. Diffie-Helman
3. ECC
4. El Gamal
5. DSA


## Objectives:
1. Get familiar with the asymmetric cryptography mechanisms.

2. Implement an example of an asymmetric cipher.

3. As in the previous task, please use a client class or test classes to showcase the execution of your programs.

   
## Implementation Description:

### RSA

RSA is an asymmetric cipher that uses the fact that it is easy to find the prime factors of a large number, but it is difficult to find the prime factors of a product of two large prime numbers. It uses two keys, one public and one private. The public key is used to encrypt the message and the private key is used to decrypt the message.

**Key generation steps:**
1. Generate two large prime numbers `p` and `q`.
```go
	p, err := rand.Prime(r, bits/2)
	if err != nil {
		return
	}

	q, err := rand.Prime(r, bits/2)
	if err != nil {
		return
	}
```

2. Compute `n = p * q`.
```go
    n := new(big.Int).Mul(p, q)
```

3. Compute `phi(n) = (p - 1) * (q - 1)`.
```go
    phi := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))
```

4. Choose an integer `e` such that `1 < e < phi(n)` and `gcd(e, phi(n)) = 1`. This is the public key.
```go
    e := randE(r, phi)
	
	[...]

	func randE(r io.Reader, phi *big.Int) *big.Int {
		e, _ := rand.Int(r, phi)
		for {
			if new(big.Int).GCD(nil, nil, e, phi).Cmp(big.NewInt(1)) == 0 {
				return e
			}
			e, _ = rand.Int(r, phi)
		}
	}
```

5. Compute `d` such that `d * e = 1 mod phi(n)`. This is the private key.
```go
    d := new(big.Int).ModInverse(e, phi)
```

**Encryption:**

To compute cypher text `c` from plain text `m` we use the following formula: `c = m^e mod n`.
```go
    c := new(big.Int).Exp(m, e, n)
```

**Decryption:**

To compute plain text `m` from cypher text `c` we use the following formula: `m = c^d mod n`.
```go
    m := new(big.Int).Exp(c, d, n)
```

## Conclusion:

After performing this laboratory work I have learned how to implement an asymmetric cipher and how it works. RSA is a very popular asymmetric cipher and it is used in many applications for encryption and decryption of data. It can be slow and it is not suitable for large amounts of data, but it is very secure.