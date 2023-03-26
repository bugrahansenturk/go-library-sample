package domain

type Role string

const (
	Member  Role = "member"
	Expired Role = "expired"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      Role   `json:"role"`
}
