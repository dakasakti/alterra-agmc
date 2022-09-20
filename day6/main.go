package main

import (
	"github.com/dakasakti/day2/config"
	"github.com/dakasakti/day2/database"
	"github.com/dakasakti/day2/internal/factory"
	"github.com/dakasakti/day2/internal/http"
	"github.com/dakasakti/day2/internal/middlewares"
	"github.com/dakasakti/day2/internal/routes"
)

func main() {
	database.InitMigrate()

	c := config.GetConfig()
	f := factory.NewFactory(c)

	e := routes.New()
	middlewares.LoggerMiddleware(e)

	http.NewHttp(e, f)
	e.Logger.Fatal(e.Start(":8000"))

}
