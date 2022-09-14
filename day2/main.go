package main

import (
	"github.com/dakasakti/day2/lib/database"
	"github.com/dakasakti/day2/routes"
)

func main() {
	database.InitMigrate()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
