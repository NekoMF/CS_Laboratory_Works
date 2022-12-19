package api

type User struct {
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Role     string `json:"role"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
