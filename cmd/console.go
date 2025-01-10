package main

import (
	"awesomeProject3/internal/models"
	"fmt"
	"os"
)

/////////////////////////////////////////////////////////

func EnterFirstName() string {

	var firstName string

	for {
		fmt.Print("Enter your first name: ")
		_, err := fmt.Scan(&firstName)
		if err != nil || isValidName(firstName) == false {
			fmt.Println("Please, enter a valid first name")
			continue
		} else {
			break
		}
	}

	return firstName

}

func EnterLastName() string {
	var lastName string

	for {
		fmt.Print("Enter your last name: ")
		_, err := fmt.Scan(&lastName)
		fmt.Println()
		if err != nil || isValidName(lastName) == false {
			fmt.Print("Please, enter a valid last name\n\n")
			continue
		} else {
			break
		}
	}

	return lastName
}

func EnterPrompt() string {
	fmt.Print("\nYour action: ")

	var prompt string
	fmt.Scan(&prompt)
	fmt.Println()
	return prompt
}

func EnterValidPrompt() {
	fmt.Print("Please, enter a valid prompt: ")
}

///////////////////////////////////////////////////////////

func (app *application) Run() {
	fmt.Println("Welcome to our application!\n[1] Login\n[2] Register")
	switch EnterPrompt() {
	case "1":
		app.Login()
	case "2":
		app.Register()
	default:
		EnterValidPrompt()
	}
}

func (app *application) MainMenu() {
	fmt.Printf("Welcome, %s! What do you want to do?\n[1] Check for balance\n[2] Withdraw money\n[3] Take a loan\n[4] Deposit money\n[5] Edit info\n[6] My history\n[7] Exit\n", app.CurrentUser.FirstName)
	for {
		switch EnterPrompt() {
		case "1":
			app.Balance()
		case "2":
			app.Withdraw()
		case "3":
			app.TakeALoan()
		case "4":
			app.Deposit()
		case "5":
			app.EditInfo()
		case "6":
			app.History()
		case "7":
			app.Exit()
		default:
			fmt.Print("Please, enter a valid prompt: ")
		}
	}

}

func (app *application) Exit() {
	fmt.Println("Thanks for using our application!")
	os.Exit(1)
}

//////////////////////////////////////////////////////////

func (app *application) Register() {

	user := models.User{
		FirstName: EnterFirstName(),
		LastName:  EnterLastName(),
	}

	isUserAlreadyExists, _ := app.IsUserAlreadyExists(user)

	if isUserAlreadyExists {
		fmt.Println("This user already exists. Do you want to sign in? \n[1] YES \n[2] NO")
		for {
			switch EnterPrompt() {
			case "1":
				{
					app.CurrentUser = user
					app.Login()
				}
			case "2":
				app.Exit()
			default:
				fmt.Print("Please, enter a valid prompt: ")
			}
		}
	} else {

		id, err := app.userModel.Insert(user.FirstName, user.LastName)
		if err != nil {
			app.serverError(err)
		}
		user.ID = id
		app.CurrentUser = user
		//err = app.accountModel.Insert(app.CurrentUser.ID, 0)
		//if err != nil {
		//	app.serverError(err)
		//}
		app.MainMenu()
	}

	app.Users = append(app.Users, user)
}

func (app *application) Login() {

	user := models.User{
		FirstName: EnterFirstName(),
		LastName:  EnterLastName(),
	}

	isUserAlreadyExists, user := app.IsUserAlreadyExists(user)

	if isUserAlreadyExists {
		app.CurrentUser = user
		app.MainMenu()
	} else {
		fmt.Println("User cannot be found\n[1] Try again\n[2] Register\n[3] Exit")
		switch EnterPrompt() {
		case "1":
			app.Login()
		case "2":
			app.Register()
		case "3":
			app.Exit()
		default:
			EnterValidPrompt()
		}
	}
}
