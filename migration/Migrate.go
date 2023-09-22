package migration

import (
	"github.com/shiv122/go-todo/app/models"
	"github.com/shiv122/go-todo/connection"
)

func Migrate() {
	connection.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)

}
