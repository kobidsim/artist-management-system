package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	db *sql.DB
}

type AuthenticationService interface {
	Login(params view.LoginView) (map[string]interface{}, error)
	Register(params view.RegisterView) error
	Logout(token string) error
}

func NewAuthenticationService(db *sql.DB) AuthenticationService {
	return authenticationService{
		db: db,
	}
}

func (service authenticationService) Login(params view.LoginView) (map[string]interface{}, error) {
	query := `
		SELECT id, email, role, password FROM user WHERE email = $1
	`
	var user domain.User
	if err := service.db.QueryRow(query, params.Email).Scan(&user.ID, &user.Email, &user.Role, &user.Password); err != nil {
		fmt.Printf("ERROR::getting data from db:: %s\n", err.Error())
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		fmt.Printf("ERROR::password does not match:: %s\n", err.Error())
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
			"iat":   time.Now().Unix(),
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("no secret set in env file")
	}

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("ERROR::error signing token:: %s\n", err.Error())
		return nil, err
	}

	response := map[string]interface{}{
		"token": signedToken,
		"role":  user.Role,
	}

	return response, err
}

func (service authenticationService) Register(params view.RegisterView) error {
	dob, err := time.Parse("2006-01-02T15:04:05.000Z", params.DOB)
	if err != nil {
		return err
	}

	createdAt := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	query := `
		INSERT INTO user (first_name, last_name, role, email, password, phone, gender, address, dob, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 15)
	if err != nil {
		return err
	}
	if _, err := service.db.Exec(query, params.FirstName, params.LastName, params.Role,
		params.Email, string(hashedPassword), params.PhoneNumber, params.Gender, params.Address, dob, createdAt); err != nil {
		return err
	}

	return nil
}

func (service authenticationService) Logout(token string) error {
	query := `
		INSERT INTO invalid_tokens
		VALUES ($1);
	`

	if _, err := service.db.Exec(query, &token); err != nil {
		return err
	}

	return nil
}
