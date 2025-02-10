package domain

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Phone     string `json:"phone"`
	Dob       string `json:"-"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
