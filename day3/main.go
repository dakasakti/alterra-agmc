package main

import (
	"github.com/dakasakti/day2/lib/database"
	"github.com/dakasakti/day2/lib/validation"
	"github.com/dakasakti/day2/middlewares"
	"github.com/dakasakti/day2/routes"
	"github.com/go-playground/validator/v10"
)

func main() {
	database.InitMigrate()

	e := routes.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	middlewares.LoggerMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))

}
