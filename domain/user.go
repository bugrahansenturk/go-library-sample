package domain

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
