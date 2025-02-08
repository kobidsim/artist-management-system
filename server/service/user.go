package service

import (
	"artist-management-system/domain"
	"database/sql"
	"fmt"
)

type userService struct {
	db *sql.DB
}

type UserService interface {
	All() ([]domain.User, error)
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
