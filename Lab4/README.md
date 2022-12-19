# Topic: Hash functions and Digital Signatures.

### Course: Cryptography & Security
### Author: Savva Nicu

----

## Overview
&ensp;&ensp;&ensp; Hashing is a technique used to compute a new representation of an existing value, message or any piece of text. The new representation is also commonly called a digest of the initial text, and it is a one way function meaning that it should be impossible to retrieve the initial content from the digest.

&ensp;&ensp;&ensp; Such a technique has the following usages:
  * Offering confidentiality when storing passwords,
  * Checking for integrity for some downloaded files or content,
  * Creation of digital signatures, which provides integrity and non-repudiation.

&ensp;&ensp;&ensp; In order to create digital signatures, the initial message or text needs to be hashed to get the digest. After that, the digest is to be encrypted using a public key encryption cipher. Having this, the obtained digital signature can be decrypted with the public key and the hash can be compared with an additional hash computed from the received message to check the integrity of it.


## Examples
1. Argon2
2. BCrypt
3. MD5 (Deprecated due to collisions)
4. RipeMD
5. SHA256 (And other variations of SHA)
6. Whirlpool


## Objectives:
1. Get familiar with the hashing techniques/algorithms.
2. Use an appropriate hashing algorithms to store passwords in a local DB.
    1. You can use already implemented algortihms from libraries provided for your language.
    2. The DB choise is up to you, but it can be something simple, like an in memory one.
3. Use an asymmetric cipher to implement a digital signature process for a user message.
    1. Take the user input message.
    2. Preprocess the message, if needed.
    3. Get a digest of it via hashing.
    4. Encrypt it with the chosen cipher.
    5. Perform a digital signature check by comparing the hash of the message with the decrypted one.

   
## Implementation Description:

&ensp;&ensp;&ensp; The laboratory implementation is separated in two packages `user` and `message`. In `user` package we have an in-memory-database that will store the users. The `User` struct contains the user email, password and private key. The password is stored as a byte array, because it is the output of the hashing algorithm. The private key is used to decrypt the digital signature.

```go
type User struct {
	Email      string
	Password   []byte
	PrivateKey *rsa.PrivateKey
}
```

&ensp;&ensp;&ensp; `UserService` has the methods for registering a new user and logging in. The `Register` method will generate a new private key for the user and store it in the database. The `Login` method will check if the user exists in the database and if the password is correct.

```go
func (s *userService) Register(email, password string) error {

	_, err := s.db.Get(email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	privKey, err := rsa.GenerateKey(rand.Reader, 1028)
	if err != nil {
		return err
	}

	user := User{
		Email:      email,
		Password:   hashedPassword,
		PrivateKey: privKey,
	}

	return s.db.Set(email, user)
}

func (s *userService) Login(email, password string) error {
	user, err := s.db.Get(email)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
```

&ensp;&ensp;&ensp; The `message` package contains the `MessageService`. It has the methods for signing a message and verifying the signature. The hashing algorithm used is SHA256. The `SignMessage` method will hash the message, encrypt it with the user private key and return the signature. The `VerifyMessage` method will decrypt the signature with the user public key, hash the message and compare the two hashes.

```go
func (s *messageService) SignMessage(privateKey *rsa.PrivateKey, msg []byte) ([]byte, []byte, error) {

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil, nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return nil, nil, err
	}

	return signature, msgHashSum, nil
}

func (s *messageService) VerifyMessage(publicKey *rsa.PublicKey, signature []byte, msgHashSum []byte) error {

	return rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
}
```

## Conclusion:

&ensp;&ensp;&ensp; Hashing is a very useful technique that can be used to store passwords, check for integrity and create digital signatures. The hashing algorithms are very fast and can be used to check for integrity of large files. The digital signature process is very useful for non-repudiation and integrity. The digital signature can be used to verify the authenticity of a message and to check if it was modified.