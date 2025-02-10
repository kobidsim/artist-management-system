package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	db *sql.DB
}

type UserService interface {
	All() ([]domain.User, error)
	Create(params view.UserView) error
	Update(id int, params view.UserView) error
	Delete(id int) error
}

func NewUserService(db *sql.DB) UserService {
	return userService{
		db: db,
	}
}

func (service userService) All() ([]domain.User, error) {
	query := `
		SELECT id, first_name, last_name, role, email, phone, gender, address, dob FROM user;
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
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Role, &user.Email, &user.Phone, &user.Gender, &user.Address, &user.Dob); err != nil {
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
		fmt.Printf("ERROR:: error hashing password: %s\n", err.Error())
		return err
	}
	if _, err := service.db.Exec(query, &params.FirstName, &params.LastName, &params.Role, &params.Email,
		&hashedPassword, &params.PhoneNumber, &params.Gender, &params.Address, &dob, &createdAt); err != nil {
		fmt.Printf("ERROR:: could not insert to user table: %s\n", err.Error())
		return err
	}

	return nil
}

func (service userService) Update(id int, params view.UserView) error {
	dob, err := time.Parse("2006-01-02T15:04:05.000Z", params.DOB)
	if err != nil {
		return err
	}

	updatedAt := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	getQuery := `
		SELECT id, first_name, last_name, role, email, phone, password, gender, address FROM user WHERE id = $1;
	`

	var user domain.User
	if err := service.db.QueryRow(getQuery, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Role,
		&user.Email, &user.Phone, &user.Password, &user.Gender, &user.Address); err != nil {
		return err
	}

	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Role = params.Role
	user.Email = params.Email
	user.Phone = params.PhoneNumber
	user.Gender = params.Gender
	user.Address = params.Address
	user.Dob = dob.Format("2006-01-02T15:04:05.000Z")

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 15)
	if err != nil {
		return err
	}
	user.Password = string(newHashedPassword)

	updateQuery := `
		UPDATE user
		SET first_name = $1, last_name = $2, role = $3, email = $4, phone = $5, password = $6, gender = $7, address = $8, dob = $9, updated_at = $10
		WHERE id = $11
	`
	if _, err := service.db.Exec(updateQuery, &user.FirstName, &user.LastName, &user.Role, &user.Email, &user.Phone,
		&user.Password, &user.Gender, &user.Address, &user.Dob, &updatedAt, &id); err != nil {
		return err
	}

	return nil
}

func (service userService) Delete(id int) error {
	query := `
		DELETE FROM user WHERE id = $1
	`

	if _, err := service.db.Exec(query, &id); err != nil {
		return err
	}

	return nil
}
