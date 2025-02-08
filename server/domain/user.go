package domain

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Role      string
	Email     string
	Password  string
	Phone     string
	Dob       time.Time
	Gender    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
