package models

type User struct {
	ID        int64  `json:"id,omitempty" db:"id"`
	FirstName string `json:"first_name,omitempty" db:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name"`
	Email     string `json:"email,omitempty" db:"email"`
}

func NewUser(firstName string, lastName string, email string) *User {
	return &User{FirstName: firstName, LastName: lastName, Email: email}
}
