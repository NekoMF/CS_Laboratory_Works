# Topic: Web Authentication & Authorisation.

### Course: Cryptography & Security
### Author: Marcel Vlasenco

----

## Overview

&ensp;&ensp;&ensp; Authentication & authorization are 2 of the main security goals of IT systems and should not be used interchangibly. Simply put, during authentication the system verifies the identity of a user or service, and during authorization the system checks the access rights, optionally based on a given user role.

&ensp;&ensp;&ensp; There are multiple types of authentication based on the implementation mechanism or the data provided by the user. Some usual ones would be the following:
- Based on credentials (Username/Password);
- Multi-Factor Authentication (2FA, MFA);
- Based on digital certificates;
- Based on biometrics;
- Based on tokens.

&ensp;&ensp;&ensp; Regarding authorization, the most popular mechanisms are the following:
- Role Based Access Control (RBAC): Base on the role of a user;
- Attribute Based Access Control (ABAC): Based on a characteristic/attribute of a user.


## Objectives:
1. Take what you have at the moment from previous laboratory works and put it in a web service / serveral web services.
2. Your services should have implemented basic authentication and MFA (the authentication factors of your choice).
3. Your web app needs to simulate user authorization and the way you authorise user is also a choice that needs to be done by you.
4. As services that your application could provide, you could use the classical ciphers. Basically the user would like to get access and use the classical ciphers, but they need to authenticate and be authorized. 

## Implementation:

### User

The user will be able to authenticate and use the services provided by the web application. The user will be able to use the classical ciphers, but only if they are authenticated and authorized.

```go
type User struct {
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Role     string `json:"role"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
```

### Web Service
The web service is implemented using *gin* framework.

For user registration we have the endpoint `api/users/register`. First we deserilize the dto, then we verify the OTP code based on user email address, and then we register the user.

```go
g.POST("/users/register", func(c *gin.Context) {
    var dto dto.CreateUser

    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := otp.Verify(dto.Email, dto.Otp); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := s.Register(dto); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, nil)
})
```

When registering a user we first check if it already exists, then we hash the password and save him in the database.

```go
func (s *userService) Register(dto dto.CreateUser) error {
	if _, ok := s.users[dto.Email]; ok {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.users[dto.Email] = User{Email: dto.Email, Password: hashedPassword, Role: RoleUser}

	return nil
}
```

When logging in, we get the user form the in-memory database, then we compare the password hash with the one from the database. If the password is correct, we generate a JWT token and return it.

```go
func (s *userService) Login(dto dto.CreateUser) (string, error) {
	user, err := s.Get(dto.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(dto.Password))
	if err != nil {
		return "", err
	}

	jwt, err := token.Generate(user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
```

### JWT Token

The JWT token is generated using the *jwt-go* library. The token is signed using the HMAC algorithm and the secret key is stored in the environment variables.

```go
func Generate(email, role string) (string, error) {

	secret := os.Getenv("SECRET")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(12 * time.Hour).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
```

### OTP

For two factor authentication is used one time passwords (OTP) that are sent to the user email address.

```go
g.POST("/otp/:email", func(c *gin.Context) {
    email := c.Param("email")

    pass, err := otp.Generate(email)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    mail := mfa.Mail{
        To:      []string{email},
        Subject: "Verification Code",
        Body:    fmt.Sprintf("Your verification code is <strong>%s</strong>.", pass),
    }

    go es.Send(mail)

    c.JSON(200, gin.H{"message": "Verification code sent to your email."})
})
```

The OTP is generated using the *mfa* module.

```go
func (s *OtpService) Generate(email string) (string, error) {
	num := 100000 + rand.Intn(800000)
	otp := strconv.Itoa(num)

	s.otps[email] = otp

	return otp, nil
}

func (s *OtpService) Verify(email string, otp string) error {
	if s.otps[email] != otp {
		return errors.New("otp is not valid")
	}

	delete(s.otps, email)

	return nil
}
```

### Authentication

The authentication is implemented by using and middleware that checks if the JWT token is valid and extracts user email and role and saves them in the request context.

```go
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, role, err := token.ExtractEmailRole(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("email", email)
		c.Set("role", role)
		c.Next()
	}
}
```

To use the authentication middleware we just need to add it to the route.

```go
	j := g.Use(middleware.JwtAuth())

	j.POST("/caesar/encrypt", func(c *gin.Context) {
		[...]
	})
```

### Authorization

The authorization is implemented by using and middleware that checks if the user has the required role.

```go
func HasRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleFromToken := c.GetString("role")
		if roleFromToken != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
```

To use the authorization middleware we just need to add it to the route.

```go
    j := g.Use(middleware.JwtAuth())

    [...]

    a := j.Use(middleware.HasRole("admin"))

    a.POST("/vigenere/encrypt", func(c *gin.Context) {
        [...]
    })
```

## Conclusion

In this laboratory work we implemented a REST API for encrypting and decrypting messages using classic ciphers. The API is secured by using JWT tokens and two factor authentication.