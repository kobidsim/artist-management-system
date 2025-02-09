package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	db *sql.DB
}

type UserService interface {
	All() ([]domain.User, error)
	Create(params view.UserView) error
	Update(id int, params view.UserView) error
}

func NewUserService(db *sql.DB) UserService {
	return userService{
		db: db,
	}
}

func (service userService) All() ([]domain.User, error) {
	query := `
		SELECT id, first_name, last_name, role, email, phone, gender, address FROM user;
	`

	var users []domain.User
	rows, err := service.db.Query(query)
	if err != nil {
		fmt.Printf("ERROR:: could not query database: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Role, &user.Email, &user.Phone, &user.Gender, &user.Address); err != nil {
			fmt.Printf("ERROR:: could not scan values from row: %s\n", err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("ERROR:: could not scan next row: %s\n", err.Error())
		return nil, err
	}

	return users, nil
}

func (service userService) Create(params view.UserView) error {
	query := `
		INSERT INTO user (first_name, last_name, role, email, password, phone, gender, address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8); 
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 15)
	if err != nil {
		fmt.Printf("ERROR:: error hashing password: %s\n", err.Error())
		return err
	}
	if _, err := service.db.Exec(query, &params.FirstName, &params.LastName, &params.Role, &params.Email,
		&hashedPassword, &params.PhoneNumber, &params.Gender, &params.Address); err != nil {
		fmt.Printf("ERROR:: could not insert to user table: %s\n", err.Error())
		return err
	}

	return nil
}

func (service userService) Update(id int, params view.UserView) error {
	getQuery := `
		SELECT id, first_name, last_name, role, email, phone, gender, address FROM user WHERE id = $1;
	`

	var user domain.User
	if err := service.db.QueryRow(getQuery, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Role,
		&user.Email, &user.Phone, &user.Gender, &user.Address); err != nil {
		return err
	}

	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Role = params.Role
	user.Email = params.Email
	user.Phone = params.PhoneNumber
	user.Gender = params.Gender
	user.Address = params.Address

	updateQuery := `
		UPDATE user
		SET first_name = $1, last_name = $2, role = $3, email = $4, phone = $5, gender = $6, address = $7
		WHERE id = $8
	`
	if _, err := service.db.Exec(updateQuery, &user.FirstName, &user.LastName, &user.Role, &user.Email, &user.Phone,
		&user.Gender, &user.Address, &id); err != nil {
		return err
	}

	return nil
}
