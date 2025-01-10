package main

import "fmt"

func (app *application) serverError(err error) {

	fmt.Print("\nInternal server error. Try again later!\n\n")
	app.logger.Error(err.Error())
}
