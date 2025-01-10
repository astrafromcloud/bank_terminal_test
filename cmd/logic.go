package main

import (
	"awesomeProject3/internal/models"
	"fmt"
)

func (app *application) IsUserAlreadyExists(currentUser models.User) (bool, models.User) {

	var exists bool
	exists, user := app.userModel.Exists(currentUser.FirstName, currentUser.LastName)
	return exists, user
}

func isValidName(name string) bool {

	for _, char := range name {
		if !((char >= 65 && char <= 90) || (char >= 96 && char <= 122)) {
			return false
		}
	}

	return true
}

func (app *application) Balance() {
	fmt.Printf("Your balance is %.2f₸\n", app.accountModel.CheckBalance(app.CurrentUser.ID))
}

func (app *application) Withdraw() {

	fmt.Print("Enter how much money do you want to withdraw: ")
	var money float64
	var userID = app.CurrentUser.ID
	_, err := fmt.Scan(&money)
	if err != nil {
		app.serverError(err)
	}

	if money > app.accountModel.CheckBalance(app.CurrentUser.ID) {
		fmt.Print("Not enough money! Try again later!\n\n")
		app.MainMenu()
	} else {
		app.accountModel.ChangeBalance(app.accountModel.CheckBalance(app.CurrentUser.ID)-money, userID)
		err := app.logModel.Withdraw(userID, money)
		if err != nil {
			app.serverError(err)
		}
		fmt.Printf("Withdrawed successfully! In your balance: %.2f₸", app.accountModel.CheckBalance(userID))
	}

}

func (app *application) TakeALoan() {

	fmt.Print("Enter how much money do you want to get: (max - 10,000,000₸) ")

	var money float64
	var userID = app.CurrentUser.ID
	_, err := fmt.Scan(&money)
	if err != nil {
		app.serverError(err)
	}
	if money > 10000000 {
		fmt.Println("Maximum is 10,000,000₸")
		app.MainMenu()
	} else {
		app.accountModel.ChangeBalance(app.accountModel.CheckBalance(app.CurrentUser.ID)+money, app.CurrentUser.ID)
		err := app.logModel.LoanOn(userID, money)
		if err != nil {
			app.serverError(err)
		}
		fmt.Printf("Took successfully! In your balance: %.2f₸", app.accountModel.CheckBalance(app.CurrentUser.ID))
	}

}

func (app *application) Deposit() {

	fmt.Print("Enter how much money do you want to deposit: (with 15%) ")
	var money float64
	_, err := fmt.Scan(&money)
	if err != nil {
		app.serverError(err)
	}

	userID := app.CurrentUser.ID

	if money > app.accountModel.CheckBalance(userID) {
		fmt.Print("Not enough money! Try again later!\n\n")
		app.MainMenu()
	} else {

		deposit := app.accountModel.CheckDeposit(userID)

		if app.userModel.HasDeposit(userID) {
			err := app.accountModel.ChangeDeposit(deposit+money, userID)
			if err != nil {
				app.serverError(err)
			}

			deposit += money

			err = app.logModel.DepositAdd(userID, deposit)
			if err != nil {
				app.serverError(err)
			}
		} else {
			err := app.accountModel.ChangeDeposit(money, userID)
			if err != nil {
				app.serverError(err)
			}
			err = app.userModel.ChangeStatusDeposit(userID)
			if err != nil {
				app.serverError(err)
			}

			deposit = money

			err = app.logModel.DepositOn(userID, deposit)
			if err != nil {
				app.serverError(err)
			}
		}
		app.accountModel.ChangeBalance(app.accountModel.CheckBalance(userID)-money, userID)

		fmt.Printf("Successfully! Your deposit: %.2f₸", deposit)
	}

}

func (app *application) EditInfo() {

	var userID = app.CurrentUser.ID

	var newFirstName, newLastName string
	fmt.Println("Your previous first name is ", app.CurrentUser.FirstName)
	fmt.Print("Enter name you want to switch: ")
	_, err := fmt.Scan(&newFirstName)
	if err != nil {
		app.serverError(err)

	}
	fmt.Println()
	err = app.userModel.Update("first_name", newFirstName, userID)
	if err != nil {
		app.serverError(err)
	}

	fmt.Println("Your previous last name is ", app.CurrentUser.LastName)
	fmt.Print("Enter last name you want to switch: ")
	_, err = fmt.Scan(&newLastName)
	if err != nil {
		app.serverError(err)

	}
	fmt.Println()
	err = app.userModel.Update("last_name", newLastName, userID)
	if err != nil {
		app.serverError(err)
	}

	fmt.Println("Your info changed successfully!")

}

func (app *application) History() {
	logs, err := app.logModel.Get(app.CurrentUser.ID)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(logs); i++ {
		fmt.Printf("%d. %s\n", i+1, logs[i])
	}
	fmt.Println()
}
