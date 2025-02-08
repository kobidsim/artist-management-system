package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"database/sql"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	db *sql.DB
}

type AuthenticationService interface {
	Login(params view.LoginView) (string, error)
	Register(params view.RegisterView) error
}

func NewAuthenticationService(db *sql.DB) AuthenticationService {
	return authenticationService{
		db: db,
	}
}

func (service authenticationService) Login(params view.LoginView) (string, error) {
	query := `
		SELECT id, email, role, password FROM user WHERE email = $1
	`
	var user domain.User
	if err := service.db.QueryRow(query, params.Email).Scan(&user.ID, &user.Email, &user.Role, &user.Password); err != nil {
		fmt.Printf("ERROR::getting data from db:: %s\n", err.Error())
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		fmt.Printf("ERROR::password does not match:: %s\n", err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("ERROR::error signing token:: %s\n", err.Error())
		return "", err
	}

	return signedToken, err
}

func (service authenticationService) Register(params view.RegisterView) error {
	//TODO::need to find date time format compatible with client, server and db
	//so skipping date fields for now
	query := `
		INSERT INTO user (first_name, last_name, role, email, password, phone, gender, address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 15)
	if err != nil {
		return err
	}
	if _, err := service.db.Exec(query, params.FirstName, params.LastName, params.Role,
		params.Email, string(hashedPassword), params.PhoneNumber, params.Gender, params.Address); err != nil {
		return err
	}

	return nil
}
