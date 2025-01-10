package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID         int
	FirstName  string
	LastName   string
	HasLoan    bool
	HasDeposit bool
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Index() ([]User, error) {
	var users []User

	query := `SELECT * FROM users`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.HasLoan, &user.HasDeposit)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (m *UserModel) Insert(firstName string, lastName string) (int, error) {
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id`

	var id int

	err := m.DB.QueryRow(query, firstName, lastName).Scan(&id)
	if err != nil {
		return id, err
	}

	queryAfter := `INSERT INTO accounts (user_id, balance) VALUES ($1, 0)`
	m.DB.QueryRow(queryAfter, id)

	return id, nil
}

func (m *UserModel) Exists(firstName string, lastName string) (bool, User) {
	query := `SELECT EXISTS( SELECT 1 FROM users WHERE first_name = $1 AND last_name = $2) `
	queryIfExists := `SELECT * FROM users WHERE first_name = $1 AND last_name = $2`

	var exists bool
	err := m.DB.QueryRow(query, firstName, lastName).Scan(&exists)
	if err != nil {
		return exists, User{}
	}

	var user User

	if exists {
		m.DB.QueryRow(queryIfExists, firstName, lastName).Scan(&user.ID, &user.FirstName, &user.LastName, &user.HasLoan, &user.HasDeposit)
		return exists, user
	}

	return exists, User{}
}

func (m *UserModel) ChangeStatusLoan(id int) error {
	query := `UPDATE users SET has_loan = NOT has_loan WHERE id = $1`
	_, err := m.DB.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) ChangeStatusDeposit(id int) error {
	query := `UPDATE users SET has_deposit = true WHERE id = $1`
	_, err := m.DB.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) CheckStatusDeposit(id int) (bool, error) {
	query := `SELECT has_deposit FROM users WHERE id = $1`
	var hasDeposit bool
	err := m.DB.QueryRow(query, id).Scan(&hasDeposit)
	if err != nil {
		return false, err
	}
	return hasDeposit, nil
}

func (m *UserModel) HasDeposit(id int) bool {
	var hasDeposit bool
	query := `SELECT has_deposit FROM users WHERE id = $1`
	m.DB.QueryRow(query, id).Scan(&hasDeposit)
	return hasDeposit
}

func (m *UserModel) Update(column string, name string, userID int) error {

	query := fmt.Sprintf(`UPDATE users SET %s = $1 WHERE id = $2`, column)
	_, err := m.DB.Query(query, name, userID)
	if err != nil {
		return err
	}

	return nil
}
