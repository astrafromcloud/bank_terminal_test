package models

import (
	"database/sql"
	"time"
)

type BankAccount struct {
	ID      int
	UserID  int
	Balance float64
}

type BankAccountModel struct {
	DB *sql.DB
}

func (m *BankAccountModel) Insert(userId int, balance float64) error {
	query := `INSERT INTO accounts (user_id, balance) VALUES ($1, $2) RETURNING id`

	var id int

	err := m.DB.QueryRow(query, userId, balance).Scan(&id)
	if err != nil {
		return err
	}

	return nil

}

func (m *BankAccountModel) CheckBalance(id int) float64 {
	query := `SELECT balance FROM accounts WHERE user_id = $1`
	var balance float64
	m.DB.QueryRow(query, id).Scan(&balance)

	return balance
}

func (m *BankAccountModel) ChangeBalance(money float64, userId int) {
	query := `UPDATE accounts SET balance = $1 WHERE id = $2`
	m.DB.QueryRow(query, money, userId)
}

func (m *BankAccountModel) CheckDeposit(id int) float64 {
	query := `SELECT deposit FROM accounts WHERE user_id = $1`
	queryDeposit := `SELECT max(created_at) FROM logs WHERE type = 'deposit_on'`

	var deposit float64
	var depositCreatedAt time.Time

	m.DB.QueryRow(query, id).Scan(&deposit)
	m.DB.QueryRow(queryDeposit).Scan(&depositCreatedAt)

	deposit = time.Since(depositCreatedAt).Hours()*1.15*deposit + deposit
	return deposit
}

func (m *BankAccountModel) ChangeDeposit(money float64, userId int) error {
	query := `UPDATE accounts SET deposit = $1 WHERE id = $2`
	_, err := m.DB.Query(query, money, userId)
	if err != nil {
		return err
	}
	return nil
}
