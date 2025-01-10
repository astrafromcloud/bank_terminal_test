package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Log struct {
	ID        int
	UserID    int
	Log       string
	Action    string
	CreatedAt time.Time
}

type LogModel struct {
	DB *sql.DB
}

func (m *LogModel) DepositOn(userID int, deposit float64) error {
	query := `INSERT INTO logs (user_id, log, type, created_at) VALUES ($1, $2, $3, $4)`
	_, err := m.DB.Query(query, userID, fmt.Sprintf("Deposit has been opened for %.2f₸", deposit), "deposit_on", time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (m *LogModel) DepositAdd(userID int, deposit float64) error {
	query := `INSERT INTO logs (user_id, log, type, created_at) VALUES ($1, $2, $3, $4)`
	_, err := m.DB.Query(query, userID, fmt.Sprintf("Added to deposit %.2f₸", deposit), "deposit_add", time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (m *LogModel) LoanOn(userID int, money float64) error {
	query := `INSERT INTO logs (user_id, log, type, created_at) VALUES ($1, $2, $3, $4)`
	_, err := m.DB.Query(query, userID, fmt.Sprintf("Loan taken from bank for %.2f₸", money), "loan_on", time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (m *LogModel) Withdraw(userID int, money float64) error {
	query := `INSERT INTO logs (user_id, log, type, created_at) VALUES ($1, $2, $3, $4)`
	_, err := m.DB.Query(query, userID, fmt.Sprintf("Withdrawed %.2f₸", money), "withdraw", time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (m *LogModel) Get(userID int) ([]string, error) {

	var logs []string

	query := `SELECT log FROM logs WHERE user_id = $1`
	rows, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var log string
		err := rows.Scan(&log)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil

}
