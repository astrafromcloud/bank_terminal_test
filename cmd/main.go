package main

import (
	"awesomeProject3/internal/models"
	"database/sql"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

type application struct {
	CurrentUser  models.User
	Users        []models.User
	Accounts     []models.BankAccount
	logger       *slog.Logger
	userModel    *models.UserModel
	accountModel *models.BankAccountModel
	logModel     *models.LogModel
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   false,
		Level:       nil,
		ReplaceAttr: nil,
	}))

	db, err := openDB()
	if err != nil {
		logger.Error(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(db)

	app := application{
		CurrentUser:  models.User{},
		userModel:    &models.UserModel{DB: db},
		accountModel: &models.BankAccountModel{DB: db},
		logModel:     &models.LogModel{DB: db},
		logger:       logger,
	}

	app.userModel.Exists("Rome", "Zhailau")
	app.Run()

}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=bank_test sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
