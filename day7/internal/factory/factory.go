package factory

import (
	"github.com/dakasakti/day2/config"
	"github.com/dakasakti/day2/database"
	"github.com/dakasakti/day2/internal/repository/user"
)

type Factory struct {
	UserRepository user.UserRepository
}

func NewFactory(config *config.AppConfig) *Factory {
	db := database.InitMySQL(config)
	client := database.InitMongoDB(config)
	mc := database.InitCollection(client, config, "users")

	return &Factory{
		user.NewUserRepository(db, mc),
	}
}
