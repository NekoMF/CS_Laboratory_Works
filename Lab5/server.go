package api

import (
	"fmt"

	"github.com/Marcel-MD/cs-labs/api/dto"
	"github.com/Marcel-MD/cs-labs/api/mfa"
	"github.com/Marcel-MD/cs-labs/api/middleware"
	"github.com/Marcel-MD/cs-labs/classic/caesar"
	"github.com/Marcel-MD/cs-labs/classic/playfair"
	"github.com/Marcel-MD/cs-labs/classic/vigenere"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ListenAndServe() error {
	godotenv.Load()

	s := NewUserService()
	otp := mfa.NewOtpService()
	es := mfa.NewMailService()
	e := gin.Default()
	g := e.Group("/api")

	g.GET("/users", func(c *gin.Context) {
		users, err := s.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, users)
	})

	g.GET("/users/:email", func(c *gin.Context) {
		email := c.Param("email")

		user, err := s.Get(email)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, user)
	})

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

	g.POST("/users/login", func(c *gin.Context) {
		var dto dto.CreateUser

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := otp.Verify(dto.Email, dto.Otp); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		jwt, err := s.Login(dto)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": jwt})
	})

	j := g.Use(middleware.JwtAuth())

	j.POST("/caesar/encrypt", func(c *gin.Context) {
		var dto dto.Caesar

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		cypher := caesar.Encrypt(dto.Alphabet, dto.Shift, dto.Text)

		dto.Text = cypher

		c.JSON(200, dto)
	})

	j.POST("/caesar/decrypt", func(c *gin.Context) {
		var dto dto.Caesar

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		plain := caesar.Decrypt(dto.Alphabet, dto.Shift, dto.Text)

		dto.Text = plain

		c.JSON(200, dto)
	})

	j.POST("/playfair/encrypt", func(c *gin.Context) {
		var dto dto.Playfair

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		cypher := playfair.Encrypt(dto.Key, dto.Text)

		dto.Text = cypher

		c.JSON(200, dto)
	})

	j.POST("/playfair/decrypt", func(c *gin.Context) {
		var dto dto.Playfair

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		plain := playfair.Decrypt(dto.Key, dto.Text)

		dto.Text = plain

		c.JSON(200, dto)
	})

	a := j.Use(middleware.HasRole(RoleAdmin))

	a.POST("/vigenere/encrypt", func(c *gin.Context) {
		var dto dto.Vigenere

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		cypher := vigenere.Encrypt(dto.Alphabet, dto.Key, dto.Text)

		dto.Text = cypher

		c.JSON(200, dto)
	})

	a.POST("/vigenere/decrypt", func(c *gin.Context) {
		var dto dto.Vigenere

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		plain := vigenere.Decrypt(dto.Alphabet, dto.Key, dto.Text)

		dto.Text = plain

		c.JSON(200, dto)
	})

	return e.Run(":8080")
}
