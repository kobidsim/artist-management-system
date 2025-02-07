package view

type RegisterView struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,len=10"`
	DOB         string `json:"dob"`
	Gender      string `json:"gender" validate:"oneof=m f o"`
	Address     string `json:"address"`
	Role        string `json:"role" validate:"oneof=super_admin artist_manager artist"`
	Password    string `json:"password" validate:"required,min=6,max=72"`
}

type LoginView struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=72"`
}
